package main

import (
	"bytes"
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/util/json"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"k8s.io/klog/v2"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"
)

var (
	kubeconfig *string
	namespace  string
	all        bool

	cmd = &cobra.Command{
		Use:  "kubectl-really-get-all",
		Long: "kubectl-really-get-all",
		Run:  cli,
	}
	out = bytes.NewBufferString("")
)

type TableRoundtripper struct {
	http.RoundTripper

	wrapRoundTripper http.RoundTripper
}

func (t *TableRoundtripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Accept", "application/json;as=Table;g=meta.k8s.io;v=v1")

	return t.wrapRoundTripper.RoundTrip(req)
}

func init() {
	cmd.Flags().StringVarP(&namespace, "Namespace", "n", v1.NamespaceDefault, "")
	cmd.Flags().BoolVarP(&all, "All", "A", false, "")

	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "(optional) absolute path to the kubeconfig file")
	}

	namespace = strings.TrimSpace(namespace)
}

func main() {
	_, err := cmd.ExecuteC()
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
	}
}

func cli(cmd *cobra.Command, args []string) {
	klog.SetOutput(io.Discard)
	flags := &flag.FlagSet{}
	klog.InitFlags(flags)
	flags.Set("logtostderr", "false")

	client, clientset, err := buildClients()
	if err != nil {
		_ = fmt.Errorf("%s", err.Error())
		return
	}

	_, resources, err := clientset.ServerGroupsAndResources()
	if err != nil {
		_ = fmt.Errorf("error fetching groups and resources")
		return
	}

	for _, resource := range resources {
		groupVersionPart := strings.Split(resource.GroupVersion, "/")
		var group string
		var version string

		if len(groupVersionPart) >= 2 {
			group = groupVersionPart[0]
			version = groupVersionPart[1]
		} else {
			group = ""
			version = groupVersionPart[0]
		}

		for _, apiResources := range resource.APIResources {
			parts := strings.Split(apiResources.Name, "/")
			if len(parts) > 1 {
				continue
			}
			gvr := schema.GroupVersionResource{Group: group, Version: version, Resource: apiResources.Name}
			var list *unstructured.UnstructuredList
			if all {
				list, err = client.Resource(gvr).Namespace(v1.NamespaceAll).List(context.Background(), v1.ListOptions{})
			} else {
				list, err = client.Resource(gvr).Namespace(namespace).List(context.Background(), v1.ListOptions{})
			}
			if err != nil {
				continue
			}
			table := &v1.Table{}
			err = runtime.DefaultUnstructuredConverter.FromUnstructured(list.Object, table)
			if err != nil {
				panic(err)
			}
			isNamespaced := checkIfNamespaced(table)

			// skip printing if there is 0 elements in resourceList
			if len(table.Rows) <= 0 {
				continue
			}

			printTable(table, &gvr, isNamespaced)
		}
	}
	_, _ = fmt.Fprint(os.Stdout, out.String())
}

func buildClients() (*dynamic.DynamicClient, *kubernetes.Clientset, error) {
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {
		return nil, nil, errors.New("error creating kubernetes client")
	}
	dynamicClient, err := rest.HTTPClientFor(config)
	if err != nil {
		return nil, nil, errors.New("error creating kubernetes client")
	}
	restTransport, err := rest.TransportFor(config)
	if err != nil {
		return nil, nil, errors.New("error creating kubernetes client")
	}
	dynamicClient.Transport = &TableRoundtripper{wrapRoundTripper: restTransport}
	client, err := dynamic.NewForConfigAndClient(config, dynamicClient)
	if err != nil {
		return nil, nil, errors.New("error creating kubernetes client")
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		return nil, nil, errors.New("error creating kubernetes client")
	}
	return client, clientset, nil
}

func getNamespaceFromRawExtension(raw *runtime.RawExtension) *string {
	intoMap := make(map[string]interface{})

	err := json.Unmarshal(raw.Raw, &intoMap)
	if err != nil {
		return nil
	}

	v, found, _ := unstructured.NestedString(intoMap, "metadata", "namespace")
	if !found {
		return nil
	}

	return &v
}

func checkIfNamespaced(table *v1.Table) bool {
	if len(table.Rows) < 0 {
		first := table.Rows[0]

		namespaceValue := getNamespaceFromRawExtension(&first.Object)
		return namespaceValue != nil
	}
	return false
}

func printTable(table *v1.Table, gvr *schema.GroupVersionResource, isNamespaced bool) {
	//b := bytes.NewBufferString("")
	tabw := tabwriter.NewWriter(out, 8, 8, 2, '\t', 0)

	headerValues := ""
	if isNamespaced {
		headerValues += fmt.Sprintf("%v\t", strings.ToUpper("Namespace"))
	}
	for _, columnDefinition := range table.ColumnDefinitions {

		if columnDefinition.Priority <= 0 {
			// columnDefitions = append(columnDefitions, columnDefinition.Name)
			headerValues += fmt.Sprintf("%v\t", strings.ToUpper(columnDefinition.Name))
		}

	}
	_, _ = fmt.Fprintln(tabw, headerValues)

	for _, tableItem := range table.Rows {
		cellValues := ""
		f := getNamespaceFromRawExtension(&tableItem.Object)
		if isNamespaced {
			cellValues += fmt.Sprintf("%v\t", *f)
		}

		for i, cellValue := range tableItem.Cells {
			if table.ColumnDefinitions[i].Priority == 0 {
				if strings.ToLower(table.ColumnDefinitions[i].Name) == "name" {
					cellValues += fmt.Sprintf("%s/%v\t", optionallyTranslateToSingular(gvr.Resource), cellValue)
				} else {
					cellValues += fmt.Sprintf("%v\t", cellValue)
				}
				// cellValues = append(cellValues, fmt.Sprintf("[%s : %v]", table.ColumnDefinitions[i].Name, cellValue))
			}
		}
		_, _ = fmt.Fprintln(tabw, cellValues)
	}

	_ = tabw.Flush()
	_, _ = fmt.Fprintln(out)
	//tstr := b.String()
	//fmt.Println(tstr)
}

func optionallyTranslateToSingular(in string) string {
	last := in[len(in)-1]
	if last == 's' {
		return in[0 : len(in)-1]
	}

	return in
}

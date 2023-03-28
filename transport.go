package main

import "net/http"

const (
	// KubernetesTableHeader is a special Header, when set kube-api will return Kinds wrapped in a metav1.Table
	KubernetesTableHeader = "application/json;as=Table;g=meta.k8s.io;v=v1"
)

// ProxyTransport takes in a http.RoundTripper and proxies calls to it, adding a specific Accept header
type ProxyTransport struct {
	http.RoundTripper

	ProxiedTransport http.RoundTripper
}

func (t *ProxyTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Set("Accept", KubernetesTableHeader)

	return t.ProxiedTransport.RoundTrip(req)
}

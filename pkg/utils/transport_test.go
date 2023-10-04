package utils

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestProxyTransport_RoundTrip(t *testing.T) {
	headerFound := false
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		headerStr := r.Header.Get("Accept")
		if strings.Contains(headerStr, KubernetesTableHeader) {

			headerFound = true
		}
	}))
	defer testServer.Close()
	httpClient := http.Client{
		Transport: &ProxyTransport{ProxiedTransport: http.DefaultTransport},
	}

	_, err := httpClient.Get(testServer.URL)
	if err != nil {
		t.Error(err)
	}

	if headerFound == false {
		t.Fatalf("kubernetes table header missing")
	}
}

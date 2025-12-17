package utils

import (
	"fmt"
	
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	// testMux is the HTTP request multiplexer used with the test server.
	testMux *http.ServeMux

	// testClient is the Jira client being tested.
	testDownloader *Downloader

	// testServer is a test HTTP server used to provide mock API responses.
	testServer *httptest.Server
)

// setup sets up a test HTTP server along with a jira.Client that is configured to talk to that test server.
// Tests should register handlers on mux which provide mock responses for the API method being tested.
func setup() {
	// Test server
	testMux = http.NewServeMux()
	testServer = httptest.NewServer(testMux)

	// downloader configured to use test server
	testDownloader = NewDownloader().WithClient(testServer.Client())
	testDownloader.URL = testServer.URL
}

// teardown closes the test HTTP server.
func teardown() {
	testServer.Close()
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testRequestURL(t *testing.T, r *http.Request, want string) {
	if got := r.URL.String(); !strings.HasPrefix(got, want) {
		t.Errorf("Request URL: %v, want %v", got, want)
	}
}

func testRequestParams(t *testing.T, r *http.Request, want map[string]string) {
	params := r.URL.Query()

	if len(params) != len(want) {
		t.Errorf("Request params: %d, want %d", len(params), len(want))
	}

	for key, val := range want {
		if got := params.Get(key); val != got {
			t.Errorf("Request params: %s, want %s", got, val)
		}

	}

}

func TestDownload(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/Go.gitignore", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("test call handled", "request_uri", r.RequestURI)
		testMethod(t, r, "GET")
		fmt.Fprint(w, "test content")
	})

	body, err := testDownloader.Download()
	if err != nil {
		fmt.Println(testDownloader.GetLog())
		t.Fatalf("Download returned error: %v", err)
	}

	want := "test content"
	if string(body) != want {
		t.Errorf("Download = %q, want %q", string(body), want)
	}

}

func TestDownload_NotFound(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/NonExistentLang.gitignore", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("test call handled", "request_uri", r.RequestURI)
		testMethod(t, r, "GET")
		http.NotFound(w, r)
	})

	testDownloader.Language = "NonExistentLang"
	_,  err := testDownloader.Download()
	if err == nil {
		fmt.Println(testDownloader.GetLog())
		t.Fatalf("Download expected error but got none")
	}
	
}


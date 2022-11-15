package clients

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func startTestHTTPServer() *httptest.Server {
	ts := httptest.NewServer(
		http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				fmt.Fprintf(w, "You passed the test!")
			}))
	return ts
}

func TestTimeoutClient(t *testing.T) {
	ts := startTestHTTPServer()
	defer ts.Close()

	testClient := TimeoutClient(5 * time.Second)

	expected := "You passed the test!"

	got, err := testClient.Get(ts.URL)
	if err != nil {
		t.Fatal(err)
	}

	if expected != string(got) {
		t.Errorf("Expected: %s, Got: %s", expected, got)
	}
}

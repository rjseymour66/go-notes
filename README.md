# HTTP Clients

## Creating the client

You do not want to use the default client in the golang http package because you don't have access to fields like Timeout and RoundTripper.

## Testing the client

Use the following structure for a basic client test:

```golang
// Create test server
func startTestHTTPServer() *httptest.Server {
    testServer := httptest.NewServer (
        http.HandlerFunc(
            func(w http.ResponseWriter, r *http.Request) {
                // values that you want to test for
            }))
    return testServer
}

// Create the main test method and do the following:
// 1. Start the test server that you just created
// 2. defer the Close() for the test server
// 3. Create any clients and make any calls that you plan to test.
// 4. Check the expected value vs the actual returned value (got)

func TestXXXX(t *testing.T) {
    ts := startTestHTTPServer()
    defer ts.Close()

    expected := // values that you want to test for
    client := // client constructor
    resp, err := // client method call

    got, err := clientName(ts.URL)
    if err != nil {
        t.Fatal(err)
    }
    if expected != string(data) {
        t.Errorf("Expected %s, Got: %s", expected, got)
    }

```
The following methods are useful:
- `t.Fatal(error)` ends a test because of failure
- `t.Errorf(formatted string, vars)` to describe why expected != got


## Command line tools

# Print(x) statements
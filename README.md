# HTTP Clients

## Creating the client

You do not want to use the default client in the golang http package because you don't have access to fields like Timeout and RoundTripper.

## Testing the client

Client test files contain the following:
- A method that creates a test server. The method must return an `*httptest.Server` object. To create one, you must run the `httptest.NewServer` function, and pass it an `http.HandlerFunc()` function.
 `http.HandlerFunc()` accepts an anonymous function that writes the response to the http.ResponseWriter.
- The main test method that actually tests your client and methods. Generally, the main test method does the following:
  1. Start the test server that you just created
  2. defer the Close() for the test server
  3. Create any clients and make any calls that you plan to test.
  4. Check the expected value vs the actual returned value (got)

```golang
// test server setup
func startTestHTTPServer() *httptest.Server {
    testServer := httptest.NewServer (
        http.HandlerFunc(
            func(w http.ResponseWriter, r *http.Request) {
                // values that you want to test for
            }))
    return testServer
}

// main test method
func TestXXXX(t *testing.T) {
    ts := startTestHTTPServer()
    defer ts.Close()

    expected := // values written to the w in the body 
                // of the testServer HandlerFunc()
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
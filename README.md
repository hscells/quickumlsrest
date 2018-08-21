# quickumlsrest

package quickumlsrest provides an API client to the quickUMLS-rest service.
The server can be found as a docker image from https://hub.docker.com/r/aehrc/quickumls-rest/.

Documentation: https://godoc.org/github.com/hscells/quickumlsrest

Example:

```go
client := quickumlsrest.NewClient("http://localhost:5000")
candidates, err := client.Match("cancer")
if err != nil {
    panic(err)
}

for _, candidate := range candidates {
    fmt.Println(candidate.CUI)
}
```
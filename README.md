# grun

### A go library helps developer to handle errors gracefully.

## Example

```go
grun.Run(func(handleErr grun.HandleErrFunc) {
    resp, err := http.Get("https://www.google.com")
    handleErr("httpGet", err)

    data, err := ioutil.ReadAll(resp.Body)
    handleErr("readBody", err)

    fmt.Println(string(data))
}).Catch(func(caughtError grun.CaughtError) {
    switch caughtError.Name {
    case "httpGet":
        fmt.Printf("error calling http.Get(): %s", caughtError.Err.Error())
    case "readBody":
        fmt.Printf("error reading response body: %s", caughtError.Err.Error())
    }
})
```
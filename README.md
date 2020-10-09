# grun

### A go library helps developer to handle errors gracefully.

## Use cases
When your go program becomes long, there should be cases that you need to handle a lot of errors within a function. 

For example:
```go
func makeHttpRequest(url string) error {
    resp, err := http.Get(url)
    if err != nil {
        // handle error for http.Get
    }

    data, err := ioutil.ReadAll(resp)
    if err != nil {
        // handle error for ioutil.ReadAll
    }

    err := processData(data)
    if err != nil {
        // handle error for processData
    }
    
    // Doing other things, which may generate other errors.
    // You need to handle them one by one.

    return nil
}
```

With `grun`, the code would become prettier:

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

With `grun` you code can handle all the errors within the function in a centralised place.

## API

### `Run` function

### `HandleErrFunc` function

### `Catch` function

### `CaughtError` struct
package main

import (
	"fmt"
	"github.com/vence722/grun"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	grun.Run(func(try grun.TryFunc) {
		conn := try(http.Get("https://www.google.com"))[0].(*http.Response)
		data := try(ioutil.ReadAll(conn.Body))[0].([]byte)
		fmt.Println("response:", string(data))
	}).Catch(func(err error) {
		log.Fatal("unexpected errors:", err)
	})
}

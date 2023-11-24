package main

import (
	"net/http"
	myapp "src/web1/myapp"
)

func main() {

	http.ListenAndServe(":3001", myapp.NewHttpHandler())
}

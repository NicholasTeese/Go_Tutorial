package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloServer)
	http.ListenAndServe(":8080", nil)

	//fmt.Println("hello world")
}

// HelloServer reads from the file "index" and returns the data inside, in the event of an error it will print the error instead
func HelloServer(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	//w.Write([]byte("<h1>Hello Server</h1>"))
	data, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println(err)
	}
	w.Write(data)
}

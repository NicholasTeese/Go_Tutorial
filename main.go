package main

import (
	"fmt"
)

//func main() {
//http.HandleFunc("/", HelloServer)
//http.ListenAndServe(":8080", nil)
/*
	fmt.Println("hello world")
  	Everything above was a short re-introduction to coding.
	Everything hence forth will be the official tutorial into Golang
	some of which will be re-hashing previously covered things,
	but for the sake of effective study I will continue with them all the same.	*/
//tutorial can be found at https://www.tutorialspoint.com/go/go_program_structure.htm

func main() {
	/*This is my first sample program.*/
	fmt.Println("Hello, World!")

	//variables tutorial
	var x float64
	x = 20.0
	y := 42
	c = "foo"

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(c)
	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)
	fmt.Printf("c is of type %T\n", c)
}

// HelloServer reads from the file "index" and returns the data inside
// in the event of an error it will print the error instead
/*func HelloServer(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
	//w.Write([]byte("<h1>Hello Server</h1>"))
	data, err := ioutil.ReadFile("index.html")
	if err != nil {
		fmt.Println(err)
	}
	w.Write(data)
}*/

//Using http package to build a simple server
package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

var name string

func sayHello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello "+name+"!\n") //prints to webpage
	fmt.Fprintf(w, "You are welcome :)")   //Also prints to webpage
}

func main() {
	//Ask user for name to be printed on webpage
	fmt.Println("What is your name?")
	fmt.Scanf("%s", &name)

	fmt.Println("Visit http://localhost:8080/ on your web browser if no error message displayed underneath: ")
	http.HandleFunc("/", sayHello)                 //Register the handler function (sayHello) for the given pattern ("/")
	log.Fatalln(http.ListenAndServe(":8080", nil)) //Set listening port. Handler is nil indicating that DefaultServeMux should be used. log.Fatal checks for error and outputs if any.
}

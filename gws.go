package main

import (
	"fmt"
	"net/http"
	"github.com/GroverSoans/GoWebServer/hello"         
    "github.com/GroverSoans/GoWebServer/syllabus" 
	"github.com/GroverSoans/GoWebServer/help"
)




func main() {
	//Routing
	http.HandleFunc("/hello-world", hello.HelloWorld)
	http.HandleFunc("/hello-world-html", hello.HelloWorldHTML)
	http.HandleFunc("/hello-world-json", hello.HelloWorldJSON)

	http.HandleFunc("/syllabus", syllabus.SyllabusHandler)

    // Create and Update Stubbed
    http.HandleFunc("/syllabus/create", syllabus.CreateStub)
    http.HandleFunc("/syllabus/update", syllabus.UpdateStub)

	http.HandleFunc("/help", help.Help)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}


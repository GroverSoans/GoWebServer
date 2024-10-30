package main

import (
	"os"
	"fmt"
	"net/http"
	"github.com/gorilla/handlers"
	"github.com/GroverSoans/GoWebServer/hello"         
    "github.com/GroverSoans/GoWebServer/syllabus" 
	"github.com/GroverSoans/GoWebServer/help"
	"github.com/GroverSoans/GoWebServer/dice"
)

func printHelp() {
	//Run help function: go run gws.go help
    helpMessage := `
		Go Web Server (GWS) Help:

		This application provides the following APIs:

		1. /hello-world         - Responds with plain text "Hello World – GWS"
		2. /hello-world-html    - Responds with HTML "Hello World &mdash; GWS"
		3. /hello-world-json    - Responds with JSON {"message": "Hello World - GWS"}
		4./syllabus   - Returns the syllabus JSON file.
   		- DELETE /syllabus/delete   - Returns "deleted – stubbed".
   		- POST   /syllabus/create   - Returns "create-stubbed".
   		- PUT    /syllabus/update   - Returns "update-stubbed".

		Run the server by typing: go run gws.go

		Example: http://localhost:8080/hello-world
	`
    fmt.Println(helpMessage)
}

func main() {
	if len(os.Args) > 1 && os.Args[1] == "help" {
        printHelp()
        return
    }

	corsObj := handlers.AllowedOrigins([]string{"*"})
	//Routing
	http.HandleFunc("/hello-world", hello.HelloWorld)
	http.HandleFunc("/hello-world-html", hello.HelloWorldHTML)
	http.HandleFunc("/hello-world-json", hello.HelloWorldJSON)

	http.HandleFunc("/syllabus", syllabus.SyllabusHandler)

    // Create Update and Delete
    http.HandleFunc("/syllabus/create", syllabus.CreateStub)
    http.HandleFunc("/syllabus/update", syllabus.UpdateStub)
	http.HandleFunc("/syllabus/delete", syllabus.DeleteStub)

	http.HandleFunc("/help", help.Help)

	//Roll dice Functionality
	http.HandleFunc("/api/rollDice", dice.DiceRoll)

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", handlers.CORS(corsObj)(http.DefaultServeMux))
}


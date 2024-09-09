package hello

import (
    "encoding/json"
    "net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello World - GWS"))
}

func HelloWorldHTML(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Hello World &mdash; GWS</h1>"))
}

func HelloWorldJSON(w http.ResponseWriter, r *http.Request){
	response := map[string]string{"message": "Hello World - GWS"}
    json.NewEncoder(w).Encode(response)
}

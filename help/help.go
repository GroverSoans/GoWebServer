package help

import (
	"net/http"
	"encoding/json"
)

func Help(w http.ResponseWriter, r *http.Request){
	helpInfo := map[string]string{
        "hello-world":       "Returns plain text: 'Hello World - GWS'",
        "hello-world-html":  "Returns HTML: 'Hello World — GWS'",
        "hello-world-json":  "Returns JSON: {'message': 'Hello World - GWS'}",
        "syllabus (GET)":    "Returns JSON of embedded syllabus data",
        "syllabus (DELETE)": "Stubbed response: 'deleted - stubbed'",
        "syllabus/create":   "Stubbed response: 'create - stubbed'",
        "syllabus/update":   "Stubbed response: 'update - stubbed'",
    }
	w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(helpInfo)
}
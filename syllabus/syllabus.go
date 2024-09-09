package syllabus

import (
    _ "embed"
    "encoding/json"
    "net/http"
)

// Embed the Syllabus JSON file
//go:embed syllabus.json
var syllabusJSON []byte

// Syllabus structure to parse JSON
type Syllabus struct {
    Course string   `json:"course"`
    Topics []string `json:"topics"`
}

// SyllabusHandler handles GET and DELETE
func SyllabusHandler(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        serveSyllabus(w)
    case http.MethodDelete:
        DeleteStub(w)
    default:
        http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
    }
}

// Serve the syllabus on GET request
func serveSyllabus(w http.ResponseWriter) {
    var syllabus Syllabus
    if err := json.Unmarshal(syllabusJSON, &syllabus); err != nil {
        http.Error(w, "Error parsing syllabus", http.StatusInternalServerError)
        return
    }
}

// Stubbed Delete
func DeleteStub(w http.ResponseWriter) {
    w.Write([]byte("deleted - stubbed"))
}

// Stubbed Create and Update
func CreateStub(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("create - stubbed"))
}

func UpdateStub(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("update - stubbed"))
}

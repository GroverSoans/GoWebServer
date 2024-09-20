package syllabus

import (
    _"embed"
    "encoding/json"
    "net/http"
)

//go:embed syllabus.json
var syllabusJSON []byte

// Syllabus structure to parse JSON
type Syllabus struct {
    ID                 string   `json:"id"`
    Section            string   `json:"section"`
    CRN                string   `json:"crn"`
    MeetingDays        string   `json:"meetingDays"`
    MeetingTimes       string   `json:"meetingTimes"`
    FinalExam          string   `json:"finalExam"`
    MeetingLocation    string   `json:"meetingLocation"`
    Course             Course   `json:"course"`
    Instructor         string   `json:"instructor"`
    Calendar           string   `json:"calendar"`
}

// Course structure nested within Syllabus
type Course struct {
    Name                string   `json:"name"`
    ID                  string   `json:"id"`
    Number              int      `json:"number"`
    CreditHours         int      `json:"creditHours"`
    Description         string   `json:"description"`
    Prerequisites       string   `json:"prerequisites"`
    LearningOutcomes    []string `json:"learningOutcomes"`
    ProgramOutcomes     []Outcome `json:"programOutcomes"`
    BaccalaureateCharacteristics []Outcome `json:"baccalaureateCharacteristics"`
    TextBook            string   `json:"textBook"`
    Modules             []string `json:"modules"`
}

// Outcome structure for program and baccalaureate characteristics
type Outcome struct {
    Value   int    `json:"value"`
    Outcome string `json:"outcome"`
}

// SyllabusHandler handles GET requests
func SyllabusHandler(w http.ResponseWriter, r *http.Request) {
    var syllabus Syllabus
    if err := json.Unmarshal(syllabusJSON, &syllabus); err != nil {
        http.Error(w, "Error parsing syllabus", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")

    // Use MarshalIndent to pretty-print the JSON
    prettyJSON, err := json.MarshalIndent(syllabus, "", "  ")
    if err != nil {
        http.Error(w, "Error encoding syllabus", http.StatusInternalServerError)
        return
    }

    w.Write(prettyJSON)
}

// Stubbed Create and Update
func CreateStub(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Syllabus created"))
}

func UpdateStub(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Syllabus updated"))
}

func DeleteStub(w http.ResponseWriter, r *http.Request){
    w.Write([]byte("Syllabus deleted"))
}

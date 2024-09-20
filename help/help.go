package help

import (
    "net/http"
)

func Help(w http.ResponseWriter, r *http.Request) {
    baseUrl := r.Host
	
	// Create an HTML response
    htmlResponse := `
    <html>
    <body>
        <h1>Available Endpoints</h1>
        <ul>
            <li><a href="http://` + baseUrl + `/hello-world">Hello World</a></li>
            <li><a href="http://` + baseUrl + `/hello-world-html">Hello World - HTML</a></li>
            <li><a href="http://` + baseUrl + `/hello-world-json">Hello World - JSON</a></li>
            <li><a href="http://` + baseUrl + `/syllabus">Get Syllabus</a></li>
            <li><a href="http://` + baseUrl + `/syllabus/create">Create Syllabus</a></li>
            <li><a href="http://` + baseUrl + `/syllabus/update">Update Syllabus</a></li>
            <li><a href="http://` + baseUrl + `/syllabus/delete">Delete Syllabus</a></li>
        </ul>
    </body>
    </html>`
    
    w.Write([]byte(htmlResponse))
}
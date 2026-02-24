/*   
Client selects file
 ↓
Browser sends multipart/form-data
 ↓
Go reads multipart form
 ↓
You access file
 ↓
Save it on disk
 */

package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Write([]byte(`
			<form method="POST" enctype="multipart/form-data">
				<input type="file" name="file">
				<button type="submit">Upload</button>
			</form>
		`))
		return
	}

	// Parse multipart form (10 MB max)
	r.ParseMultipartForm(10 << 20)

	// Get uploaded file
	file, handler, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error reading file", 400)
		return
	}
	defer file.Close()

	log.Println("Uploaded file:", handler.Filename)

	// Create file on server
	dst, err := os.Create(handler.Filename)
	if err != nil {
		http.Error(w, "Cannot create file", 500)
		return
	}
	defer dst.Close()

	// Copy uploaded content to server file
	io.Copy(dst, file)

	fmt.Fprintf(w, "File uploaded successfully: %s", handler.Filename)
}

func main() {

	http.HandleFunc("/upload", uploadHandler)

	log.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}

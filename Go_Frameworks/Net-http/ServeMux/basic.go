/* 
What is ServeMux? (plain words)

ServeMux = traffic police for HTTP requests.

It decides:

👉 /hello goes to which handler
👉 /upload goes to which handler
👉 /users goes to which handler


=>http.HandleFunc("/upload", uploadHandler)
  
You were secretly using:

http.DefaultServeMux

Go creates it automatically.
 
*/

package main
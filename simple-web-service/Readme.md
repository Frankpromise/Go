# Go Web Service

This simple Go web service demonstrates basic concepts of handling HTTP requests and serving content using the `net/http` package.

## Concepts Explained

### 1. HTTP Server Setup

```go
import (
	"fmt"
	"net/http"
)
```

### 2. Handling the Root Path ("/")
Set up a handler for the root path to respond with a plain text message.

```go
http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Simple web service written in Go")
})
```

### 3. Serving a File at "/home" Path
Set up a handler for the "/home" path to serve the content of the home.html file.
```go
http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "./home.html")
})
```

### 4. Start the HTTP server, making it listen on port 3000 to handle incoming requests.
```go
http.ListenAndServe(":3000", nil)
```
### Customization
Modify the plain text message in the root path response.
Update the content of the home.html file to suit your needs.

### Dependencies
. fmt: For formatted output.
. net/http: For handling HTTP requests and serving content.
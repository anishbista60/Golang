## Go HTTP Server Example

This repository contains a simple Go HTTP server example. The code sets up a server that serves static files from the "Template" directory and handles form submissions and a greeting route.

### Code Explanation

The provided Go code creates an HTTP server with the following features:

- Serves static files from the "Template" directory.
- Handles a "/form" route for processing form submissions.
- Handles a "/hello" route for a simple greeting.

```go
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// ... (see original code)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// ... (see original code)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	// ... (see original code)
}

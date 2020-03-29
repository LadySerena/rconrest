package api

import (
	"fmt"
	"net/http"
)

//SerenaHandler extra http handler to test pkg imports in app engine
func SerenaHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/serena" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello, mortal!")
}

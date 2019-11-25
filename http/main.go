package functions

import (
	"fmt"
	"net/http"
)

func TriggerHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprint(w, "Cloud Function より Hello World!(GET)")
	case http.MethodPost:
		fmt.Fprint(w, "Hello World!(POST)")
	default:
		http.Error(w, "405 - Method Not Allowed", http.StatusMethodNotAllowed)
	}
}


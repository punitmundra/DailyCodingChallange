package handler

import (
	"fmt"
	"net/http"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

// home is a simple HTTP handler function which writes a response.
func create(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Your create request is processed.")
}

// Router register necessary routes and returns an instance of a router.
func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/create", create).Methods("POST")
	return r
}

func TestRouter(t *testing.T) {
	//r := Router()
	//ts := httptest.NewServer(r)
	//defer ts.Close()
	payload := fmt.Sprintf(`{"name": {"name": %s},"type": {“message”: "%s"}}`, "bug", "issue")
	res, err := http.Post("http://127.0.0.1:5000/api/incident/create", "application/json", strings.NewReader(payload))
	if err != nil {
		t.Fatal(err)
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Status code for /create is wrong. Have: %d, want: %d.", res.StatusCode, http.StatusOK)
	}
}

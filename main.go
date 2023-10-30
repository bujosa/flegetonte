package main

import (
	"flegetonte/config"
	"flegetonte/email"
	"flegetonte/templates"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv()

	r := mux.NewRouter()

	r.HandleFunc("/send-email", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseMultipartForm(32 << 20)
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}

		to := r.FormValue("email")
		subject := r.FormValue("subject")
		firstName := r.FormValue("firstName")
		lastName := r.FormValue("lastName")

		if to == "" || subject == "" || firstName == "" || lastName == "" {
			http.Error(w, "Missing form data", http.StatusBadRequest)
			return
		}

		m := templates.UsePrimaryTemplate(to, subject, firstName, lastName)
		email.SendEmail(m)
	})

	// Launch the server.
	http.ListenAndServe(":8080", r)
}

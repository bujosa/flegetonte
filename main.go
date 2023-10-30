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
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}

		to := r.Form.Get("receiptEmail")
		subject := r.Form.Get("subject")
		firstName := r.Form.Get("firstName")
		lastName := r.Form.Get("lastName")

		m := templates.UsePrimaryTemplate(to, subject, firstName, lastName)
		email.SendEmail(m)
	})

	// Launch the server.
	http.ListenAndServe(":8080", r)
}

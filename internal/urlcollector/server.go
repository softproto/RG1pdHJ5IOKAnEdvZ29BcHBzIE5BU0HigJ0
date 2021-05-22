package urlcollector

import (
	// "context"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func RunServer() {
	fmt.Println("RunServer()")

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root"))
	})

	r.Get("/pictures", func(w http.ResponseWriter, r *http.Request) {
		startDate := r.URL.Query().Get("start_date")
		endDate := r.URL.Query().Get("end_date")
		err := addToQueue(startDate, endDate)
		if err != nil {
			w.Write([]byte(err.Error()))
		}
	})

	http.ListenAndServe(":80", r)
}

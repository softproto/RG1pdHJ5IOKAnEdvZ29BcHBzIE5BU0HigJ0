package urlcollector

import (
	// "context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var out struct {
	URLs   []string `json:",omitempty"`
	Errors []string `json:",omitempty"`
}

func RunServer(apiKey string, port string) {
	fmt.Println("RunServer()")

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root"))
	})

	r.Get("/pictures", func(w http.ResponseWriter, r *http.Request) {
		startDate := r.URL.Query().Get("start_date")
		endDate := r.URL.Query().Get("end_date")

		cd := runCollector(apiKey, startDate, endDate)

		out.URLs = cd.urls
		out.Errors = cd.errors
		b, err := json.Marshal(out)
		if err != nil {
			collectError("json.Marshal()", err, cd)
		}
		w.Write(b)
	})
	addr := ":" + port
	http.ListenAndServe(addr, r)
}

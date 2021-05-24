package urlcollector

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)


//Prepar and running the urlcollector server with basic routing
func RunServer(config *Config) {
	fmt.Println("RunServer()")

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("root"))
	})

	r.Get("/pictures", func(w http.ResponseWriter, r *http.Request) {
		startDate := r.URL.Query().Get("start_date")
		endDate := r.URL.Query().Get("end_date")

		cd := runCollector(config, startDate, endDate)

		w.Write(*cd.json())
	})
	addr := ":" + config.port
	http.ListenAndServe(addr, r)
}

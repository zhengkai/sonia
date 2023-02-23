package web

import (
	"fmt"
	"net/http"
	"project/zj"
	"time"

	"github.com/arl/statsviz"
)

// Server ...
func Server(port int) {

	addr := fmt.Sprintf(`localhost:%d`, port)

	mux := http.NewServeMux()
	mux.HandleFunc(`/`, failbackHandle)

	statsviz.Register(mux)

	s := &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  30 * time.Second,
	}

	zj.J(`start web server`, addr)

	s.ListenAndServe()
}

func failbackHandle(w http.ResponseWriter, r *http.Request) {
	zj.J(`failback handle`, r.URL.String())
}

package listener

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/hanchchch/gimi/packages/chopstick/pkg/container"
	"github.com/hanchchch/gimi/packages/chopstick/pkg/utils"
)

type HTTPListener struct {
	router *http.ServeMux
	addr   string
}

type HTTPListenerOptions struct {
	Addr string
}

func NewHTTPListener(options HTTPListenerOptions) *HTTPListener {
	return &HTTPListener{
		router: http.NewServeMux(),
		addr:   options.Addr,
	}
}

func (l *HTTPListener) Listen() error {
	log.Println("Listening on", l.addr)
	return http.ListenAndServe(l.addr, l.router)
}

func (l *HTTPListener) OnInvoke(callback HandlerFunc) error {
	l.router.HandleFunc("/invoke", logging(func(w http.ResponseWriter, req *http.Request) {
		var args container.TryArgs
		if err := json.NewDecoder(req.Body).Decode(&args); err != nil {
			http.Error(w, utils.ErrorJsonString(err), http.StatusBadRequest)
			return
		}
		data, err := callback(args)
		if err != nil {
			http.Error(w, utils.ErrorJsonString(err), http.StatusInternalServerError)
			return
		}
		body, _ := json.Marshal(data)

		w.WriteHeader(http.StatusOK)
		w.Write(body)
	}))
	return nil
}

func logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var start time.Time
		var end time.Time
		defer func() {
			log.Println(r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent(), end.Sub(start))
		}()
		start = time.Now()
		next(w, r)
		end = time.Now()
	}
}

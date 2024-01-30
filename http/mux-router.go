package http

import (
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

type muxRouter struct {
	client *mux.Router
}

var (
	mR 			*muxRouter
	routerOnce 	sync.Once
)

// Making muxRouter instance as singleton
func NewMuxRouter() IRouter {
	if mR == nil {
		routerOnce.Do(func() {
			client := mux.NewRouter().StrictSlash(true)
			mR = &muxRouter{client}
		})
	}
	return mR
}

func (r *muxRouter) ADDVERSION(uri string) {
	r.client = r.client.PathPrefix(uri).Subrouter()
}

func (r *muxRouter) GET(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	r.client.HandleFunc(uri,f).Methods("GET", "OPTIONS")
}

func (r *muxRouter) POST(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	r.client.HandleFunc(uri,f).Methods("POST", "OPTIONS")
}

func (r *muxRouter) PUT(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	r.client.HandleFunc(uri,f).Methods("PUT", "OPTIONS")
}

func (r *muxRouter) DELETE(uri string, f func(w http.ResponseWriter, r *http.Request)) {
	r.client.HandleFunc(uri,f).Methods("DELETE", "OPTIONS")
}

func (r *muxRouter) SERVE(port string) {
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "withCredentials"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	originsOk := handlers.AllowedOrigins([]string{"*"})
	credentialsOk := handlers.AllowCredentials()
	log.Printf("Mux HTTP server running on port %v", port)
	log.Fatal(http.ListenAndServe(":" + port, handlers.CORS(headersOk,methodsOk, originsOk, credentialsOk)(r.client)))
}



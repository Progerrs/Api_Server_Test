package apiserver

import (
	"awesomeProject/internal/app/store"
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

type server struct {
	router *mux.Router
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/users", s.handleUsersCreate())
	s.router.HandleFunc("/", s.handleHomePage())
}

func (s *server) handleHomePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Home")
	}
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		temp, _ := template.ParseFiles("static/index.html")
		temp.Execute(w, nil)
	}
}
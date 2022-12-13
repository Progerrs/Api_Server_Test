package apiserver

import (
	"awesomeProject/internal/app/model"
	"awesomeProject/internal/app/store"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
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
	s.router.HandleFunc("/api/users", s.handleUsersCreate()).Methods("POST")
	s.router.HandleFunc("/api/delete", s.handleUsersDelete()).Methods("POST")
	s.router.HandleFunc("/api/update", s.handleUsersUpdate()).Methods("POST")
	s.router.HandleFunc("/api/", s.handleHomePage()).Methods("GET")
}

func (s *server) handleHomePage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, s.store.User().Get())
	}
}

func (s *server) handleUsersDelete() http.HandlerFunc {
	type request struct {
		Id int `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		s.respond(w, r, http.StatusOK, s.store.User().Delete(req.Id))
	}
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	type request struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		BirthDay  string `json:"birth_day"`
		Gender    string `json:"gender"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &model.User{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			BirthDay:  req.BirthDay,
			Gender:    req.Gender,
		}
		if err := s.store.User().Create(u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusCreated, u)
	}

}

func (s *server) handleUsersUpdate() http.HandlerFunc {
	type request struct {
		Id        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		BirthDay  string `json:"birth_day"`
		Gender    string `json:"gender"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		req := &request{}
		if err := json.NewDecoder(r.Body).Decode(req); err != nil {
			s.error(w, r, http.StatusBadRequest, err)
			return
		}
		u := &model.User{
			FirstName: req.FirstName,
			LastName:  req.LastName,
			BirthDay:  req.BirthDay,
			Gender:    req.Gender,
		}

		if err := s.store.User().Update(req.Id, u); err != nil {
			s.error(w, r, http.StatusUnprocessableEntity, err)
			return
		}
		s.respond(w, r, http.StatusOK, u)
	}
}

func (s *server) error(w http.ResponseWriter, r *http.Request, code int, err error) {
	s.respond(w, r, code, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

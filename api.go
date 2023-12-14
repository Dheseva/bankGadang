package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int , v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func (http.ResponseWriter, *http.Request) error

type ApiError struct{
	Error string
}

func makeHTTPHandleFunc (f apiFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		if err := f(w,r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run(){
	router := mux.NewRouter()

	router.HandleFunc("/account", makeHTTPHandleFunc(s.handlerAccount))

	log.Println("JSON API server running on port: ", s.listenAddr)

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handlerAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return s.handlerGetAccount(w,r)
	}
	if r.Method == "POST" {
		return s.handlerCreateAccount(w,r)
	}
	if r.Method == "DELETE" {
		return s.handlerDeleteAccount(w,r)
	}
	return fmt.Errorf("request %s invalid", r.Method)
}

func (s *APIServer) handlerGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handlerCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handlerDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handlerTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}
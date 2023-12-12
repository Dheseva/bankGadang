package main

import "net/http"

type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run(){
	
}

func (s *APIServer) handlerAccount(w *http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handlerGetAccount(w *http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handlerCreateAccount(w *http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handlerDeleteAccount(w *http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handlerTransfer(w *http.ResponseWriter, r *http.Request) error {
	return nil
}
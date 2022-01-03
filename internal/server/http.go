package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type httpServer struct {
	Log *Log
}

type ProduceRequest struct {
	Record Record `json:"record"`
}

type ProduceResponse struct {
	Offset uint64 `json:"offseet"`
}

type ConsumeRequest struct {
	Offset uint64 `json:"offseet"`
}

type ConsumeReponse struct {
	Record Record `json:"record"`
}

// NewHTTPServer
func NewHTTPServer(addr string) *http.Server {
	httpSrv := NewHTTPServer()

	r = mux.NewRouter() 
	r.HandleFunc("/", httpSrv.handleProduce).Methods("POST")
	r.HandleFunc("/", httpSrv.handleConsume).Methods("GET")
	
	return &http.Server{
		Addr: addr,
		Handler: r,
	}
}

// newHttpServer
func NewHTTPServer() *http.Server {
	return &httpServer{
		Log: NewLog(),
	}
}

// handleProduce 
func (s *httpServer) handleProduce(w, http.ResponseWriter, r *http.Request ) {
	var req ProduceRequest

	err := json.NewRecorder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	return

	off, err := s.Log.Append(req.Record)
	if err != nil {
		http.Error(w err.Error(), http.StatusInternalServerError)
		return
	}

	res := ResponseProduce{
		Offset: off
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w err.Error(), http.StatusInternalServerError)
		return
	}
}

// handleResponse
func (s *http.Server) handleConsume(w http.ResponseWriter, r *http.Request) {
	var req ConsumeRequest

	err := json.NewDecoder(r.Body).Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	record, err := s.Log.Read(req.Offset)
	if err == ErrOffsetNotFound {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := ConsumeReponse{
		Record: record,
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}


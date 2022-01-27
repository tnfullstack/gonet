package server

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// httpServer
type httpServer struct {
	Log *Log
}

// WriteLogRequest
type WriteLogRequest struct {
	Record Record `json:"record"`
}

// WriteLogResponse
type WriteLogResponse struct {
	Offset uint64 `json:"offseet"`
}

// ReadLogRequest
type ReadLogRequest struct {
	Offset uint64 `json:"offseet"`
}

// ReadLogReponse
type ReadLogReponse struct {
	Record Record `json:"record"`
}

// NewHTTPServer
func NewHTTPServer(addr string) *http.Server {
	httpSrv := newHTTPServer()

	r := mux.NewRouter()
	r.HandleFunc("/", httpSrv.handleProduce).Methods("POST")
	r.HandleFunc("/", httpSrv.handleConsume).Methods("GET")

	return &http.Server{
		Addr:    addr,
		Handler: r,
	}
}

// newHttpServer
func newHTTPServer() *httpServer {
	return &httpServer{
		Log: NewLog(),
	}
}

// handleProduce
func (s *httpServer) handleProduce(w http.ResponseWriter, r *http.Request) {
	var req WriteLogRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	off, err := s.Log.Append(req.Record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	res := WriteLogResponse{
		Offset: off,
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// handleResponse
func (s *httpServer) handleConsume(w http.ResponseWriter, r *http.Request) {
	var req ReadLogRequest

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

	res := ReadLogReponse{
		Record: record,
	}

	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

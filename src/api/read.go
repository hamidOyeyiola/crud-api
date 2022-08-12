package api

import (
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/hamidOyeyiola/crud-api/controllers"
	model "github.com/hamidOyeyiola/crud-api/models"
)

type Read struct {
	model.Model
	controller.Retriever
}

func (r *Read) Read(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		return
	}
	q, ok := r.QueryToSelectWhere(id)
	h, b, ok := r.Retrieve(q, r.Model)
	var res controller.Response
	res.AddHeader(h).
		AddBody(b).
		Write(rw)
}

func (r *Read) ReadAll(rw http.ResponseWriter, req *http.Request) {
	q, _ := r.QueryToSelectAll()
	h, b, _ := r.Retrieve(q, r.Model)
	var res controller.Response
	res.AddHeader(h).
		AddBody(b).
		Write(rw)
}

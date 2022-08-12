package api

import (
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/hamidOyeyiola/crud-api/controllers"
	model "github.com/hamidOyeyiola/crud-api/models"
)

type Remove struct {
	model.Model
	controller.Deleter
}

func (d *Remove) Remove(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {

	}
	q, ok := d.QueryToDeleteWhere(id)
	h, b, ok := d.Delete(q)
	var res controller.Response
	res.AddHeader(h).
		AddBody(b).
		Write(rw)
}

func (r *Remove) RemoveAll(rw http.ResponseWriter, req *http.Request) {
	q, _ := r.QueryToDeleteAll()
	h, b, _ := r.Delete(q)
	var res controller.Response
	res.AddHeader(h).
		AddBody(b).
		Write(rw)
}

package api

import (
	"net/http"

	"github.com/gorilla/mux"
	controller "github.com/hamidOyeyiola/crud-api/controllers"
	model "github.com/hamidOyeyiola/crud-api/models"
)

type Change struct {
	model.Model
	controller.Updater
}

func (c *Change) Change(rw http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id, ok := vars["id"]
	if !ok {
		return
	}
	q, ok := c.QueryToUpdateWhere(req.Body, id)
	if !ok {
		return
	}
	h, b, _ := c.Update(q)
	var res controller.Response
	res.AddHeader(h).
		AddBody(b).
		Write(rw)
}

package api

import (
	"net/http"

	controller "github.com/hamidOyeyiola/crud-api/controllers"
	model "github.com/hamidOyeyiola/crud-api/models"
)

type New struct {
	model.Model
	controller.Creater
}

func (n *New) New(rw http.ResponseWriter, req *http.Request) {
	q, ok := n.QueryToInsert(req.Body)
	if !ok {

	}
	h, b, _ := n.Create(q)
	var res controller.Response
	res.AddHeader(h).
		AddBody(b).
		Write(rw)
}

package model

import (
	"io"

	interfaces "github.com/hamidOyeyiola/crud-api/interfaces"
)

type SQLQueryToInsert string

type SQLQueryToSelect string

type SQLQueryToDelete string

type SQLQueryToUpdate struct {
	Stmts  []string
	Values []string
	ID     string
}

type JSONObject string

type Model interface {
	FromQueryResult(interfaces.Iterator) (JSONObject, int)
	QueryToInsert(jsonObject io.Reader) (SQLQueryToInsert, bool)
	QueryToUpdateWhere(jsonObject io.Reader, value string) (SQLQueryToUpdate, bool)
	QueryToSelectWhere(value string) (SQLQueryToSelect, bool)
	QueryToSelectAll() (SQLQueryToSelect, bool)
	QueryToDeleteWhere(value string) (SQLQueryToDelete, bool)
	QueryToDeleteAll() (SQLQueryToDelete, bool)
}

/*
func GetQueryToSelect(m Model, id string) (q SQLQueryToSelect, ok bool) {
	q = SQLQueryToSelect(fmt.Sprintf("SELECT * FROM %s WHERE id IN (%s)", m.Name(), id))
	return q, true
}

func GetQueryToSelectAll(m Model) (q SQLQueryToSelect, ok bool) {
	q = SQLQueryToSelect(fmt.Sprintf("SELECT * FROM %s ", m.Name()))
	return q, true
}

func GetQueryToDelete(m Model, id string) (q SQLQueryToDelete, ok bool) {
	q = SQLQueryToDelete(fmt.Sprintf("DELETE FROM %s WHERE id IN (%s)", m.Name(), id))
	return q, ok
}

func GetParamFromRequest(req *http.Request, param string) (string, bool) {
	vars := mux.Vars(req)
	id, ok := vars[param]
	return id, ok
}

*/

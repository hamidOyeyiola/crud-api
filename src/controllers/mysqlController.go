package controller

import (
	"database/sql"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	model "github.com/hamidOyeyiola/crud-api/models"
)

type MySQLCRUDController struct {
	dataSource string
	db         *sql.DB
	conns      int
	mu         sync.Mutex
}

func (sc *MySQLCRUDController) Open() bool {
	sc.mu.Lock()
	if sc.db == nil {
		db, err := sql.Open("mysql", sc.dataSource)
		if err != nil {
			return false
		}
		db.SetConnMaxLifetime(3 * time.Minute)
		db.SetMaxOpenConns(10)
		db.SetMaxIdleConns(10)
		sc.db = db
	}
	sc.conns++
	sc.mu.Unlock()
	return true
}

func (sc *MySQLCRUDController) Close() bool {
	sc.mu.Lock()
	sc.conns--
	if sc.conns == 0 {
		err := sc.db.Close()
		if err != nil {
			return false
		}
		sc.db = nil
	}
	sc.mu.Unlock()
	return true
}

func NewMySQLCRUDController(datasrc string) *MySQLCRUDController {
	return &MySQLCRUDController{dataSource: datasrc}
}

func (sc *MySQLCRUDController) Create(q model.SQLQueryToInsert) (h *Header, b *Body, ok bool) {
	ok = sc.Open()
	if !ok {
		h, b = GetStatusFailedDependencyRes()
		return
	}
	defer sc.Close()
	//fmt.Println(q)
	res, err := sc.db.Exec(string(q))
	if err != nil {
		h, b = GetStatusNotModifiedRes()
		ok = false
		return
	}
	id, _ := res.LastInsertId()
	fmt.Printf("%d\n", id)
	h, b = GetStatusOKRes()
	ok = true
	return
}

func (sc *MySQLCRUDController) Retrieve(q model.SQLQueryToSelect, m model.Model) (h *Header, b *Body, ok bool) {
	ok = sc.Open()
	if !ok {
		h, b = GetStatusFailedDependencyRes()
		return
	}
	defer sc.Close()
	res, err := sc.db.Query(string(q))
	if err != nil {
		h, b = GetStatusFailedDependencyRes()
		ok = false
		return
	}
	defer res.Close()
	d, n := m.FromQueryResult(res)
	if n == 0 {
		h, b = GetStatusNotFoundRes()
		ok = false
		return
	}
	b = new(Body).
		AddContentType("application/json").
		AddContent(string(d))
	h, _ = GetStatusOKRes()
	return
}

func (sc *MySQLCRUDController) Update(q model.SQLQueryToUpdate) (h *Header, b *Body, ok bool) {
	ok = sc.Open()
	if !ok {
		h, b = GetStatusFailedDependencyRes()
		return
	}
	defer sc.Close()

	for j := 0; j < len(q.Stmts); j++ {
		stmt, err := sc.db.Prepare(string(q.Stmts[j]))
		if err != nil {
			h, b = GetStatusNotModifiedRes()
			ok = false
			return
		}
		_, err = stmt.Exec(q.Values[j], q.ID)
		if err != nil {
			h, b = GetStatusNotModifiedRes()
			ok = false
			return
		}
	}
	h, b = GetStatusOKRes()
	return
}

func (sc *MySQLCRUDController) Delete(q model.SQLQueryToDelete) (h *Header, b *Body, ok bool) {
	ok = sc.Open()
	if !ok {
		h, b = GetStatusFailedDependencyRes()
		return
	}
	defer sc.Close()
	_, err := sc.db.Exec(string(q))
	if err != nil {
		h, b = GetStatusBadRequestRes()
		ok = false
		return
	}
	h, b = GetStatusOKRes()
	return
}

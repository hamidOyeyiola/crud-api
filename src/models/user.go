package model

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/hamidOyeyiola/crud-api/interfaces"
	"github.com/hamidOyeyiola/crud-api/utils"
)

const (
	name   = "users"
	priKey = "email"
)

type User struct {
	FirstName string             `json:"firstname"`
	LastName  string             `json:"lastname"`
	Email     utils.EmailAddress `json:"email"`
	PhoneNo   string             `json:"phoneno"`
	Password  string             `json:"password"`
	MessageID string             `json:"-"`
	CreatedOn string             `json:"-"`
	UpdatedOn string             `json:"-"`
	ID        int                `json:"-"`
}

func (u *User) QueryToInsert(r io.Reader) (ins SQLQueryToInsert, ok bool) {
	err := json.NewDecoder(r).Decode(&u)
	if err != nil {
		return
	}
	ok = u.Email.IsValid()
	ins = SQLQueryToInsert(fmt.Sprintf("INSERT INTO users(firstname, lastname, email, phoneno, password, messageID, createdOn,updatedOn,id) VALUES ('%s','%s','%s','%s','%s','%s','%s','%s',%d)",
		u.FirstName, u.LastName, u.Email, u.PhoneNo, utils.EncryptPassword(u.Password), u.MessageID, utils.NewDate(), "", 0))
	return
}

func (u *User) QueryToUpdateWhere(r io.Reader, value string) (upt SQLQueryToUpdate, ok bool) {
	err := json.NewDecoder(r).Decode(&u)
	if ok = err == nil; !ok {
		return
	}
	upt.ID = value
	if u.FirstName != "" {
		upt.Stmts = append(upt.Stmts, string("UPDATE users SET firstname=? WHERE email=?"))
		upt.Values = append(upt.Values, u.FirstName)
	}
	if u.LastName != "" {
		upt.Stmts = append(upt.Stmts, string("UPDATE users SET lastname=? WHERE email=?"))
		upt.Values = append(upt.Values, u.LastName)
	}
	if u.PhoneNo != "" {
		upt.Stmts = append(upt.Stmts, string("UPDATE users SET phoneno=? WHERE email=?"))
		upt.Values = append(upt.Values, u.PhoneNo)
	}
	if u.Password != "" {
		upt.Stmts = append(upt.Stmts, string("UPDATE users SET password=? WHERE email=?"))
		upt.Values = append(upt.Values, utils.EncryptPassword(u.Password))
	}
	if u.MessageID != "" {
		upt.Stmts = append(upt.Stmts, string("UPDATE users SET messageID=? WHERE email=?"))
		upt.Values = append(upt.Values, string(u.MessageID))
	}
	upt.Stmts = append(upt.Stmts, string("UPDATE users SET updatedOn=? WHERE email=?"))
	date := fmt.Sprintf("%s", utils.NewDate())
	upt.Values = append(upt.Values, date)
	return
}

func (u *User) FromQueryResult(s interfaces.Iterator) (JSONObject, int) {
	v := []User{}
	var n int
	var id int
	for n = 0; s.Next(); n++ {
		err := s.Scan(&u.FirstName, &u.LastName, &u.Email, &u.PhoneNo, &u.Password, &u.MessageID, &u.CreatedOn, &u.UpdatedOn, &id)
		if err != nil {
			break
		}
		v = append(v, *u)
	}
	o, _ := json.MarshalIndent(v, "", "  ")
	return JSONObject(o), n
}

func (u *User) QueryToSelectWhere(value string) (q SQLQueryToSelect, ok bool) {
	q = SQLQueryToSelect("SELECT * FROM " + name)
	q = q + SQLQueryToSelect(fmt.Sprintf(" WHERE %s = '%s'", priKey, value))
	return q, true
}

func (u *User) QueryToDeleteWhere(key string) (q SQLQueryToDelete, ok bool) {
	q = SQLQueryToDelete(fmt.Sprintf("DELETE FROM %s WHERE %s = '%s'", name, priKey, key))
	return q, ok
}

func (u *User) QueryToSelectAll() (q SQLQueryToSelect, ok bool) {
	q = SQLQueryToSelect(fmt.Sprintf("SELECT * FROM users"))
	return q, true
}

func (u *User) QueryToDeleteAll() (q SQLQueryToDelete, ok bool) {
	q = SQLQueryToDelete(fmt.Sprintf("DELETE FROM users"))
	return q, true
}

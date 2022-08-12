package controller

import (
	model "github.com/hamidOyeyiola/crud-api/models"
)

type Creater interface {
	Create(model.SQLQueryToInsert) (*Header, *Body, bool)
}

type Retriever interface {
	Retrieve(model.SQLQueryToSelect, model.Model) (*Header, *Body, bool)
}

type Updater interface {
	Update(model.SQLQueryToUpdate) (*Header, *Body, bool)
}

type Deleter interface {
	Delete(model.SQLQueryToDelete) (*Header, *Body, bool)
}

type CreateRetrieveUpdateDeleter interface {
	Creater
	Retriever
	Updater
	Deleter
}

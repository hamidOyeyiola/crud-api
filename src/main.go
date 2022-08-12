package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	model "github.com/hamidOyeyiola/crud-api/models"

	"github.com/hamidOyeyiola/crud-api/api"
	controller "github.com/hamidOyeyiola/crud-api/controllers"
)

const (
	dataSource = "hamid:kolajoke2055@tcp(localhost:3306)/registrationandlogin"
)

func main() {
	rt := mux.NewRouter()
	c := controller.NewMySQLCRUDController(dataSource)
	r := api.Read{
		Model:     &model.User{},
		Retriever: c,
	}
	n := api.New{
		Model:   &model.User{},
		Creater: c,
	}
	d := api.Remove{
		Model:   &model.User{},
		Deleter: c,
	}
	u := api.Change{
		Model:   &model.User{},
		Updater: c,
	}
	rt.HandleFunc("/read/{id}", r.Read)
	rt.HandleFunc("/readall", r.ReadAll)
	rt.HandleFunc("/new", n.New)
	rt.HandleFunc("/remove/{id}", d.Remove)
	rt.HandleFunc("/removeall", d.RemoveAll)
	rt.HandleFunc("/change/{id}", u.Change)

	go func() {
		err := http.ListenAndServe("localhost:8000", rt)
		if err != nil {
			log.Println(err)
		}
	}()

	fmt.Println("Server is now ready to take your orders!")
	e := make(chan os.Signal, 1)
	signal.Notify(e, os.Interrupt)
	<-e
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Println("Shutting Down!")
	os.Exit(1)
}

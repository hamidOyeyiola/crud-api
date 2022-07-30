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

func main() {
	rt := mux.NewRouter()
	c := controller.NewMySQLCRUDController("hamid:@tcp(localhost:3306)/userinfo")
	api.MakeCreaterAPI(rt, c, "/api/create", "", "", model.User{}, nil)
	r := controller.NewMySQLCRUDController("hamid:@tcp(localhost:3306)/userinfo")
	api.MakeCreaterAPI(rt, r, "/api/retrieve", "id", "", model.User{}, nil)
	u := controller.NewMySQLCRUDController("hamid:@tcp(localhost:3306)/userinfo")
	api.MakeUpdaterAPI(rt, u, "/api/update", "id", "", model.User{}, nil)
	d := controller.NewMySQLCRUDController("hamid:@tcp(localhost:3306)/userinfo")
	api.MakeRetrieverAPI(rt, d, "/api/delete", "id", "", model.User{}, nil)

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

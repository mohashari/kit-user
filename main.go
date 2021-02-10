package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/mohashari/kit-user/endpoint"
	"github.com/mohashari/kit-user/repo"
	"github.com/mohashari/kit-user/service"
	"github.com/mohashari/kit-user/trasport"
)

func main() {
	const (
		host     = "localhot"
		port     = 5432
		user     = "muklis"
		password = "1234"
		dbname   = "test"
	)
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	repository := repo.NewUserRepo(db)
	svc := service.NewUserService(repository)
	userEndpoint := endpoint.NewUserEndpoint(svc)
	h := trasport.MakeHandler(userEndpoint)

	http.ListenAndServe(":8000", h)

}

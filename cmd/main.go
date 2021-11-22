package main

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4/middleware"

	"github.com/detohm/todo-app-study/config"
	"github.com/detohm/todo-app-study/handler"
	"github.com/detohm/todo-app-study/route"
	todoRepo "github.com/detohm/todo-app-study/todo/repository"
	todoService "github.com/detohm/todo-app-study/todo/service"
	"github.com/labstack/echo/v4"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	conf := config.LoadConfig()

	db, err := sql.Open("mysql", conf.MySQLConn)
	if err != nil {
		log.Panic("error: database initialization")
	}
	defer db.Close()

	repo := todoRepo.NewToDoRepository(db)
	service := todoService.NewToDoService(repo)

	handler := handler.NewHandler(service)
	e := echo.New()

	route := route.NewRouteV1(e, handler)
	route.Bind()

	e.Use(middleware.Logger())
	e.Start(":5000")

	// fs := http.FileServer(http.Dir("./frontend/build"))
	// http.Handle("/", fs)
	// err := http.ListenAndServe(":5500", nil)
	// if err != nil {
	// 	log.Fatal(err)
	// }

}

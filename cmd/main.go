package main

import (
	"context"
	"database/sql"
	"log"
	"os"
	"os/signal"
	"time"

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

	e.Static("/", "./frontend/build")

	e.Use(middleware.Logger())

	// start server in another go routine
	go func() {
		err := e.Start(":" + conf.Port)
		if err != nil {
			log.Println("couldn't initialize the server")
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	<-sig

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = e.Shutdown(ctx)
	if err != nil {
		log.Panic(err)
	}
}

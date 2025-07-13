package main

import (
	"backend/internal/config"
	"backend/internal/config/db"
	"backend/internal/config/http"
	scheduler "backend/internal/layers/business"
	asgn "backend/internal/layers/database/challenge/assignmnt"
	inst "backend/internal/layers/database/challenge/instance"
	show "backend/internal/layers/database/challenge/showcase"
	tmpl "backend/internal/layers/database/challenge/template"
	user "backend/internal/layers/database/coolusers"
	"context"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	config.Load()
	ctx := context.Background()

	dbConn := db.Connect(ctx, config.AppConf.DbServer)
	if dbConn == nil {
		log.Fatal("Failed to connect to DB")
	}
	defer dbConn.Close(ctx)

	userRepo := user.NewPgRepo(dbConn)
	tmplRepo := tmpl.NewPgRepo(dbConn)
	instRepo := inst.NewPgRepo(dbConn)
	showRepo := show.NewPgRepo(dbConn)
	asgnRepo := asgn.NewPgRepo(dbConn)

	go func() {
		time.Sleep(5 * time.Second)
		chgeScheduler := scheduler.NewChgeScheduler(tmplRepo, instRepo)
		chgeScheduler.Start(ctx)
	}()

	router := http.GetRouter(userRepo, showRepo, asgnRepo)
	http.StartServer(config.AppConf.HTTPServer, router)
}

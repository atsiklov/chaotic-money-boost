package http

import (
	asgn "backend/internal/layers/database/challenge/assignmnt"
	show "backend/internal/layers/database/challenge/showcase"
	user "backend/internal/layers/database/coolusers"
	"backend/internal/layers/transport/handlers"

	"github.com/gin-gonic/gin"
)

func GetRouter(
	userRepo user.Repository,
	showRepo show.Repository,
	asgnRepo asgn.Repository,
) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	userHandler := handlers.NewUserHandler(userRepo)
	showHandler := handlers.NewShowHandler(showRepo)
	asgnHandler := handlers.NewAsgnHandler(asgnRepo)

	router.GET("/users/:id", userHandler.GetUser)
	router.POST("/users", userHandler.CreateUser)

	router.GET("/challenges/", showHandler.GetShowcases)
	router.GET("/challenges/:id", showHandler.GetShowcase)
	router.POST("/challenges/:id/assign", asgnHandler.CreateAssignment)
	router.PATCH("/challenges/:id/submit", asgnHandler.UpdateAssignment)

	return router
}

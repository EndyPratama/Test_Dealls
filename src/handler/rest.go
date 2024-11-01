package handler

import (
	"fmt"
	"test_dealls/src/business/usecase"
	"test_dealls/src/middleware"
	"test_dealls/src/utils/config"
	"test_dealls/src/utils/configreader"
	"test_dealls/src/utils/log"

	"github.com/gin-gonic/gin"
)

type REST interface {
	Run()
}

type rest struct {
	http         *gin.Engine
	conf         config.Application
	configreader configreader.Interface
	log          log.Interface
	uc           *usecase.Usecase
}

func Init(conf config.Application, configreader configreader.Interface, log log.Interface, uc *usecase.Usecase) REST {
	r := &rest{}

	httpServer := gin.New()

	r = &rest{
		conf:         conf,
		configreader: configreader,
		log:          log,
		http:         httpServer,
		uc:           uc,
	}

	// Set Middleware
	r.http.Use(middleware.DeallsMiddleware(uc.User, conf, log))

	r.Register()

	return r
}

func (r *rest) Run() {
	if r.conf.Gin.Port != "" {
		r.http.Run(fmt.Sprintf(":%s", r.conf.Gin.Port))
	} else {
		r.http.Run(":8080")
	}
}

func (r *rest) Register() {
	// server health and testing purpose
	r.http.GET("/ping", r.Ping)

	// Done
	r.http.POST("/login", r.Login)
	r.http.POST("/register", r.RegisterUser)

	// Done
	profileGroup := r.http.Group("/profile")
	profileGroup.Use(middleware.CheckToken())
	{
		profileGroup.GET("/searchPeople", r.SearchPeople)
		profileGroup.GET("/detail", r.detailProfile)
		profileGroup.POST("/create", r.createProfile)
		profileGroup.POST("/update", r.UpdateProfile)
		profileGroup.POST("/upgrade", r.UpgradeProfile)
	}

	// Done
	photoGroup := r.http.Group("/photo")
	photoGroup.Use(middleware.CheckToken())
	{
		photoGroup.POST("/add", r.createPhoto)
		photoGroup.POST("/update", r.UpdatePhoto)
	}

	likesGroup := r.http.Group("/likes")
	likesGroup.Use(middleware.CheckToken())
	{
		likesGroup.GET("/", r.listLikes)
		likesGroup.POST("/approve", r.approve)
		likesGroup.POST("/skip", r.skip)
		likesGroup.POST("/delete", r.DeleteLikes)
	}
}

func (r *rest) Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "PONG!"})
}

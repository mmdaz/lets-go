package server


import (
	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"lets_go/gorm_practice/pkg/log"
)

type Server struct {
	isReady chan bool
}

func New() *Server {
	return &Server{
		isReady: make(chan bool),
	}
}

func (s *Server) Run() bool {
	go s.start()
	return <-s.isReady
}

func (s *Server) start() {
	e := echo.New()
	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover()) // Recover is recover from panics anywhere in the chain.
	//CORS
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	e.POST("/add_user", func(context echo.Context) error {
		fmt.Print(context)
		return echo.ErrBadRequest
	})
	log.Logger.Info(e.Start(":1323"))
	s.isReady <- true
}

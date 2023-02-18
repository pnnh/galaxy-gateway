package server

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type IResource interface {
	RegisterRouter(router *gin.Engine, name string)
}

type WebServer struct {
	router    *gin.Engine
	resources map[string]IResource
}

func NewWebServer() (*WebServer, error) {
	router := gin.Default()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	server := &WebServer{
		router:    router,
		resources: make(map[string]IResource)}

	return server, nil
}

func (s *WebServer) Init() error { 

	s.router.GET("/account/restful", accountBackendHandler)
	s.router.GET("/account", accountHandler)
	s.router.GET("/restful", mainBusinessHandler)
	s.router.Use(func(ctx *gin.Context) {
		if strings.HasPrefix(ctx.Request.URL.Path, "/account/restful") {
			accountBackendHandler(ctx)
		} else if strings.HasPrefix(ctx.Request.URL.Path, "/account") {
			accountHandler(ctx)
		} else if strings.HasPrefix(ctx.Request.URL.Path, "/restful") {
			mainBusinessHandler(ctx)
		} else  {
			mainSiteHandler(ctx)
		} 
	})

	return nil
}

func (s *WebServer) Start() error {
	port := os.Getenv("port")
	if len(port) < 1 {
		port = "8080"
	}
	var handler http.Handler = s

	serv := &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	if err := serv.ListenAndServe(); err != nil {
		return fmt.Errorf("服务出错停止: %w", err)
	}
	return nil
}

func (s *WebServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
 
	s.router.ServeHTTP(w, r)
}

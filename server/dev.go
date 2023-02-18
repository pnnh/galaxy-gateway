package server

import ( 
	"net/http/httputil"
	"net/url" 

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func mainSiteHandler(gctx *gin.Context) { 
	rs1 := "http://127.0.0.1:8101"
	targetUrl, err := url.Parse(rs1)
	if err != nil {
		logrus.Fatalln("mainSiteHandler: %w", err)
		return
	}
	httputil.NewSingleHostReverseProxy(targetUrl).ServeHTTP(gctx.Writer, gctx.Request)
}

func mainBusinessHandler(gctx *gin.Context) { 
	rs1 := "http://127.0.0.1:8102"
	targetUrl, err := url.Parse(rs1)
	if err != nil {
		logrus.Fatalln("mainSiteHandler: %w", err)
		return
	}
	httputil.NewSingleHostReverseProxy(targetUrl).ServeHTTP(gctx.Writer, gctx.Request)
}

func accountHandler(gctx *gin.Context) { 
	rs1 := "http://127.0.0.1:8103"
	targetUrl, err := url.Parse(rs1)
	if err != nil {
		logrus.Fatalln("mainSiteHandler: %w", err)
		return
	}
	httputil.NewSingleHostReverseProxy(targetUrl).ServeHTTP(gctx.Writer, gctx.Request)
}

func accountBackendHandler(gctx *gin.Context) { 
	rs1 := "http://127.0.0.1:8104"
	targetUrl, err := url.Parse(rs1)
	if err != nil {
		logrus.Fatalln("mainSiteHandler: %w", err)
		return
	}
	httputil.NewSingleHostReverseProxy(targetUrl).ServeHTTP(gctx.Writer, gctx.Request)
}


func debugStaticWebHandler(gctx *gin.Context) {
	// target := "localhost:8080"
	// //devUrl := fmt.Sprintf("http://localhost:3000%s", realPath)
	// proxy := &httputil.ReverseProxy{Director: func(req *http.Request) {
	// 	req.URL.Scheme = "http"
	// 	req.URL.Host = target
	// 	req.URL.Path = r.URL.Path //"/svc/css/index.scss"
	// 	//r.Host = target
	// }}

	// proxy.ServeHTTP(w, r)

	rs1 := "http://127.0.0.1:3500"
	targetUrl, err := url.Parse(rs1)
	if err != nil {
		logrus.Fatalln("debugStaticWebHandler: %w", err)
		return
	}
	httputil.NewSingleHostReverseProxy(targetUrl).ServeHTTP(gctx.Writer, gctx.Request)

}

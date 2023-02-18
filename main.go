package main

import ( 
	"github.com/pnnh/quantum-go/config"
	"github.com/sirupsen/logrus"
)

func main() {
	config.InitApp() 

	app := NewApplication()
 
	if err := app.Init(); err != nil {
		logrus.Fatalln("应用程序初始化出错: %w", err)
	}
	if err := app.Start(); err != nil {
		logrus.Fatalln("应用程序执行出错: %w", err)
	}
}

package http

import (
	koanfx "berlin/utils/koanfx"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"sync"
)

func SetupRouter() *gin.Engine {
	app := gin.New()
	app.NoRoute(func(c *gin.Context) {
		log.Errorf("No route found")
	})
	return app
}

func Run(wg *sync.WaitGroup) {

	defer wg.Done()
	host := koanfx.K.String("host")
	port := koanfx.K.String("port")
	app := SetupRouter()
	err := app.Run(host + ":" + port)
	if err != nil {
		panic(err)
	}
}

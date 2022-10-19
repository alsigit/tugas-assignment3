package webserver

import (
	"assignment3/webserver/configs"
	"assignment3/webserver/router"
)

func Start() {
	// engine := gin.Default()
	router.CreateRouter().Run(configs.PORT)
}

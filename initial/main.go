package main

import (
	"os"

	klog "k8s.io/klog/v2"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("INITIAL_PORT")
	if port == "" {
		port = "9876"
		klog.Warning("no env variable INITIAL_PORT. Using default port " + port)
	} else {
		klog.Info("port set by env variable to " + port)
	}
	router := gin.Default()
	router.POST("/calc", handleCalc)
	router.Run("0.0.0.0:" + port)
}

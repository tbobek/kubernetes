package main

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	klog "k8s.io/klog/v2"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Id         string
	Calls      int `json:"calls" binding:"required"`      // number of calls
	Iterations int `json:"iterations" binding:"required"` // number of iterations
	WaitTime   int `json:"wait_time"`                     // number of milliseconds to wait
}

func main() {
	port := os.Getenv("WORKER_PORT")
	if port == "" {
		klog.Info("Env variable WORKER_PORT not set, using default port 8765")
		port = "8765"
	} else {
		klog.Info("Port set by env variable to ", port)
	}
	//flag.StringVar(&port, "p", "8765", "set port for service")
	//flag.Parse()
	r := gin.Default()
	r.POST("/calc", func(c *gin.Context) {
		var rq Request
		err := c.ShouldBindBodyWithJSON(&rq)
		if err == nil {
			klog.Infof("handling call with id %s", rq.Id)
			for i := 0; i < rq.Iterations; i++ {
				time.Sleep(time.Duration(rq.WaitTime) * time.Millisecond)
				klog.Infof("iteration %d of %d finished", i+1, rq.Iterations)
			}
			c.JSON(http.StatusOK, gin.H{"result": rand.Intn(100)})
		} else {
			klog.Warning("can't decode body, aborting")
			c.JSON(http.StatusInternalServerError, gin.H{"result": -1})
		}
	})
	r.Run("0.0.0.0:" + port)
}

package main

import (
	"eth-agent/modules/agent"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/DeanThompson/ginpprof"
	nice "github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"eth-agent/config"
)

// AgentError represent error type in each opearative error condition.
type AgentError struct {
	// Type:0: System Error ,Type: 1; Client Error
	ErrorType        int
	ErrorDescription string
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New() // gin.Default() installs gin.Recovery() so use gin.New() instead
	router.Use(nice.Recovery(recoveryHandler))

	f, _ := os.Create(config.SysConf.EthProxy.LogFile)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// middleware.VerifyToken is for SSO service
	// router.Use(middleware.VerifyToken())

	router.POST("/agent", agent.Redirect)

	ginpprof.Wrap(router)

	router.Run(":" + config.SysConf.EthProxy.Port)
}

func recoveryHandler(c *gin.Context, err interface{}) {
	//TODO: return eth offcial json-rpc format
	switch err.(type) {
	case *logrus.Entry:
		logrusEntryStruct := err.(*logrus.Entry)
		logrusMessage := logrusEntryStruct.Message
		panicType := logrusMessage[1:2]
		panicMessage := logrusMessage[3 : len(logrusMessage)-1]
		fmt.Println(panicMessage)
		switch panicType {
		case "0":
			c.JSON(http.StatusInternalServerError, panicMessage)
		case "1":
			c.JSON(http.StatusInternalServerError, "500 Internal Server Error")
		}

	default:
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "Unknown Error")
	}
}

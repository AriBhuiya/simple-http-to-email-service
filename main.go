package main

import (
	"emailMS/EmailService"
	"fmt"
	"github.com/gin-gonic/gin"
	"os"
)



func main()  {
	var err error
	EmailService.GlobalConfig, err =EmailService.ReadConfig()
	if err!=nil{
		os.Exit(-1)
		fmt.Println("Service exiting because of lack of configuration")
	}

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)

	router.GET("/", EmailService.Check)
	router.POST("send/", EmailService.SendMailHandler)
	err1 := router.Run(":7000")
	if err1 != nil {
		fmt.Fprintf(os.Stderr, "Unable to run router: %v\n", err1)
	}
}
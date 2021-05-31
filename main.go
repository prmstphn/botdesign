package main

import (
	"fmt"
	//"inittest/stringutil"
	"net/http"

	//	"io/ioutil"
	//	"encoding/base64"
	//	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func PostMethod(c *gin.Context) {
	fmt.Println("\napi.go 'PostMethod' called")
	message := "PostMethod called"
	c.JSON(http.StatusOK, message)
}

func GetMethod(c *gin.Context) {
	fmt.Println("\napi.go 'GetMethod' called")
	message := "GetMethod called"
	c.JSON(http.StatusOK, message)
}
func main() {
	//fmt.Println(stringutil.Reverse("!selpmaxe ,oGolleH"))
	logrus.Info("this is to check debug")
	router := gin.Default()

	router.POST("/", PostMethod)
	router.GET("/", GetMethod)

	listenPort := "4000"
	// Listen and Server on the LocalHost:Port
	router.Run(":" + listenPort)

}

package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/golang/glog"
)

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Passwd   string `form:"passwd"   json:"passwd" bdinding:"required"`
	Age      int    `form:"age"      json:"age"`
}

func main() {
	flag.Parse()
	defer glog.Flush()
	router := gin.Default()
	router.POST("/echo", echo)
	router.POST("/user", userJson)

	certFile := "/home/ansen/keys/gv_svc.crt"
	keyFile := "/home/ansen/keys/gv_svc.key"
	// enable http
	// router.Run(":8080")
	// enable https
	glog.Infof("Using self-signed cert (%s, %s)", certFile, keyFile)
	glog.Infof("another line of log\n")
	router.RunTLS(":443", certFile, keyFile)
}

func echo(c *gin.Context) {
	buf := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	fmt.Println(string(buf[0:n]))
	//resp := map[string]string{"hello": "world"}
	//c.JSON(http.StatusOK, resp)
	c.String(http.StatusOK, string(buf[0:n]))
	/*post_gwid := c.PostForm("name")
	fmt.Println(post_gwid)*/
}

func userJson(c *gin.Context) {
	var user User
	var err error
	contentType := c.Request.Header.Get("Content-Type")
	switch contentType {
	case "application/json":
		err = c.BindJSON(&user)
	case "application/x-www-form-urlencoded":
		err = c.BindWith(&user, binding.Form)
	}
	if err != nil {
		fmt.Println(err.Error())
		glog.Error(err.Error())
	}

	c.JSON(http.StatusOK,
		gin.H{"user": user.Username, "passwd": user.Passwd, "age": user.Age})

}

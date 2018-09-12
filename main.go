package main
 
import (
	"fmt"
        "flag"
	"net/http"

        "github.com/golang/glog"
	"github.com/gin-gonic/gin"
)
 
func main() {
        flag.Parse()
        defer glog.Flush()
	router := gin.Default()
	router.POST("/echo", echo)
	certFile := "/home/ansen/keys/gv_svc.crt"
	keyFile  := "/home/ansen/keys/gv_svc.key"
        // run http
	// router.Run(":8080")
        // run https
        glog.Infof("Using self-signed cert (%s, %s)", certFile, keyFile)
        glog.Infof("another line of log\n")
        router.RunTLS(":443", certFile, keyFile)
}     
 
func echo(c *gin.Context) {
	buf  := make([]byte, 1024)
	n, _ := c.Request.Body.Read(buf)
	fmt.Println(string(buf[0:n]))
	//resp := map[string]string{"hello": "world"}
	//c.JSON(http.StatusOK, resp)
        c.String(http.StatusOK, string(buf[0:n]))
	/*post_gwid := c.PostForm("name")
	fmt.Println(post_gwid)*/
}

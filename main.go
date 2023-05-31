package main
import (
	"github.com/gin-gonic/gin"
    "net/http"
)

func main(){
	server := gin.Default()
	server.Any("/",RootWeb)
	server.Run(":8010")
}

func RootWeb(context *gin.Context){
	context.String(http.StatusOK, "hello world")
}
package routers

import (
	"baoxiu/methods"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.LoadHTMLGlob("templates/**/*")
	r.StaticFS("/static", http.Dir("./static"))
	record := methods.Newrecord()
	baoxiu := methods.NewBx()

	detail := methods.NewDetail()
	apiv1 := r.Group("/")
	//apiv1.StaticFS("/static", http.Dir("/static"))

	{
		apiv1.GET("/", record.Index) ///index

		apiv1.GET("/detail/:id", detail.Index)

		/////////////
		apiv1.GET("/baoxiu", baoxiu.Index)
		apiv1.POST("/bxpost", baoxiu.CreateBx)
		apiv1.DELETE("/delbx/:id", baoxiu.DelBx)
		apiv1.GET("/baoxiu/success", baoxiu.Success)

	}

	//apiv2 := r.Group("/admin")
	return r
}

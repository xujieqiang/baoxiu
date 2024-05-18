package main

import (
	"baoxiu/routers"
	"fmt"
)

func main() {
	fmt.Println("网络测试")
	// r := gin.Default()
	// r.LoadHTMLGlob("templates/*")
	// r.Static("/static", "./static")
	// r.GET("/", Index)
	r := routers.NewRouter()

	r.Run(":8080")
	//s.ListenAndServe()
}

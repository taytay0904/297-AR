package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zijianguan0204/297AR/controller"
)

//http://localhost:8080/assets/.jpg/road.jpg
//http://localhost:8080/assets/..mod/pa_drone
func main() {

	r := gin.Default()
	c := controller.NewController()
	r.GET("/", c.GetModelListFromString)
	r.Static("/assets", "E:/VScode/go/go_workspace/go/src/github.com/zijianguan0204/297AR/assets") //need to be changed in server
	v1 := r.Group("/api/v1")
	{
		model := v1.Group("/model")
		{
			// model.POST("getModelList", c.GetModelList)
			// model.GET("getModelList", c.GetModelList)
			model.GET("getModelList", c.GetModelListFromString)
		}

		detail := v1.Group("/detail")
		{
			// detail.POST("getDetail", c.GetDetail)
			detail.GET("getDetailFromString/:modelID", c.GetDetailFromString) //get info from String
		}
	}

	r.Run()
}

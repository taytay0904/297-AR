package main

import (
	"github.com/gin-gonic/gin"
	"github.com/zijianguan0204/297AR/controller"
)

func main() {

	r := gin.Default()
	c := controller.NewController()
	r.Static("/assets", "./assets")
	v1 := r.Group("/api/v1")
	{
		model := v1.Group("/model")
		{
			model.POST("getModelList", c.GetModelList)
			model.StaticFile("/getModel/road.jpg", "./assets/.jpg/road.jpg")
		}

		detail := v1.Group("/company")
		{
			detail.POST("getDetail", c.GetDetail)
		}
	}

	r.Run()
}

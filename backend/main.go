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


// 	http.HandleFunc("/", dog)
// 	http.HandleFunc("/model", dogPic)
// 	http.ListenAndServe(":8080", nil)
// }

// func dog(w http.ResponseWriter, req *http.Request) {
// 	w.Header().Set("Content-Type", "text/html; charset=utf-8")
// 	io.WriteString(w, `<img src = "/toby.jpg">`)
// }

// func dogPic(w http.ResponseWriter, req *http.Request) {
// 	f, err := os.Open("pa_warrior")
// 	if err != nil {
// 		http.Error(w, "file not found", 404)
// 		return

	}

	r.Run()
}

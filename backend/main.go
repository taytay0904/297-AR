package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	// fmt.Println("Hello World")

	// r := gin.Default()
	// // r.GET("/", getModel)
	// r.POST("api/addModel", addModel)
	// r.POST("api/getModel", getModel)
	// r.Static("/assets", "./assets")
	// r.StaticFile("/favicon.png", "./assets/favicon.png")
	// r.Use(static.Serve("/", static.LocalFile("./frontend", true)))

	// r.Run()

	http.HandleFunc("/", dog)
	http.HandleFunc("/toby.jpg", dogPic)
	http.ListenAndServe(":8080", nil)
}

func dog(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src = "/toby.jpg">`)
}

func dogPic(w http.ResponseWriter, req *http.Request) {
	f, err := os.Open("toby.jpg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer f.Close()

	fi, err := f.Stat()
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	http.ServeContent(w, req, f.Name(), fi.ModTime(), f)
}

//Company sighup function
// func addModel(c *gin.Context) {
// 	//getting data from request
// 	body := c.Request.Body
// 	value, err := ioutil.ReadAll(body)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	type Model struct {
// 		Name            string `json:"name"`
// 		ICON_Path       string `json:"ICON_Path"`
// 		Model_File_Path string `json:"model_File_Path"`
// 		Detail_Info_ID  string `json:"detail_Info_ID"`
// 		Sound_Path      string `json:"sound_Path"`
// 	}

// 	var model Model
// 	er := json.Unmarshal([]byte(value), &model)
// 	if er != nil {
// 		fmt.Println(err.Error())
// 	}
// 	e := reflect.ValueOf(&model).Elem()
// 	name := fmt.Sprint(e.Field(0).Interface())
// 	icon_path := fmt.Sprint(e.Field(1).Interface())
// 	model_file_path := fmt.Sprint(e.Field(2).Interface())
// 	detail_info_ID := fmt.Sprint(e.Field(3).Interface())
// 	sound_path := fmt.Sprint(e.Field(4).Interface())

// 	// connect to the db
// 	db, err := connectDB(c)
// 	if err != nil {
// 		fmt.Println("DB error")
// 		c.JSON(500, gin.H{
// 			"message": "DB connection problem",
// 		})
// 		panic(err.Error())
// 	}

// 	//insert to db
// 	insert, err := db.Query("INSERT INTO model VALUES(DEFAULT,?,?,?,?,?,DEFAULT)", name, icon_path, model_file_path, detail_info_ID, sound_path)

// 	if err != nil {
// 		fmt.Println("model fail")
// 		if strings.Contains(err.Error(), "Access denied") {
// 			c.JSON(500, gin.H{
// 				"message": "DB access error, username or password is wrong",
// 			})
// 			panic(err.Error())
// 		}
// 		c.JSON(404, gin.H{
// 			"message": "model fail",
// 		})
// 		panic(err.Error())
// 	} else {
// 		c.JSON(201, gin.H{
// 			"message": "This record has been assigned or created",
// 		})
// 	}
// 	defer db.Close()
// 	defer insert.Close()
// }

// func getModel(c *gin.Context) {

// 	//make struct and array
// 	type Model struct {
// 		id              int
// 		name            string
// 		icon_path       string
// 		model_file_path string
// 		detail_info_id  string
// 		sound_path      string
// 		time_stamp      string
// 	}
// 	var models [5]Model
// 	index := 0
// 	// connect to the db
// 	db, err := connectDB(c)
// 	if err != nil {
// 		fmt.Println("DB error")
// 		c.JSON(500, gin.H{
// 			"message": "DB connection problem",
// 		})
// 		panic(err.Error())
// 	}

// 	defer db.Close()

// 	var (
// 		id              int
// 		name            string
// 		icon_path       string
// 		model_file_path string
// 		detail_info_id  string
// 		sound_path      string
// 		time_stamp      string
// 	)
// 	rows, err := db.Query("SELECT * from model")
// 	if err != nil {
// 		fmt.Println("Login error")
// 		if strings.Contains(err.Error(), "Access denied") {
// 			c.JSON(500, gin.H{
// 				"message": "DB access error, username or password is wrong",
// 			})
// 			panic(err.Error())
// 		}
// 		c.JSON(404, gin.H{
// 			"message": "model does not exist",
// 		})
// 		panic(err.Error())
// 	} else {
// 		for rows.Next() {
// 			err := rows.Scan(&id, &name, &icon_path, &model_file_path, &detail_info_id, &sound_path, &time_stamp)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			model := Model{id: id, name: name, icon_path: icon_path, model_file_path: model_file_path, detail_info_id: detail_info_id, sound_path: sound_path, time_stamp: time_stamp}
// 			models[index] = model
// 			fmt.Println(models[index].id)
// 			fmt.Println(models[index].name)
// 			fmt.Println(models[index].icon_path)
// 			fmt.Println(models[index].model_file_path)
// 			fmt.Println(models[index].detail_info_id)
// 			fmt.Println(models[index].sound_path)
// 			index = index + 1
// 		}

// 		modelJson, err := json.Marshal(models)
// 		// fmt.Fprintf(os.Stdout, "%s", models)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 		c.JSON(http.StatusOK, gin.H{
// 			"code":    http.StatusOK,
// 			"message": string(modelJson), // cast it to string before showing
// 		})
// 	}
// 	defer rows.Close()
// }

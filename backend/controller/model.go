package controller

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zijianguan0204/297AR/model"
)

func (ct *Controller) GetModelList(c *gin.Context) {

	// connect to the db
	db, err := connectDB(c)
	if db == nil {
		message := model.Message{
			Message: "DB has something wrong",
		}
		c.JSON(500, message)
		panic(err.Error())
	}

	var (
		id            string
		name          string
		modelPath     string
		detailPath    string
		soundPath     string
		animationPath string
	)
	modelList := []model.ModelOutput{}
	rows, err := db.Query("SELECT * from mode")
	if err != nil {
		if strings.Contains(err.Error(), "Access denied") {
			message := model.Message{
				Message: "DB access error, username or password is wrong",
			}
			c.JSON(500, message)
			panic(err.Error())
		}
		message := model.Message{
			Message: "Query statements has something wrong",
		}
		c.JSON(500, message)
		panic(err.Error())
	} else {
		for rows.Next() {
			err := rows.Scan(&id, &name, &modelPath, &detailPath, &soundPath, &animationPath)
			if err != nil {
				log.Fatal(err)
			}
		}
		if id == "0" {
			message := model.Message{
				Message: "Model does not exist",
			}
			c.JSON(404, message)
		} else {
			model := model.ModelOutput{
				ID:            id,
				Name:          name,
				ModelPath:     modelPath,
				DetailPath:    detailPath,
				SoundPath:     soundPath,
				AnimationPath: animationPath,
			}
			modelList = append(modelList, model)
		}
	}
	c.JSON(200, modelList)
	defer rows.Close()
	defer db.Close()

}

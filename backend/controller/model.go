package controller

import (
	"log"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zijianguan0204/297AR/model"
)

func (ct *Controller) GetModelListFromString(c *gin.Context) {

	jsonString := "[{\"id\":\"1\",\"name\":\"A103\",\"modelPath\":\"/assets/.mod/pa_drone\",\"detailPath\":\"fake_detail_path\",\"soundPath\":\"fake_sound_path\",\"animationPath\":\"fake_am_path\"},{\"id\":\"2\",\"name\":\"A404\",\"modelPath\":\"/assets/.mod/pa_warrior\",\"detailPath\":\"fake_detail_path2\",\"soundPath\":\"fake_sound_path2\",\"animationPath\":\"fake_am_path2\"},{\"id\":\"3\",\"name\":\"E45\",\"modelPath\":\"/assets/.mod/E45\",\"detailPath\":\"fake_detail_path\",\"soundPath\":\"fake_sound_path\",\"animationPath\":\"fake_am_path\"},{\"id\":\"4\",\"name\":\"A103\",\"modelPath\":\"/assets/.mod/B25\",\"detailPath\":\"fake_detail_path\",\"soundPath\":\"fake_sound_path\",\"animationPath\":\"fake_am_path\"},{\"id\":\"5\",\"name\":\"B25\",\"modelPath\":\"/assets/.mod/C73\",\"detailPath\":\"fake_detail_path\",\"soundPath\":\"fake_sound_path\",\"animationPath\":\"fake_am_path\"},{\"id\":\"6\",\"name\":\"C73\",\"modelPath\":\"/assets/.mod/A103\",\"detailPath\":\"fake_detail_path\",\"soundPath\":\"fake_sound_path\",\"animationPath\":\"fake_am_path\"}]"
	c.String(200, jsonString)
}

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
	rows, err := db.Query("SELECT * from model")
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
			if id == "" {
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

	}
	c.JSON(200, modelList)
	defer rows.Close()
	defer db.Close()

}

package controller

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zijianguan0204/297AR/model"
)

func homePage(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello World",
	})
}

func (ct *Controller) GetDetail(c *gin.Context) {

	body := c.Request.Body
	value, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println(err.Error())
	}

	var input model.DetailInput
	er := json.Unmarshal([]byte(value), &input)
	if er != nil {
		fmt.Println(err.Error())
	}
	e := reflect.ValueOf(&input).Elem()
	modelID := fmt.Sprint(e.Field(0).Interface())

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
		weight   string
		height   string
		built    string
		designer string
	)
	rows, err := db.Query("SELECT weight, height, built, designer from detail where id = ?", modelID)
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
			err := rows.Scan(&weight, &height, &built, &designer)
			if err != nil {
				log.Fatal(err)
			}
		}
		if weight == "" {
			message := model.Message{
				Message: "detail of the model does not exist",
			}
			c.JSON(404, message)
		} else {
			detail := model.DetailOutput{
				Weight:   weight,
				Height:   height,
				Built:    built,
				Designer: designer,
			}
			c.JSON(200, detail)
		}
	}

	defer rows.Close()
	defer db.Close()
}

func (ct *Controller) GetDetailFromString(c *gin.Context) {

	// body := c.Request.Body
	// value, err := ioutil.ReadAll(body)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// var input model.DetailInput
	// er := json.Unmarshal([]byte(value), &input)
	// if er != nil {
	// 	fmt.Println(err.Error())
	// }
	// e := reflect.ValueOf(&input).Elem()
	// modelID := fmt.Sprint(e.Field(0).Interface())

	modelID := c.Param("modelID")
	//get info from string
	var jsonString string
	if string(modelID) == "1" {
		jsonString = "{\"weight\":\"15kg\",\"height\":\"1m\",\"built\":\"iron\",\"desinger\":\"Robert_Cai\"}"
	} else if string(modelID) == "2" {
		jsonString = "{\"weight\":\"50kg\",\"height\":\"5m\",\"built\":\"Copper, iron, Alloy\",\"desinger\":\"Kun_Su\"}"
	} else if string(modelID) == "3" {
		jsonString = "{\"weight\":\"23kg\",\"height\":\"3m\",\"built\":\"Alloy, Plastic, Glass\",\"desinger\":\"Taylor_Yang\"}"
	} else if string(modelID) == "4" {
		jsonString = "{\"weight\":\"13kg\",\"height\":\"1m\",\"built\":\"Alloy, Plastic, Glass\",\"desinger\":\"Robert_Cai\"}"
	}
	// else if string(modelID) == "5" {
	// 	jsonString = "{\"weight\":\"43kg\",\"height\":\"4m\",\"built\":\"Plastic, Glass\",\"desinger\":\"Taylor_Yang\"}"
	// } else if string(modelID) == "6" {
	// 	jsonString = "{\"weight\":\"61kg\",\"height\":\"5m\",\"built\":\"Iron, Plastic, Glass\",\"desinger\":\"Taylor_Yang\"}"
	// }
	c.String(200, jsonString)

	// connect to the db

	// db, err := connectDB(c)
	// if db == nil {
	// 	message := model.Message{
	// 		Message: "DB has something wrong",
	// 	}
	// 	c.JSON(500, message)
	// 	panic(err.Error())
	// }

	// var (
	// 	weight   string
	// 	height   string
	// 	built    string
	// 	designer string
	// )
	// rows, err := db.Query("SELECT weight, height, built, designer from detail where id = ?", modelID)
	// if err != nil {
	// 	if strings.Contains(err.Error(), "Access denied") {
	// 		message := model.Message{
	// 			Message: "DB access error, username or password is wrong",
	// 		}
	// 		c.JSON(500, message)
	// 		panic(err.Error())
	// 	}
	// 	message := model.Message{
	// 		Message: "Query statements has something wrong",
	// 	}
	// 	c.JSON(500, message)
	// 	panic(err.Error())
	// } else {
	// 	for rows.Next() {
	// 		err := rows.Scan(&weight, &height, &built, &designer)
	// 		if err != nil {
	// 			log.Fatal(err)
	// 		}
	// 	}
	// 	if weight == "" {
	// 		message := model.Message{
	// 			Message: "detail of the model does not exist",
	// 		}
	// 		c.JSON(404, message)
	// 	} else {
	// 		detail := model.DetailOutput{
	// 			Weight:   weight,
	// 			Height:   height,
	// 			Built:    built,
	// 			Designer: designer,
	// 		}
	// 		c.JSON(200, detail)
	// 	}
	// }

	// defer rows.Close()
	// defer db.Close()
}

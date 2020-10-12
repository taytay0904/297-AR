package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Hello World")

	r := gin.Default()
	r.GET("/", getModel)
	// r.POST("api/signup", querySignUp)
	// r.GET("api/login", queryLogin)
	// r.GET("api/reset", queryResetPassword)
	// r.POST("api/companySignup", queryCompanySignUp)
	// r.POST("api/addressCompanyAssociate", queryCompanyAddressAssociate)
	// r.POST("api/addItem", queryAddItem)
	// r.POST("api/addOrder", queryAddOrder)
	// r.POST("api/companyList", queryGetCompanyList)
	// r.POST("api/updateUserStatus", queryUpdateUserStatus)
	r.Run()
}

func getModel(c *gin.Context) {
	//getting data from request
	// email := c.Query("email")
	// password := c.Query("password")
	// email = strings.ToLower(string(email))

	//TODO: MD5 for encode and decode password

	// connect to the db
	db, err := connectDB(c)
	if err != nil {
		fmt.Println("DB error")
		c.JSON(500, gin.H{
			"message": "DB connection problem",
		})
		panic(err.Error())
	}

	defer db.Close()
	//select userID from db
	var (
		id              int
		name            string
		ICON_Path       string
		model_File_Path string
		detail_Info_ID  string
		sound_Path      string
	)
	rows, err := db.Query("SELECT * from user")
	if err != nil {
		fmt.Println("Login error")
		if strings.Contains(err.Error(), "Access denied") {
			c.JSON(500, gin.H{
				"message": "DB access error, username or password is wrong",
			})
			panic(err.Error())
		}
		c.JSON(404, gin.H{
			"message": "Such username or password is incorrect",
		})
		panic(err.Error())
	} else {
		for rows.Next() {
			err := rows.Scan(&userID, &status)
			if err != nil {
				log.Fatal(err)
			}
		}
		if userID == 0 {
			c.JSON(404, gin.H{
				"message": "Such username or password is incorrect",
			})
		} else if status == "pending" {
			c.JSON(404, gin.H{
				"message": "Such user has not been approved by admin",
			})
		} else {
			c.JSON(200, gin.H{
				"userID":  userID,
				"message": "User Found",
			})
		}

	}
	defer rows.Close()
}

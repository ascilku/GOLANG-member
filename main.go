package main

import (
	"fmt"
	"go-test-2/handler"
	"go-test-2/member"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/member?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err.Error())
	} else {
		// db.Debug().AutoMigrate(&member.Member{})
		new_repository := member.NewRepository(db)
		new_service := member.NewService(new_repository)
		new_handler := handler.NewHandler(new_service)

		router := gin.Default()
		api := router.Group("v1")
		api.POST("handler", new_handler.SaveHandler)
		api.POST("login", new_handler.LoginHandler)
		api.POST("check-email", new_handler.CheckEmailAvailable)
		router.Run()
	}

}

// func handler(g *gin.Context) {
// 	dsn := "root:@tcp(127.0.0.1:3306)/member?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {

// 		g.JSON(http.StatusBadRequest, err.Error())
// 	} else {
// 		var key_member []member.Member
// 		db.Find(&key_member)
// 		g.JSON(http.StatusOK, key_member)
// 	}
// }

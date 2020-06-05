package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/ifconfigure/gorm-paginate/paginage"
	"github.com/jinzhu/gorm"
	"strconv"
)

type UserLikes struct {
	FromUserID int `json:"from_user_id"`
	ToUserID   int `json:"to_user_id"`
}

func main() {

	r := gin.Default()

	r.GET("user_likes", func(c *gin.Context) {

		toUserID := c.MustGet("UserID").(uint32)

		//current page
		currentPage, _ := strconv.ParseInt(c.Query("page_index"), 10, 64)
		if currentPage == 0 {
			currentPage = 1
		}

		var userLikes []UserLikes

		db, _ := gorm.Open("mysql", "xxxx")

		//with oder and preload
		userLikesTx := db.Model(UserLikes{}).
			Order("created_at desc").
			Where("to_user_id = ?", toUserID).
			Preload("User.Country")

		res, err := paginage.Paginate(userLikesTx, int(currentPage), &userLikes)

		if err != nil {
			fmt.Println(err)
			return
		}

		c.JSON(200, res)
	})

}

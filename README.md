# gorm-paginate
Gorm Pagination

### Usage
`go get github.com/ifconfigure/gorm-paginate`


### Example

```
userLikesTx := db.Model(UserLikes{}).
    Order("created_at desc").
    Where("to_user_id = ?", toUserID).
    Preload("User.Country")

res, err := paginage.Paginate(userLikesTx, int(currentPage), &userLikes)

if err != nil {
    fmt.Println(err)
    return
}

c.JSON(200 ,res)

```
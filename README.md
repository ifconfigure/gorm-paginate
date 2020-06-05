# gorm-paginate
Gorm Pagination

### Usage && 使用方式 
`go get github.com/ifconfigure/gorm-paginate`


### Example && 例子

```
// 1、Chaining - 链式操作查询
userLikesTx := db.Model(UserLikes{}).
    Order("created_at desc").
    Where("to_user_id = ?", toUserID).
    Preload("User.Country")

//2、use paginate - 调用分页类
res, err := paginage.Paginate(userLikesTx, int(currentPage), &userLikes)

if err != nil {
    fmt.Println(err)
    return
}

3、output - 输出
c.JSON(200 ,res)

```
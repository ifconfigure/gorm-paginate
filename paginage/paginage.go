package paginage

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

const PageSize = 15

func Paginate(queryTx *gorm.DB, currentPage int, model interface{}) (map[string]interface{}, error) {
	//get total count
	var total int
	err := queryTx.Count(&total).Error
	if err != nil {
		fmt.Println(err)
	}

	//get page count
	var pageMax int
	if total%PageSize == 0 {
		pageMax = total / PageSize
	} else {
		pageMax = total/PageSize + 1
	}

	//get offset
	offset := PageSize * (currentPage - 1)
	err = queryTx.Limit(PageSize).Offset(offset).Find(model).Error

	if err != nil {
		fmt.Println(err)
	}

	return map[string]interface{}{
		"list":       model,
		"pageIndex":  currentPage,
		"pageSize":   PageSize,
		"pages":      pageMax,
		"totalItems": total,
	}, err
}

package dao

import "gorm.io/gorm"

// QueryOption dao层查询结构体
type QueryOption struct {
	Offset   int
	PageSize int
}

// Paginate 分页排序操作
func Paginate(co *QueryOption) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := co.Offset
		if page == 0 {
			page = 1
		}

		pageSize := co.PageSize
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		db.Offset(offset).Limit(pageSize)

		return db
	}
}

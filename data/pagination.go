package data

import "math"

type Paging struct {
	CurrentPage int `json:"currentPage"` // 当前页码
	PageSize    int `json:"pageSize"`    // 每页显示多少条
	TotalCount  int `json:"totalCount"`  // 总条数
	PageCount   int `json:"pageCount"`   // 总共多少页
	Offset      int `json:"-"`           // 起始页
}

func NewPagination(page, pageSize, total int) *Paging {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	pageCount := int(math.Ceil(float64(total) / float64(pageSize)))

	if page > pageCount {
		page = pageCount
	}
	offset := pageSize * (page - 1)
	paging := new(Paging)
	paging.CurrentPage = page
	paging.PageSize = pageSize
	paging.TotalCount = total
	paging.PageCount = pageCount
	paging.Offset = offset
	return paging
}

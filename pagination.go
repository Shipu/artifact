package artifact

import (
	"context"
	"gorm.io/gorm"
)

type pagingQuery struct {
	DB          *gorm.DB
	Ctx         context.Context
	FilterQuery interface{}
	LimitCount  int64
	PageCount   int64
}

type PagingQuery interface {
	Limit(limit int64) PagingQuery
	Page(page int64) PagingQuery
	Context(ctx context.Context) PagingQuery
}

type PaginationData struct {
	Total     int64 `json:"total"`
	Page      int64 `json:"page"`
	PerPage   int64 `json:"perPage"`
	Prev      int64 `json:"prev"`
	Next      int64 `json:"next"`
	TotalPage int64 `json:"totalPage"`
}

//type Paginator struct {
//	TotalRecord int64 `json:"total_record"`
//	TotalPage   int64 `json:"total_page"`
//	Limit       int64 `json:"limit"`
//	Page        int64 `json:"page"`
//	PrevPage    int64 `json:"prev_page"`
//	NextPage    int64 `json:"next_page"`
//}
//
//func (p *Paginator) PaginationData() *PaginationData {
//	data := PaginationData{
//		Total:     p.TotalRecord,
//		Page:      p.Page,
//		PerPage:   p.Limit,
//		Prev:      0,
//		Next:      0,
//		TotalPage: p.TotalPage,
//	}
//	if p.Page != p.PrevPage && p.TotalRecord > 0 {
//		data.Prev = p.PrevPage
//	}
//	if p.Page != p.NextPage && p.TotalRecord > 0 && p.Page <= p.TotalPage {
//		data.Next = p.NextPage
//	}
//
//	return &data
//}

func (paging *pagingQuery) Page(page int64) PagingQuery {
	if page < 1 {
		paging.PageCount = 1
	} else {
		paging.PageCount = page
	}
	return paging
}

func (paging *pagingQuery) Limit(limit int64) PagingQuery {
	if limit < 1 {
		paging.LimitCount = 10
	} else {
		paging.LimitCount = limit
	}
	return paging
}

func (paging *pagingQuery) Context(ctx context.Context) PagingQuery {
	paging.Ctx = ctx
	return paging
}

func NewPaginate(db *gorm.DB) *pagingQuery {
	return &pagingQuery{
		DB: db,
	}
}

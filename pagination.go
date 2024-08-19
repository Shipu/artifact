package artifact

import (
	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
	"gorm.io/gorm"
	"math"
	"strconv"
)

var Pagination paginator.Paginator

type PaginationMeta struct {
	CurrentPage int `json:"current_page"`
	PerPage     int `json:"per_page"`
	LastPage    int `json:"last_page"`
	Total       int `json:"total"`
}

type Paginator struct {
	Meta PaginationMeta `json:"meta"`
	*paginator.Paginator
}

func (paginator *Paginator) PaginateScope() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (paginator.Meta.CurrentPage - 1) * paginator.Meta.PerPage
		return db.Offset(offset).Limit(paginator.Meta.PerPage)
	}
}

func (paginator *Paginator) CursorPaginate(query *gorm.DB, v interface{}) (*gorm.DB, paginator.Cursor, error) {
	return paginator.Paginate(query, v)
}

func (paginator *Paginator) updateMeta(v interface{}, request map[string]interface{}) {
	limit, _ := strconv.ParseInt(request["limit"].(string), 10, 64)
	page, _ := strconv.ParseInt(request["page"].(string), 10, 64)

	if request["after"] != nil && request["after"] != "" {
		paginator.SetAfterCursor(request["after"].(string))
	}

	if request["before"] != nil && request["before"] != "" {
		paginator.SetBeforeCursor(request["before"].(string))
	}

	if limit == 0 {
		limit = 10
	}
	if page == 0 {
		page = 1
	}

	paginator.SetLimit(int(limit))
	paginator.Meta.CurrentPage = int(page)
	paginator.Meta.PerPage = int(limit)

	var totalRows int64
	DB.Model(v).Count(&totalRows)
	paginator.Meta.Total = int(totalRows)
	paginator.Meta.LastPage = int(math.Ceil(float64(totalRows) / float64(limit)))
}

func NewPaginator(v interface{}, request map[string]interface{}) *Paginator {
	opts := []paginator.Option{
		&paginator.Config{
			Order: paginator.ASC,
		},
	}

	p := paginator.New(opts...)

	newInstance := &Paginator{Meta: PaginationMeta{}, Paginator: p}
	newInstance.updateMeta(v, request)

	return newInstance
}

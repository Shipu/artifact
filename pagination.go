package artifact

import (
	"github.com/pilagod/gorm-cursor-paginator/v2/paginator"
	"gorm.io/gorm"
	"math"
	"strconv"
)

var Pagination paginator.Paginator

type PaginationMeta struct {
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	LastPage    int   `json:"last_page"`
	Total       int64 `json:"total"`
}

type Paginator struct {
	Meta PaginationMeta `json:"meta"`
	db   *gorm.DB
	*paginator.Paginator
	model  interface{}
	filter map[string]interface{}
}

func (paginator *Paginator) PaginateScope(page int, limit int) func(db *gorm.DB) *gorm.DB {
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}
	var totalRows int64

	paginator.Meta.CurrentPage = page
	paginator.Meta.PerPage = limit

	DB.Model(paginator.model).Where(paginator.filter).Count(&totalRows)
	paginator.Meta.Total = totalRows
	paginator.Meta.LastPage = int(math.Ceil(float64(totalRows) / float64(limit)))

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
	paginator.Meta.Total = totalRows
	paginator.Meta.LastPage = int(math.Ceil(float64(totalRows) / float64(limit)))
}

func NewPaginator(v interface{}, queryFilter map[string]interface{}) *Paginator {
	opts := []paginator.Option{
		&paginator.Config{
			Order: paginator.ASC,
		},
	}

	p := paginator.New(opts...)

	newInstance := &Paginator{Meta: PaginationMeta{}, Paginator: p, model: v, filter: queryFilter}
	return newInstance
}

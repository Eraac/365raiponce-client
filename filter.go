package raiponce

import "fmt"

const (
	OrderASC  = true
	OrderDESC = false
)

type filters map[string][]string
type orders map[string]bool
type groups []string

type QueryFilter struct {
	filters    filters
	orders     orders
	groups     groups
	MaxPerPage int
	Page       int
}

func NewFilter() *QueryFilter {
	return &QueryFilter{
		filters: map[string][]string{},
		orders:  map[string]bool{},
		groups:  []string{},
	}
}

func (f *QueryFilter) buildQuery() string {
	var query string

	query += f.filters.buildQuery()
	query += f.orders.buildQuery()
	query += f.groups.buildQuery()

	if f.Page > 0 {
		query += fmt.Sprintf("_page=%d&", f.Page)
	}

	if f.MaxPerPage > 0 {
		query += fmt.Sprintf("_max_per_page=%d&", f.MaxPerPage)
	}

	return query
}

func (f filters) buildQuery() string {
	var query string

	for key, value := range f {
		var array string

		if len(value) > 1 {
			array = "[]"
		}

		for _, v := range value {
			query += fmt.Sprintf("filter[%s]%s=%s&", key, array, v)
		}
	}

	return query
}

func (o orders) buildQuery() string {
	var query string

	for key, value := range o {
		order := "ASC"

		if value == OrderDESC {
			order = "DESC"
		}

		query += fmt.Sprintf("filter[_order][%s]=%s&", key, order)
	}

	return query
}

func (g groups) buildQuery() string {
	var query string

	for _, value := range g {
		query += fmt.Sprintf("filter[_group][]=%s&", value)
	}

	return query
}

func (f *QueryFilter) AddFilter(key string, value string) {
	f.filters[key] = append(f.filters[key], value)
}

func (f *QueryFilter) AddOrder(orderBy string, order bool) {
	f.orders[orderBy] = order
}

func (f *QueryFilter) AddGroup(groupBy string) {
	f.groups = append(f.groups, groupBy)
}

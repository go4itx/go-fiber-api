package pagination

import "home/pkg/constant"

// Pagination unified paging parameters
type Pagination struct {
	Page int `validate:"numeric"`
	Size int `validate:"numeric"`
}

// paging Processing paging parameters
func (p *Pagination) Paging() (offset int, pageSize int) {
	if p.Page == 0 {
		p.Page = 1
	}

	if p.Size == 0 {
		p.Size = constant.DefaultPageSize
	}

	return (p.Page - 1) * p.Size, p.Size
}

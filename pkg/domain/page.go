// Package domain provides the domain model for the application.
package domain

// Page represents a paginated response.
type Page[Entity any] struct {
	Items      []Entity
	Total      int
	PageSize   int
	PageNumber int
	TotalPages int
}

// EmptyPage returns an empty page with default values.
func EmptyPage[Entity any]() Page[Entity] {
	return Page[Entity]{
		Items:      []Entity{},
		Total:      0,
		PageSize:   0,
		PageNumber: 0,
		TotalPages: 0,
	}
}

// IPageBuilder is an interface for building a paginated response.
type IPageBuilder[Entity any] interface {
	SetItems(items []Entity) IPageBuilder[Entity]
	SetTotal(total int) IPageBuilder[Entity]
	SetPageSize(pageSize int) IPageBuilder[Entity]
	SetPageNumber(pageNumber int) IPageBuilder[Entity]
	SetTotalPages(totalPages int) IPageBuilder[Entity]
	Build() Page[Entity]
}

// PageBuilder is a struct that implements the IPageBuilder interface.
type PageBuilder[Entity any] struct {
	items      []Entity
	total      int
	pageSize   int
	pageNumber int
	totalPages int
}

// SetItems sets the items for the paginated response.
func (builder *PageBuilder[Entity]) SetItems(items []Entity) IPageBuilder[Entity] {
	builder.items = items
	return builder
}

// SetTotal sets the total number of items for the paginated response.
func (builder *PageBuilder[Entity]) SetTotal(total int) IPageBuilder[Entity] {
	builder.total = total
	return builder
}

// SetPageSize sets the page size for the paginated response.
func (builder *PageBuilder[Entity]) SetPageSize(pageSize int) IPageBuilder[Entity] {
	builder.pageSize = pageSize
	return builder
}

// SetPageNumber sets the page number for the paginated response.
func (builder *PageBuilder[Entity]) SetPageNumber(pageNumber int) IPageBuilder[Entity] {
	builder.pageNumber = pageNumber
	return builder
}

// SetTotalPages sets the total number of pages for the paginated response.
func (builder *PageBuilder[Entity]) SetTotalPages(totalPages int) IPageBuilder[Entity] {
	builder.totalPages = totalPages
	return builder
}

// Build builds the paginated response.
func (builder *PageBuilder[Entity]) Build() Page[Entity] {
	totalPages := builder.totalPages
	if totalPages == 0 {
		totalPages = builder.total / builder.pageSize
		if builder.total%builder.pageSize > 0 {
			totalPages++
		}
	}
	return Page[Entity]{
		Items:      builder.items,
		Total:      builder.total,
		PageSize:   builder.pageSize,
		PageNumber: builder.pageNumber,
		TotalPages: totalPages,
	}
}

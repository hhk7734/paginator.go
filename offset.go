package paginator

type Offset struct {
	// Page is the current page. Page numbering starts at 0.
	Page int64
	// Size is the number of items per page.
	Size int64
}

type OffsetModel struct {
	// Page is the current page. Page numbering starts at 0.
	Page int64 `json:"page"`
	// Size is the number of items per page.
	Size int64 `json:"size"`
	// TotalItems is the total number of items.
	TotalItems int64 `json:"total_items"`
	// TotalPages is the total number of pages.
	TotalPages int64 `json:"total_pages"`
	// First is true if the current page is the first page.
	First bool `json:"first"`
	// Last is true if the current page is the last page.
	Last bool `json:"last"`
}

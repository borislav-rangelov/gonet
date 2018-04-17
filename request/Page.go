package request

import (
	"strconv"
)

// DefaultPageSize The default page size
const DefaultPageSize = 10

// Page A page request
type Page struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

// GetPage Get the currently requested page
func (p Page) GetPage() int {
	return p.Page
}

//GetSize Get the requested content size
func (p Page) GetSize() int {
	return p.Size
}

// GetPage constructs a page request
func GetPage(page int, size int) Page {
	if page < 0 {
		page = 0
	}
	if size <= 0 {
		size = DefaultPageSize
	}

	return Page{Page: page, Size: size}
}

// MapPage Maps values to a page object
func MapPage(values map[string]string) Page {

	page := parseInt(values["page"])
	size := parseInt(values["size"])

	return GetPage(page, size)
}

func parseInt(val string) int {
	intVal, _ := strconv.Atoi(val)
	return intVal
}

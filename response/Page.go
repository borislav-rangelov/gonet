package response

import (
	"github.com/borislav-rangelov/gonet"
	"github.com/borislav-rangelov/gonet/request"
)

const defaultPageSize = request.DefaultPageSize

// Page A page response
type Page struct {
	Content *[]interface{} `json:"content"`
	Page    *PageInfo      `json:"page"`
}

func (p *Page) GetContent() *[]interface{} {
	return p.Content
}

func (p *Page) GetPage() *PageInfo {
	return p.Page
}

// PageInfo Page information
type PageInfo struct {
	TotalPages int   `json:"total_pages"`
	Page       int   `json:"page"`
	Size       int   `json:"size"`
	Count      int64 `json:"count"`
}

func (pi *PageInfo) GetTotalPages() int {
	return pi.TotalPages
}

func (pi *PageInfo) GetPage() int {
	return pi.Page
}

func (pi *PageInfo) GetSize() int {
	return pi.Size
}

func (pi *PageInfo) GetCount() int64 {
	return pi.Count
}

func BuildPage(content *[]interface{}, reqPage gonet.PageRequest, count int64) Page {
	size := reqPage.GetSize()
	pages := 0

	if size > 0 {
		pages = int(count / int64(size))
		if count%int64(size) > 0 {
			pages++
		}
	}

	current := reqPage.GetPage()
	info := PageInfo{TotalPages: pages, Page: current, Size: size, Count: count}
	return Page{Content: content, Page: &info}
}

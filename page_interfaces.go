package gonet

// PageRequest A page request
type PageRequest interface {
	GetPage() int
	GetSize() int
}

// PageResult A page result
type PageResult interface {
	GetContent() *[]interface{}
	GetPage() *PageInfo
}

// PageInfo Page information
type PageInfo interface {
	PageRequest
	GetTotalPages() int
	GetCount() int64
}

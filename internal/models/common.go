package models

type Pagination struct {
	Page     int `json:"page"`
	PageSize int `json:"pageSize"`
}

type SortBy struct {
	Column    string `json:"column"`
	Direction string `json:"direction"`
}

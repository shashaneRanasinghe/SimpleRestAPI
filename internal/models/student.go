package models

type StudentResponse struct {
	Status  string  `json:"status"`
	Data    Student `json:"data"`
	Message string  `json:"message"`
}

type StudentSearchRequest struct {
	SearchString string     `json:"searchString"`
	SortBy       SortBy     `json:"sortBy"`
	Pagination   Pagination `json:"pagination"`
}

type StudentSearchData struct {
	TotalElements int       `json:"totalElements"`
	Data          []Student `json:"data"`
}

type StudentSearchResponse struct {
	Status  string            `json:"status"`
	Data    StudentSearchData `json:"data"`
	Message string            `json:"message"`
}

type StudentListResponse struct {
	Status  string    `json:"status"`
	Data    []Student `json:"data"`
	Message string    `json:"message"`
}

type Student struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Year      int    `json:"year"`
}

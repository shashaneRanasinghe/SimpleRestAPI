package models

type LecturerResponse struct {
	Status  string   `json:"status"`
	Data    Lecturer `json:"data"`
	Message string   `json:"message"`
}

type LecturerSearchRequest struct {
	SearchString string     `json:"searchString"`
	SortBy       SortBy     `json:"sortBy"`
	Pagination   Pagination `json:"pagination"`
}

type LecturerSearchData struct {
	TotalElements int        `json:"totalElements"`
	Data          []Lecturer `json:"data"`
}

type LecturerSearchResponse struct {
	Status  string             `json:"status"`
	Data    LecturerSearchData `json:"data"`
	Message string             `json:"message"`
}

type LecturerListResponse struct {
	Status  string     `json:"status"`
	Data    []Lecturer `json:"data"`
	Message string     `json:"message"`
}

type Lecturer struct {
	ID        int    `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Year      int    `json:"year"`
}

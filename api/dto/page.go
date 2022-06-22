package dto

type Page struct {
    Page        int `json:"page"`
    PerPage     int `json:"per_page"`
    TotalCounts int `json:"total_count"`
    TotalPages  int `json:"total_pages"`
}

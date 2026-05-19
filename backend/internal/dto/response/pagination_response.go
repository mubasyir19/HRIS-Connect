package response

type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

type Pagination struct {
	Page       int   `json:"page"`
	Limit      int   `json:"limit"`
	TotalRows  int64 `json:"totalRows"`
	TotalPages int   `json:"totalPages"`
}

func NewPaginatedResponse(data interface{}, page, limit int, totalRows int64) *PaginatedResponse {
	totalPages := int((totalRows + int64(limit) - 1) / int64(limit))
	if totalPages < 1 {
		totalPages = 1
	}

	return &PaginatedResponse{
		Data: data,
		Pagination: Pagination{
			Page:       page,
			Limit:      limit,
			TotalRows:  totalRows,
			TotalPages: totalPages,
		},
	}
}

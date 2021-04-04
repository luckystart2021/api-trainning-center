package class

type Class struct {
	Id        int64  `json:"id"`
	Code      string `json:"code"`
	IdCourse  int64  `json:"id_course"`
	Quantity  int64  `json:"quantity"`
	IsDelete  bool   `json:"is_delete"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
	UpdatedAt string `json:"updated_at"`
	UpdatedBy string `json:"updated_by"`
}

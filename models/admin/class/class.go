package class

type Class struct {
	Id        int64  `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	IdCourse  int64  `json:"id_course"`
	Quantity  int64  `json:"quantity"`
	IdTeacher int64  `json:"id_teacher"`
	IsDelete  bool   `json:"is_delete"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
	UpdatedAt string `json:"updated_at"`
	UpdatedBy string `json:"updated_by"`
}

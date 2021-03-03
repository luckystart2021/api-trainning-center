package student

type Student struct {
	Id          int64  `json:"id"`
	Code        string `json:"code"`
	Sex         string `json:"sex"`
	DateOfBirth string `json:"date_of_birth"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	FullName    string `json:"full_name"`
	IdClass     int64  `json:"id_class"`
	CreatedAt   string `json:"created_at"`
	CreatedBy   string `json:"created_by"`
	UpdatedAt   string `json:"updated_at"`
	UpdatedBy   string `json:"updated_by"`
}

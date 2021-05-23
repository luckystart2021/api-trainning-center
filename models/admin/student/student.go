package student

type Student struct {
	Id             int64  `json:"id"`
	Code           string `json:"code"`
	Sex            string `json:"sex"`
	DateOfBirth    string `json:"date_of_birth"`
	Phone          string `json:"phone"`
	Address        string `json:"address"`
	FullName       string `json:"full_name"`
	IdClass        int64  `json:"id_class"`
	CreatedAt      string `json:"created_at"`
	CreatedBy      string `json:"created_by"`
	UpdatedAt      string `json:"updated_at"`
	UpdatedBy      string `json:"updated_by"`
	CMND           string `json:"cmnd"`
	CNSK           bool   `json:"cnsk"`
	GPLX           string `json:"gplx"`
	Exp            int    `json:"exp"`
	NumberOfKm     int    `json:"number_of_km"`
	AmountComplete string `json:"amount"`
	AmountRemain   string `json:"amount_remain"`
	DiemLyThuyet   string `json:"diem_ly_thuyet"`
	DiemThucHanh   string `json:"diem_thuc_hanh"`
	KetQua         string `json:"ket_qua"`
	Email          string `json:"email"`
}

package student

import (
	"api-trainning-center/models/admin/student"
	"api-trainning-center/service/response"
	"database/sql"
)

type IStudentService interface {
	ShowStudents() ([]student.Student, error)
	ShowStudent(idStudent int) (student.Student, error)
	SearchStudentInformation(codeStudent string) (student.Student, error)
	CreateStudent(sex, dayOfBirth, phone, address, fullName, userName string, idClass int,
		cmnd string, cnsk bool, gplx string, exp, numberKm int, amount float64) (response.MessageResponse, error)
	UpdateStudent(idStudent int, sex, dayOfBirth, phone, address, fullName, userName string,
		idClass int, cmnd string, cnsk bool, gplx string, exp, numberKm int, amount float64,
		diemLyThuyet, diemThucHanh string, ketQua bool,
	) (response.MessageResponse, error)
	DeleteStudent(idStudent int) (response.MessageResponse, error)
}

type StoreStudent struct {
	db *sql.DB
}

func NewStoreStudent(db *sql.DB) *StoreStudent {
	return &StoreStudent{
		db: db,
	}
}

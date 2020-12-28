package models

type AccountRequest struct {
	Email    string `json:"email"`
	UserName string `json:"username"`
	PassWord string `json:"password"`
	Role     string `json:"role"`
}

type AccountReponse struct {
	Status bool `json:"status"`
}

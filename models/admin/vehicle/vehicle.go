package vehicle

type Vehicle struct {
	Id        int    `json:"id"`
	BienSoXe  string `json:"bien_so_xe"`
	LoaiXe    string `json:"loai_xe"`
	Status    bool   `json:"status"`
	IsDeleted string `json:"is_deleted"`
}

type FindOneVehicle struct {
	Id        int    `json:"id"`
	BienSoXe  string `json:"bien_so_xe"`
	LoaiXe    string `json:"loai_xe"`
	Status    bool   `json:"status"`
	IsDeleted string `json:"is_deleted"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
	UpdatedAt string `json:"updated_at"`
	UpdatedBy string `json:"updated_by"`
}

type VehicleRequest struct {
	BienSoXe string `json:"bien_so_xe"`
	LoaiXe   string `json:"loai_xe"`
}

type VehicleUpdateRequest struct {
	BienSoXe  string `json:"bien_so_xe"`
	LoaiXe    string `json:"loai_xe"`
	IsDeleted string `json:"is_deleted"`
}

package vehicle

// type Vehicle struct {
// 	Id         int    `json:"id"`
// 	BienSoXe   string `json:"bien_so_xe"`
// 	LoaiXe     string `json:"loai_xe"`
// 	Status     bool   `json:"status"`
// 	IsContract bool   `json:"is_contract"`
// 	IsDeleted  bool   `json:"is_deleted"`
// }

type FindOneVehicle struct {
	Id         int    `json:"id"`
	BienSoXe   string `json:"bien_so_xe"`
	LoaiXe     string `json:"loai_xe"`
	Status     bool   `json:"status"`
	IsContract bool   `json:"is_contract"`
	IsDeleted  string `json:"is_deleted"`
	CreatedAt  string `json:"created_at"`
	CreatedBy  string `json:"created_by"`
	UpdatedAt  string `json:"updated_at"`
	UpdatedBy  string `json:"updated_by"`
}

type VehicleRequest struct {
	BienSoXe   string `json:"bien_so_xe"`
	LoaiXe     string `json:"loai_xe"`
	IsContract bool   `json:"xe_hop_dong"`
}

type VehicleUpdateRequest struct {
	BienSoXe   string `json:"bien_so_xe"`
	LoaiXe     string `json:"loai_xe"`
	IsDeleted  bool   `json:"is_deleted"`
	IsContract bool   `json:"xe_hop_dong"`
}

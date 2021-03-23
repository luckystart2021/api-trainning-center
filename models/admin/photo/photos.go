package photo

type Photo struct {
	Id        int    `json:"id"`
	IdAlbum   int    `json:"id_album"`
	Img       string `json:"img"`
	Title     string `json:"title"`
	Meta      string `json:"meta"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"created_by"`
	UpdatedAt string `json:"updated_at"`
	UpdatedBy string `json:"updated_by"`
}

type Album struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Meta string `json:"meta"`
}

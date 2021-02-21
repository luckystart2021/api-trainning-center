package slide

type Slide struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Img       string `json:"img"`
	Hide      bool   `json:"hide"`
	CreatedAt string `json:"created_at"`
	CreatedBy string `json:"create_by"`
}

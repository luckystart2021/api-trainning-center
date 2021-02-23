package result

type Result struct {
	IdSuite int      `json:"id_suite"`
	Answers []Answer `json:"answers"`
}
type Answer struct {
	IdQuestion int    `json:"id_question"`
	IdAnswer   string `json:"answer"`
}

package course

type ErrorResponse struct {
	StatusCode string `json:"statusCode"`
	Message    string `json:"message"`
}

var (
	ErrIdCourseParseFail = &ErrorResponse{StatusCode: "courseRequest", Message: "Giá trị truyền vào không hợp lệ"}
)

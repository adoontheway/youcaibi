package defs

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	StatusCode int
	Error      Err
}

var (
	ErrorRequestBodyParseFaild = ErrorResponse{
		StatusCode: 400,
		Error: Err{
			Error:     "Parse Body failed",
			ErrorCode: "001",
		},
	}

	ErrorNotAuthed = ErrorResponse{
		StatusCode: 401,
		Error: Err{
			Error:     "User Authentication failed",
			ErrorCode: "002",
		},
	}
)

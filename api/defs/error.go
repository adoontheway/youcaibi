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
	ErrorRequestBodyParseFailed = ErrorResponse{
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

	ErrorDBError            = ErrorResponse{StatusCode: 500, Error: Err{Error: "DB ops failed", ErrorCode: "003"}}
	ErrorInternalErrorFault = ErrorResponse{StatusCode: 500, Error: Err{Error: "Internal service error", ErrorCode: "004"}}
)

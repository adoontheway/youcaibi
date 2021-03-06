package main

type HomePage struct {
	Name string
}

type UserPage struct {
	Name string
}

// api透传规范
type ApiBody struct {
	Url         string `json:"url"`
	Method      string `json:"method"`
	RequestBody string `json:"req_body"`
}

type Err struct {
	Error     string `json:"error"`
	ErrorCode string `json:"error_code"`
}

var (
	ErrorRequestNotRecognized = Err{
		Error:     "api not recognized, bad request",
		ErrorCode: "001",
	}
	ErrorRequestBodyParseFailed = Err{
		Error:     "illegal request body",
		ErrorCode: "002",
	}
	ErrorInternalFaults = Err{
		Error:     "internal service error",
		ErrorCode: "003",
	}
)

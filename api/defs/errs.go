package defs

type Error struct {
	Message string `json:"msg"`
	ErrorCode string `json:"error_code"`
}

type ErrorResponse struct {
	HttpSC int
	Error Error
}

var (
	ErrorRequestBodyParseFailed = ErrorResponse{
		HttpSC:400,Error:Error{Message: "Request body is not correct",ErrorCode: "001"},
	}
	ErrorNotAuthUser = ErrorResponse{
		HttpSC:401, Error: Error{Message:"User authentication failed.", ErrorCode:"002"},
	}
	ErrorDBError = ErrorResponse{HttpSC: 500, Error: Error{Message: "DB ops failed", ErrorCode: "003"}}
	ErrorInternalFaults = ErrorResponse{HttpSC: 500, Error: Error{Message: "Internal service error", ErrorCode: "004"}}
)

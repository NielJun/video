package defs


type Err struct {
  Error     string `json:"error"`
  ErrorCode string `json:"error_code"`
}

type ErrResponse struct {
  HttpSc int
  Error  Err
}

var (
  // 出现咩有请求的消息体错误
  ErrorRequestBodyParseFailed = ErrResponse{HttpSc: 400, Error: Err{Error: "Request body is not correct", ErrorCode: "001"}}
  //
  ErrorNotAuthUser = ErrResponse{HttpSc: 401, Error: Err{Error: "User auth not correct", ErrorCode: "002"}}

  ErrorDBError = ErrResponse{HttpSc: 500, Error: Err{Error: "DB ops failld", ErrorCode: "003"}}
  ErrorInternalFaults   = ErrResponse{HttpSc:500,Error:Err{Error:"Internal service error",ErrorCode:"004"}}
)

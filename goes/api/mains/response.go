package mains

import (
  "encoding/json"
  "github.com/daniel/video/goes/api/defs"
  "io"
  "net/http"
)


// 处理所有的返回消息的的地方  当处理过程粗错的时候向客户端卸乳错误信息

func sendErrorResponse(w http.ResponseWriter, errResp defs.ErrResponse) {
  w.WriteHeader(errResp.HttpSc)
  resStr, _ := json.Marshal(errResp.Error)
  io.WriteString(w, string(resStr))
}

func sendNormalResponse(w http.ResponseWriter, response string, sc int) {
  w.WriteHeader(sc)
  io.WriteString(w, response)

}

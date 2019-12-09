package mains

import (
  "encoding/json"
  "github.com/daniel/video/goes/api/dbops"
  "github.com/daniel/video/goes/api/defs"
  "github.com/daniel/video/goes/api/sessions"
  "github.com/julienschmidt/httprouter"
  "io"
  "io/ioutil"
  "net/http"
)

func CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  // get params from request
  res, _ := ioutil.ReadAll(r.Body)
  ubody := &defs.UserCredential{}

  // json反序列化过程 把request中11的json发序列号花城model对象
  if err := json.Unmarshal(res, ubody); err != nil {
    sendErrorResponse(w, defs.ErrorRequestBodyParseFailed) // 反序列化失败
    return
  }

  err := dbops.AddUser(ubody.Username, ubody.Pwd) // 数据库添加用户失败
  if err != nil {
    sendErrorResponse(w, defs.ErrorDBError)
    return
  }

  // 成功的处理
  id := sessions.GenerateNewSessionId(ubody.Username)
  su := &defs.SignedUp{Success: true, SessionId: id}

  if resp, err := json.Marshal(su); err != nil {
    sendErrorResponse(w, defs.ErrorInternalFaults)
    return
  } else {
    sendNormalResponse(w, string(resp), 201)
  }
}

func Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  uname := p.ByName("user_name")
  io.WriteString(w, uname)
}

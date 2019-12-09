package mains

import (
"github.com/daniel/video/goes/api/defs"
"github.com/daniel/video/goes/api/sessions"
  "net/http"
)

// 用来验证用户是否是合法用户  包括健全和校验等方法

// 两个1自定义的Http头1 用来加入http的协议1前面
var HEADER_FIELD_SESSION = "X-Session-Id"
var HEADER_FIELD_UNAME = "X-User-Name"

// 检查用户session的合法性
func ValidateUserSession(r *http.Request) bool {

  sid := r.Header.Get(HEADER_FIELD_SESSION)
  if len(sid) == 0 {
    return false
  }
  // 对 sesson的合法性做一个检验
  uname, isexpired := sessions.IsSessionExpired(sid)   //是否过期
  if isexpired {
    return false
  }
  r.Header.Add(HEADER_FIELD_SESSION, uname)
  return true
}

func VallidateUser(w http.ResponseWriter, r *http.Request) bool {
  uname := r.Header.Get(HEADER_FIELD_UNAME)
  if len(uname) == 0 {
    sendErrorResponse(w,defs.ErrorNotAuthUser)
    return false
  }
  return true

}

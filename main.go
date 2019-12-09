package main

import (
  "github.com/daniel/video/goes/api/mains"
  "github.com/julienschmidt/httprouter"
  "net/http"
)

type MiddleWareHandler struct {
  r *httprouter.Router
}

func NewMiddleWareHandler(r *httprouter.Router) http.Handler {
  m := MiddleWareHandler{}
  m.r = r
  return m
}

// 劫持Go内部的handler 并且做一层安全校验的包装 检查session的合法性
func (m MiddleWareHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
  // check sessionß
  mains.ValidateUserSession(r)

  m.r.ServeHTTP(w, r)
}

func RegiestHandlers() *httprouter.Router {
  router := httprouter.New()
  router.POST("/user", mains.CreateUser)
  router.POST("/user/:user_name", mains.Login)
  return router
}

func main() {
  r := RegiestHandlers()
  mh := NewMiddleWareHandler(r) // 在Serve之前添加安全校验
  http.ListenAndServe(":8000", mh)
}

//Handler=>validation(1.request,2.user)  ===> business->logic=>reponse
// 1. data model
// 2. error

// 流程图： mian ----> 做一些校验、流控、处理 middleware --->defs(message,err)  --->Handlers -----> dbops ----->Respons

// Go语言的 Handler机制  若存在 handler A 和 interface A 则Go内部就把两个A认为是同一个

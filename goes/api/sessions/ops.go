package sessions

import (
  "github.com/daniel/video/goes/api/dbops"
  "github.com/daniel/video/goes/api/defs"
  "github.com/daniel/video/goes/api/utils"
  "sync"
  "time"
)

var sessionMap *sync.Map

func init() {
  sessionMap = &sync.Map{}
}

// 从数据库中加载存储的session项目
func LoadSessionsFromDB() {
  r, err := dbops.RetrieveAllSessions()
  if err != nil {
    return
  }
  r.Range(func(k, v interface{}) bool {
    ss := v.(*defs.SimepleSession)
    sessionMap.Store(k, ss)
    return true
  })

}

// 生成一个新的sessionID
func GenerateNewSessionId(un string) string {
  id, _ := utils.NewUUID()
  ct := GetNowTimeForMilli() //毫秒   session 定义的时间
  ttl := ct + 30*60*1000     // session过期时间
  session := &defs.SimepleSession{un, ttl}
  sessionMap.Store(id, session)
  dbops.InserSession(id, ttl, un)
  return id

}

// 判断当前session是否过期   不过期就返回当前session的Username
func IsSessionExpired(sid string) (string, bool) {

  session, ok := sessionMap.Load(sid)

  if ok {
    ct := GetNowTimeForMilli()
    if session.(*defs.SimepleSession).TTL < ct {
      deleteExpiredSession(sid)
      return "", true
    }
    return session.(*defs.SimepleSession).Username, false

  }
  return "", true
}

func GetNowTimeForMilli() int64 {
  return time.Now().UnixNano() / 1000000
}

func deleteExpiredSession(sid string) {
  sessionMap.Delete(sid) // 从服务器session、缓存中删除session、
  dbops.DeleteSession(sid)
}


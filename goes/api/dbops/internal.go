package dbops

import (
"github.com/daniel/video/goes/api/defs"
  "log"
  "strconv"
  "sync"
)

// 为session和数据库操作链接的一些方法和必要数据交互的方法

// 向sessions表中插入session
func InserSession(sid string, ttl int64, uname string) error {
  ttlstr := strconv.FormatInt(ttl, 10)
  stmtIns, err := dbConn.Prepare("INSERT INTO sessions (session_id,TTL,login_name) VALUES (?,?,?)")

  if err != nil {
    return err
  }
  _, err = stmtIns.Exec(sid, ttlstr, uname)
  if err != nil {
    return err
  }
  defer stmtIns.Close()
  return nil
}

// 通过sessionId查找 TTl和用户名字组成的新的session model
func RetriveSession(sid string) (*defs.SimepleSession, error) {

  session := &defs.SimepleSession{}

  stmtOut, err := dbConn.Prepare("SELECT TTL,login_name from sessions where session_id = ?")
  if err != nil {
    return nil, err
  }

  var ttl string
  var uname string
  stmtOut.QueryRow(sid).Scan(&ttl, &uname)
  //if err != nil && err != sql.ErrNoRows {
  //  return nil, err
  //}

  if res, err := strconv.ParseInt(ttl, 10, 64); err == nil {
    session.TTL = res
    session.Username = uname
  } else {
    return nil, err
  }
  defer stmtOut.Close()
  return session, nil
}

///  返回一堆以session-sessionid为结构的map数据
func RetrieveAllSessions() (*sync.Map, error) {

  m := &sync.Map{}
  stmtOut, err := dbConn.Prepare("SELECT * FROM sessions")
  if err != nil {
    log.Printf("%s", err)
    return nil, err
  }

  rows, err := stmtOut.Query()
  if err != nil {
    log.Printf("%s", err)
    return nil, err
  }

  for rows.Next() {
    var id string
    var ttlstr string
    var login_name string
    if er := rows.Scan(&id, &ttlstr, &login_name); er != nil {
      log.Printf("receive sessions error: %s", er)
      break
    }
    if ttl, eer1 := strconv.ParseInt(ttlstr, 10, 64); eer1 == nil {
      ss := &defs.SimepleSession{Username: login_name, TTL: ttl}
      m.Store(id, ss)
      log.Printf("Session id: %s , ttl : %d", id, ss.TTL)
    }
  }
  return m, nil
}

// 从数据库删除一个session
func DeleteSession(sid string)error  {
  stmtOut,err := dbConn.Prepare("DELETE FROM sessions WHERE  session_id = ?")
  if err!=nil{
    log.Panicf("%s",err)
    return nil
  }
  _,err =stmtOut.Query(sid)
  if err != nil {
    return err
  }
  return nil
}

package dbops

import (
  "database/sql"
"github.com/daniel/video/goes/api/defs"
  "github.com/daniel/video/goes/api/utils"
  "log"
  "time"
)

/////////////////////////  用户操作

// 添加用户
func AddUser(loginName string, pwd string) error {

  stmtIns, err := dbConn.Prepare("INSERT INTO users (login_name,pwd) VALUES (?,?)") // 预编译SQL语句
  if err != nil {
    return err
  }
  _, err = stmtIns.Exec(loginName, pwd)
  if err != nil {
    return err
  }
  defer stmtIns.Close()
  return nil
}

// 获取用户的登陆密码是否正确
func GetUserCredentail(loginName string) (string, error) {
  stmtOut, err := dbConn.Prepare("SELECT pwd FROM users WHERE login_name = ?")
  if err != nil {
    log.Printf("%v", err)
    return "", err
  }
  // 从数据库中取数据并且检验错误
  var pwd string
  err = stmtOut.QueryRow(loginName).Scan(&pwd)
  if err != nil && err != sql.ErrNoRows { //只是没有结果的错误
    return "", nil
  }
  defer stmtOut.Close()
  return pwd, nil
}

//删除用户
func DeleteUser(LoginName string, LoginPws string) error {
  stmtDel, err := dbConn.Prepare("DELETE FROM users WHERE login_name = ? AND pwd = ?")
  if err != nil {
    log.Printf("%v", err)
    return err
  }
  _, err = stmtDel.Exec(LoginName, LoginPws)
  if err != nil {
    return err
  }
  defer stmtDel.Close()

  return nil
}

///////////////////////////////  视屏部分

// 添加新的视屏
func AddNewVedio(aid int, name string) (*defs.Vedio_info, error) {

  // create
  vid, err := utils.NewUUID()
  if err != nil {
    return nil, err
  }
  // 以当前时间写入数据库
  t := time.Now()
  ctime := t.Format("Jan 02 2006, 15:04:05") //时间格式化语句内容固定 修改会失效 M D Y, HH:MM:SS
  stmtIns, err := dbConn.Prepare(`INSERT INTO vedio_info (id,author_id,name,display_ctime) values (?,?,?,?)`)
  if err != nil {
    return nil, err
  }
  _, err = stmtIns.Exec(vid, aid, name, ctime)
  if err != nil {
    return nil, err
  }
  res := &defs.Vedio_info{Id: vid, AuthorId: aid, Name: name, DisPlayCTime: ctime}
  defer stmtIns.Close()
  return res, nil
}

// 获取视频信息接口
func GetVedioInfo(vid string) (*defs.Vedio_info, error) {
  stmtOut, err := dbConn.Prepare("select author_id,name,display_ctime from vedio_info where id = ?")
  var aid int
  var dispaly_ctime string
  var name string
  err = stmtOut.QueryRow(vid).Scan(&aid, &name, &dispaly_ctime)

  if err != nil && err != sql.ErrNoRows {
    return nil, err
  }
  if err == sql.ErrNoRows {
    return nil, nil
  }

  defer stmtOut.Close()

  res := &defs.Vedio_info{Id: vid, AuthorId: aid, Name: name, DisPlayCTime: dispaly_ctime}
  return res, nil

}

// 删除视频接口
func DeleteVedioInfo(vid string) error {
  stmtDel, err := dbConn.Prepare("DELETE FROM vedio_info WHERE id = ?")
  if err != nil {
    return err
  }
  _, err = stmtDel.Exec(vid)
  if err != nil {
    return err
  }
  defer stmtDel.Close()
  return nil
}

///////////////////////////  评论接口   ////////////////

func AddNewComments(vid string, aid int, content string) error {
  id, err := utils.NewUUID()
  if err != nil {
    return err
  }
  stmtIns, err := dbConn.Prepare("INSERT INTO  comments (id,video_id,author_id,content) values (?,?,?,?)")
  if err != nil {
    return err
  }
  _, err = stmtIns.Exec(id, vid, aid, content)
  if err != nil {
    return err
  }
  defer stmtIns.Close()
  return nil
}

// 查询弹幕列表
func GetListComments(vid string, from, to int) ([] *defs.Comment, error) {
  stmtOut, err := dbConn.Prepare(`SELECT  comments.id,users.Login_name,comments.content
from comments INNER JOIN users ON comments.author_id = users.id WHERE  comments.video_id = ?
AND  comments.time > FROM_UNIXTIME(?) AND comments.time<=FROM_UNIXTIME(?)`)
  var res [] *defs.Comment

  rows, err := stmtOut.Query(vid, from, to)

  if err != nil {
    return res, err
  }
  for rows.Next() {
    var id, name, content string
    if err := rows.Scan(&id, &name, &content); err != nil {
      return res, err
    }
    c := &defs.Comment{Id: id, VideoId: vid, Author: name, Content: content}
    res = append(res, c)
  }
  defer stmtOut.Close()

  return res, nil

}

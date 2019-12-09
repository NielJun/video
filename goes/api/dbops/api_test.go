package dbops

import (
  "fmt"
  "strconv"
  "testing"
  "time"
)

// 测试过程  init(dblogin,truncate tables) ==>run tests ==> clear tables(truncate tables)

func clearTables() {
  dbConn.Exec("truncate users")
  dbConn.Exec("truncate vedio_info")
  dbConn.Exec("truncate comments")
  dbConn.Exec("truncate sessions")
}

func TestMain(m *testing.M) {
  clearTables()
  m.Run()
  clearTables()
}

// 测试用例的执行顺序流  AddUser --> GetUser -->DeletUser -->ReGetUser
// 总体执行顺序
// clear ---> test run ---> clear
func TestUserWorkFlow(t *testing.T) {
  t.Run("Add User", TestAddUserCredentail)
  t.Run("Get User", TestGetUserCredentail)
  t.Run("Delete User", TestDeleteUser)
  t.Run("Re Get User", testReGetUser)
}

func TestAddUserCredentail(t *testing.T) {

  err := AddUser("daniel", "123")
  if err != nil {
    t.Errorf("Error of Add User:%v", err)
  }
}

func TestGetUserCredentail(t *testing.T) {
  pwd, err := GetUserCredentail("daniel")
  if pwd != "123" || err != nil {
    t.Errorf("%v", "Error of Get User")
  }
}
func TestDeleteUser(t *testing.T) {
  err := DeleteUser("daniel", "123")
  if err != nil {
    t.Errorf("Error of Add User:%v", err)
  }
}

func testReGetUser(t *testing.T) {
  pwd, err := GetUserCredentail("daniel")
  if err != nil {
    t.Errorf("Error of ReGet User:%v", err)
  }
  if pwd != "" {
    t.Errorf("Delete Testing fiald err")
  }
}

/////////////////////////////   vedio test cases
func TestVedioWorkFlow(t *testing.T) {

  t.Run("PrepareUser", TestAddUserCredentail)
  t.Run("AddVedio", TestAddNewVedio)
  t.Run("GetVedio", TestGetVedioInfo)
  t.Run("DeleteVedio", TestDeleteVedioInfo)
  t.Run("RegetVedio", TestReGetVideoInfo)

}

var (
  tempvid string
)

func TestAddNewVedio(t *testing.T) {
  vedioInfo, err := AddNewVedio(1, "first-vedio")
  if err != nil {
    t.Errorf("Error of AddVedioInfo : %v", err)
  }
  tempvid = vedioInfo.Id
}

func TestGetVedioInfo(t *testing.T) {
  _, err := GetVedioInfo(tempvid)
  if err != nil {
    t.Errorf("Error of GetVideoInfo : %v", err)
  }
}

func TestDeleteVedioInfo(t *testing.T) {

  err = DeleteVedioInfo(tempvid)
  if err!=nil{
    t.Errorf("Error of DeleteVideoInfo : %v",err)
  }
}

func TestReGetVideoInfo(t *testing.T) {
  vid, err := GetVedioInfo(tempvid)
  if err != nil  && vid !=nil{
    t.Errorf("Error of ReGetVideoInfo : %v", err)
  }
}



////////////////// comments test cases

func TestComments(t *testing.T) {
  clearTables()
  t.Run("Add User",TestAddUserCredentail)
  t.Run("Add Comments",TestAddNewComments)
  t.Run("GetListComments",TestGetListComments)
}

func TestAddNewComments(t *testing.T) {
  vid := "12345"
  aid:=1
  content := "I love this vedio"
  err:=AddNewComments(vid,aid,content)
  if err!=nil{
    t.Errorf("Errof of AddNewComments: %v",err)
  }
}

func TestGetListComments(t *testing.T) {
  vid :="12345"
  from :=1514764800
  to,_:=strconv.Atoi(strconv.FormatInt(time.Now().UnixNano()/1000000000,10))
  res ,err :=GetListComments(vid,from,to)
  if err!=nil{
    t.Errorf("Error of ListComments: %v",err)
  }
  for i,ele :=range res{
    fmt.Printf("comment: %d, %v \n",i ,ele)
  }

}

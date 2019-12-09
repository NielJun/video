package defs

// add data models  与数据库数值相对应的model

type UserCredential struct {
  Username string `json:"user_name"`
  Pwd      string `json:"pwd"`
}

// response
type SignedUp struct {
  Success   bool   `json:"success"`
  SessionId string `json:"session_id"`
}

type Vedio_info struct {
  Id           string
  AuthorId     int
  Name         string
  DisPlayCTime string
}

type Comment struct {
  Id      string
  VideoId string
  Author  string
  Content string
}

type SimepleSession struct {
  Username string
  TTL      int64
}

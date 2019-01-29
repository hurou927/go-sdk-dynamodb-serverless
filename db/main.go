package main

import (
  "fmt"
  "github.com/aws/go-crypto/db/sub"
)
type User struct {
  UserId   string `json:"userId"`
  Username string `json:"username"`
  TenantId string `json:"tenantId"`
  Email    string `json:"email"`
  Status   int    `json:"status"`
  Password string `json:"password"`
}




func main()  {
  
  userdao, err := userDAO.NewUserDaoWithRegionAndEndpoint("usersTable", "ap-northeast-1", "http://localhost:8001");
  
  userdto, err := userdao.GetUserFromUserId("id0000");
  if err != nil{
    fmt.Println(err)
    return
  }
  fmt.Println(userdto)

  userdtos, err := userdao.GetUserFromTenantIdAndUsername("tenant0001", "name0001");
  if err != nil{
    fmt.Println(err)
    return
  }
  fmt.Println(userdtos)
}
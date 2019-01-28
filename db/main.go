package main

import (
  "fmt"
  "github.com/aws/aws-sdk-go-v2/aws"
//   "github.com/aws/aws-sdk-go-v2/aws/defaults"
  "github.com/aws/aws-sdk-go-v2/aws/endpoints"
  "github.com/aws/aws-sdk-go-v2/aws/external"
  "github.com/aws/aws-sdk-go-v2/service/dynamodb"
  "github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"
  // "github.com/aws/aws-sdk-go-v2/aws/awserr"

  // "strconv"
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
  cfg, err := external.LoadDefaultAWSConfig()
  if err != nil {
    panic("unable to load SDK config, " + err.Error())
  }
  cfg.EndpointResolver = aws.ResolveWithEndpointURL("http://localhost:8001")
  cfg.Region = endpoints.ApNortheast1RegionID
  cfg.DisableEndpointHostPrefix = true

  svc := dynamodb.New(cfg)

  data, err  := svc.GetItemRequest(&dynamodb.GetItemInput{
      Key: map[string]dynamodb.AttributeValue{
        "userId": {
          S: aws.String("id0001"),
        },
      },
      TableName: aws.String("usersTable"),
    }).Send()
  
  if err != nil{
    fmt.Println(err)
    return
  }

  
  // var user map[string]interface{}
  user1 := User{}
  if err :=dynamodbattribute.UnmarshalMap(data.Item, &user1); err!= nil{
    fmt.Println(err)
    return
  }
  fmt.Printf("%+v\n",user1)


  result, err := svc.QueryRequest(&dynamodb.QueryInput{
    TableName: aws.String("usersTable"),
    IndexName: aws.String("tenantIdUserNameIndex"),
    KeyConditionExpression: aws.String("username = :username and tenantId =:tenantId"),
    ExpressionAttributeValues: map[string] dynamodb.AttributeValue{
      ":username": {
        S: aws.String("name0001"),
      },
      ":tenantId": {
        S: aws.String("tenant0001"),
      },
    },
  }).Send()
  if err != nil{
    fmt.Println(err)
    return
  }

  users := []User{}
  if err :=dynamodbattribute.UnmarshalListOfMaps(result.Items, &users); err!= nil{
    fmt.Println(err)
    return
  }
  fmt.Printf("%+v\n",users)
}
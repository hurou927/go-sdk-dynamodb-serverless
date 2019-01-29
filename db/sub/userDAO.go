package userDAO

import (
  "fmt"
  "github.com/aws/aws-sdk-go-v2/aws"
  "github.com/aws/aws-sdk-go-v2/aws/external"
  "github.com/aws/aws-sdk-go-v2/service/dynamodb"
  "github.com/aws/aws-sdk-go-v2/service/dynamodb/dynamodbattribute"

)

type UserDto struct {
  UserId   string `json:"userId"`
  Username string `json:"username"`
  TenantId string `json:"tenantId"`
  Email    string `json:"email"`
  Status   int    `json:"status"`
  Password string `json:"password"`
}

type UserDao struct {
	tableName string
	dynamoClient *dynamodb.DynamoDB
}

func NewUserDaoDefaultConfig(tableName string)(*UserDao, error) {
	cfg, err := external.LoadDefaultAWSConfig()
  cfg.DisableEndpointHostPrefix = true
	if err != nil {
		return &UserDao{}, err
	}

  	return &UserDao{
			dynamoClient: dynamodb.New(cfg),
			tableName: tableName,
	}, nil
}

func NewUserDaoWithRegion(tableName string, region string)(*UserDao, error) {
	cfg, err := external.LoadDefaultAWSConfig()
	// cfg.EndpointResolver = aws.ResolveWithEndpointURL("http://localhost:8001")
  cfg.Region = region
  cfg.DisableEndpointHostPrefix = true
  if err != nil {
		return &UserDao{}, err
	}

  	return &UserDao{
			dynamoClient: dynamodb.New(cfg),
			tableName: tableName,
	}, nil
}

func NewUserDaoWithRegionAndEndpoint(tableName string, region string, endpoint string)(*UserDao, error) {
	cfg, err := external.LoadDefaultAWSConfig()
	cfg.EndpointResolver = aws.ResolveWithEndpointURL(endpoint)
  cfg.Region = region
  cfg.DisableEndpointHostPrefix = true
  if err != nil {
		return &UserDao{}, err
	}

  	return &UserDao{
			dynamoClient: dynamodb.New(cfg),
			tableName: tableName,
	}, nil
}



func (this *UserDao) GetUserFromUserId (userId string) (*UserDto, error) {
  if this == nil {
    return &UserDto{}, fmt.Errorf("nil pointer receiver")
  }
	result, err  := this.dynamoClient.GetItemRequest(&dynamodb.GetItemInput{
      Key: map[string]dynamodb.AttributeValue{
        "userId": {
          S: aws.String(userId),
        },
      },
      TableName: aws.String(this.tableName),
    }).Send()
  
  if result.Item == nil {
    return nil, nil
  }

  if err != nil{
    return &UserDto{}, err
  }
  fmt.Printf("%+v\n", result)
  user := UserDto{}
  if err :=dynamodbattribute.UnmarshalMap(result.Item, &user); err!= nil{
    return &UserDto{}, err
  }
  return &user, nil
}

func (this *UserDao) GetUserFromTenantIdAndUsername (tenantId string, username string) ([]UserDto, error) {
  if this == nil {
    return []UserDto{}, fmt.Errorf("nil pointer receiver")
  }
  result, err := this.dynamoClient.QueryRequest(&dynamodb.QueryInput{
    TableName: aws.String(this.tableName),
    IndexName: aws.String("tenantIdUserNameIndex"),
    KeyConditionExpression: aws.String("username = :username and tenantId =:tenantId"),
    ExpressionAttributeValues: map[string] dynamodb.AttributeValue{
      ":username": {
        S: aws.String(username),
      },
      ":tenantId": {
        S: aws.String(tenantId),
      },
    },
  }).Send()

  if err != nil{
    fmt.Println(err)
    return []UserDto{}, err
  }

  users := []UserDto{}
  if err :=dynamodbattribute.UnmarshalListOfMaps(result.Items, &users); err!= nil{
    fmt.Println(err)
    return []UserDto{}, err
  }
  return users, nil
}

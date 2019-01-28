package main

import (
    "fmt"
    "testing"
    "os"
    "../lib"
    "context"
)



func TestHandlerSuccess(t *testing.T) {
    os.Setenv("LAMBDA_TEST","123")
    body := map[string] interface{} {
		"key1" : "v1",
		"key2" : 1,
		"key3" : true,
	}

	queryParams := map[string] string {
		"query": "q1",
	}
	pathParams := map[string] string {
		"path": "p1",
    }
    

    request, err := request.CreateProxyRequest(body,queryParams, pathParams)
    
    var ctx context.Context
    response, err := Handler(ctx, request)
    if err != nil {
        t.Fatalf("failed test %#v", err)
    }
    fmt.Printf("%+v\n", response)
    // if result != 1 {
    //     t.Fatal("failed test")
    // }
}
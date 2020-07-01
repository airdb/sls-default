package main

import (
    "context"
    "fmt"
    "github.com/tencentyun/scf-go-lib/cloudfunction"
    "github.com/tencentyun/scf-go-lib/cloudevents/scf"
)

type DefineEvent struct {
    // test event define
    Key1 string `json:"key1"`
    Key2 string `json:"key2"`
}

func hello(ctx context.Context, req scf.APIGatewayProxyRequest) (string, error) {
    fmt.Println("req:", req)
    // return fmt.Sprintf("Hello %s!", req.Body), nil
    return fmt.Sprintf("Hello %s!", req.QueryString), nil
}

func main() {
    // Make the handler available for Remote Procedure Call by Cloud Function
    cloudfunction.Start(hello)
}

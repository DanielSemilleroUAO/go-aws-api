package main

import (
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/awslabs/aws-lambda-go-api-proxy/gorillamux"
	"github.com/gorilla/mux"
)

func Handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	r := mux.NewRouter()
	r.HandleFunc("/endpoint1", EndPoint1Handler).Methods("POST")
	r.HandleFunc("/endpoint2", EndPoint2Handler).Methods("GET")

	adapter := gorillamux.New(r)
	res, err := adapter.Proxy(req)
	if err != nil {
		fmt.Println(err.Error())
	}
	return res, err
}

func EndPoint1Handler(rw http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(rw, "Hola mundo")

}
func EndPoint2Handler(rw http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(rw, "Hola mundo 2")
}

func main() {
	lambda.Start(Handler)
}
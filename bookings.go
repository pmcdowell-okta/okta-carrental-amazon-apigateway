package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	//"encoding/json"
)


func bookingHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	return events.APIGatewayProxyResponse{
		Headers:    map[string]string{"content-type": "text/html"},
		Body:       bookings(),
		StatusCode: 200,
	}, nil
}

func main() {

	lambda.Start(bookingHandler)
}

func bookings() string {
	return ` {
                "vehicle_id":"733-23-13",
                "estimated_cost":"$ 120.00 USD",
                "confirmation_code":"HEUWIDWHDJIWY"
            }`
}


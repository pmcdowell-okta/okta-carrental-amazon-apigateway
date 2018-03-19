package main

import (
	"fmt"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	//"encoding/json"
)


func bookingHandler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	fmt.Println("Received body: ", request.Body)

	//datas := make(map[string]string)

	//datas["Dude where is my on echo2"]=request.Body

	//returnString, err := json.Marshal(datas)
	//_=returnString
	//_=err

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
                "vehicle_id":"{okta.in.id}",
                "estimated_cost":"{okta.estCost}",
                "confirmation_code":"HEUWIDWHDJIWY"
            }`
}


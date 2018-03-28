package main

import (
	//"fmt"
	"github.com/pmcdowell-okta/oktajwt"
	"context"
	//"errors"
	//"strings"
	//
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"strings"
)

func main() {
	lambda.Start(handleRequest)
}

func handleRequest(ctx context.Context, event events.APIGatewayCustomAuthorizerRequest) (events.APIGatewayCustomAuthorizerResponse, error) {

	token := string( event.AuthorizationToken)

	//token := `eyJraWQiOiJkZFBVRER5VXBIMk41d0dTWHZucVFaeS1PbVRGU1Z1NVBZYW5zanBzb0FzIiwiYWxnIjoiUlMyNTYifQ.eyJzdWIiOiIwMHUxOGVlaHUzNDlhUzJ5WDFkOCIsIm5hbWUiOiJva3RhcHJveHkgb2t0YXByb3h5IiwidmVyIjoxLCJpc3MiOiJodHRwczovL2NvbXBhbnl4Lm9rdGEuY29tIiwiYXVkIjoidlpWNkNwOHJuNWx4ck45YVo2ODgiLCJpYXQiOjE1MjIwOTM1NDQsImV4cCI6MTUyMjA5NzE0NCwianRpIjoiSUQuSEdiV1ZxWmJ0bFY1NUhDcFM4MlFHSEVmaVhGMWIxY0JGZE1ORVI5NF9vUSIsImFtciI6WyJwd2QiXSwiaWRwIjoiMDBveTc0YzBnd0hOWE1SSkJGUkkiLCJub25jZSI6Im4tMFM2X1d6QTJNaiIsInByZWZlcnJlZF91c2VybmFtZSI6Im9rdGFwcm94eUBva3RhLmNvbSIsImF1dGhfdGltZSI6MTUyMjA4MzkyNiwiYXRfaGFzaCI6IjZzZFNZWktUV3hkVFFMdXVLbkswU1EifQ.MGt2L3IGalqt0sd4wXYrHjgTG6Hly8IvgDvwZbZ40y-9OjhZBqkZduSboif2eqvocs66gJmJs9U4Gy1lcWTuEhJu4Hhb7nNsKEniuiv6dCt1L97OrSAMJjAvFm324Tp1kbgGocudX34CHx3RwBK0nWK8sfYgUo2_Ue7Zj0yTDSGN7qJ6640NVvxICp0QepLLXaPdSJSerqaisN8yZNhGnSrab5u347YBf-UfgLQ-4ewbeqcjhS_HyBg0zAhvzfZm52XTO4BovEx02dMSN6u_KNuprjaTT0jm3_11ldpagTujlPw6JwC1WRj-xzzffWWCAUtqWPT_V9X3OhWXEqEh-A`

	token = strings.Replace(token,"Bearer ", "", -1	)
	token = strings.Replace(token,"bearer ", "", -1	)

	result, err:=oktajwt.Oktaparsejwt( token )
	_=err
	if result!=nil { //Bad token
		return generatePolicy("user", "Deny", "*"), nil

	} else { //Good Token
		return generatePolicy("user", "Allow", "*"), nil
	}

}

// Help function to generate an IAM policy
func generatePolicy(principalId, effect, resource string) events.APIGatewayCustomAuthorizerResponse {
	authResponse := events.APIGatewayCustomAuthorizerResponse{PrincipalID: principalId}

	if effect != "" && resource != "" {
		authResponse.PolicyDocument = events.APIGatewayCustomAuthorizerPolicy{
			Version: "2012-10-17",
			Statement: []events.IAMPolicyStatement{
				{
					Action:   []string{"execute-api:Invoke"},
					Effect:   effect,
					Resource: []string{resource},
				},
			},
		}
	}

	return authResponse
}


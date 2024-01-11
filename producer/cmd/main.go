package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"producer/internal/app"
	"producer/internal/domain"
	"producer/internal/infrastructure/sqs"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type LambdaHandler struct {
	buyHandler *app.BuyHandler
}

func NewLambdaHandler() *LambdaHandler {
	queueURL := os.Getenv("QUEUE_URL")
	if queueURL == "" {
		log.Fatal("La variable de entorno QUEUE_URL es requerida")
	}

	sqsHandler := sqs.NewSQSHandler(queueURL)
	buyApp := app.NewBuyApplication(sqsHandler)
	buyHandler := app.NewBuyHandler(buyApp)
	return &LambdaHandler{
		buyHandler,
	}
}

func (h *LambdaHandler) HandleRequest(ctx context.Context, request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	b := []byte(request.Body)
	var body domain.Buy
	err := json.Unmarshal(b, &body)
	if err != nil {
		log.Fatal("Error Unmarshal:", err)
		return events.APIGatewayProxyResponse{
			StatusCode: 500,
		}, err
	}
	
	err = h.buyHandler.HandleBuyCreation(&body)
	if err != nil {
		log.Fatal("Error HandleBuyCreation:", err)

		return events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Internal Server Error",
		}, nil
	}

	response := domain.Response{
		Code:    200,
		Message: "buy: send success create",
	}
	r, _ := json.Marshal(response)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(r),
	}, nil
}

func main() {
	handler := NewLambdaHandler()
	lambda.Start(handler.HandleRequest)
}

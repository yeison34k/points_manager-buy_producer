package sqs

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"producer/internal/domain"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

type SQSHandler struct {
	SQSClient *sqs.SQS
	QueueURL  string
}

func NewSQSHandler(queueURL string) *SQSHandler {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-1"), // Cambia esto a tu región de AWS deseada
	})
	if err != nil {
		log.Fatal("Error creando la sesión:", err)
	}

	sqsClient := sqs.New(sess)

	return &SQSHandler{
		SQSClient: sqsClient,
		QueueURL:  queueURL,
	}
}

func (h *SQSHandler) CreateBuy(buy *domain.Buy) error {
	id, err := exec.Command("uuidgen").Output()
	if err != nil {
		log.Fatal("fallo al crear el uuid: %w", err)
	}
	buy.ID = string(id)
	buyJSON, err := json.Marshal(buy)
	if err != nil {
		return fmt.Errorf("fallo al convertir el punto a JSON: %w", err)
	}

	sendMessageInput := &sqs.SendMessageInput{
		MessageBody:  aws.String(string(buyJSON)),
		QueueUrl:     aws.String(h.QueueURL),
		DelaySeconds: aws.Int64(0),
	}

	_, err = h.SQSClient.SendMessage(sendMessageInput)
	if err != nil {
		return fmt.Errorf("fallo al enviar el mensaje a SQS: %w", err)
	}

	return nil
}

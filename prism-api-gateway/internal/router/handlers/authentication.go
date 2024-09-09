package handlers

import (
	"log"
	"prism-api-gateway/internal/producer"

	"github.com/IBM/sarama"
	"github.com/gofiber/fiber/v2"
)

func sendAuthRequest(email, password string) error {
	msg := &sarama.ProducerMessage{
		Topic: "register_request",
		Value: sarama.StringEncoder(email + " " + password),
	}

	partition, offset, err := producer.Get().SendMessage(msg)
	if err != nil {
		return err
	}

	log.Printf("Message sent to partition %d at offset %d\n", partition, offset)

	return nil
}

func SignUp(c *fiber.Ctx) error {

	type SignupRequest struct {
		Login    string `json:"email"`
		Password string `json:"password"`
	}

	var req SignupRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	err := sendAuthRequest(req.Login, req.Password)
	if err != nil {
		return c.JSON(fiber.Map{
			"status": "failed",
			"msg":    "Failed to send request to kafka",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"msg":    "account success created",
	})
}

func SignIn(c *fiber.Ctx) error {
	type SignupRequest struct {
		Login    string `json:"email"`
		Password string `json:"password"`
	}

	var req SignupRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
	}

	err := sendAuthRequest(req.Login, req.Password)
	if err != nil {
		return c.JSON(fiber.Map{
			"status": "failed",
			"msg":    "Failed to send request to kafka",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"msg":    "account success authorized",
	})
}

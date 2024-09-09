package authentication

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type ResultObj struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

func SignIn(c *fiber.Ctx) error {
	return c.Render("signin", fiber.Map{})
}

func LogIn(c *fiber.Ctx) error {
	email, password := c.FormValue("e"), c.FormValue("p")

	requestBody, err := json.Marshal(map[string]string{
		"email":    email,
		"password": password,
	})

	if err != nil {
		return c.Redirect("/sign_in?err=Failed to get data from client")
	}

	resp, err := http.Post("http://localhost:8080/api/v1/sign_in", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return c.Redirect("/sign_in?err=Failed to get data from server")
	}

	defer resp.Body.Close()
	fmt.Println(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return c.Redirect("/sign_in?err=Failed to read data from server")
	}

	result := ResultObj{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return c.Redirect("/sign_in?err=Failed to read data from response")
	}

	fmt.Println(result.Status)
	fmt.Println(result.Msg)
	return c.Redirect("/dashboard")
}

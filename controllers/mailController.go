package controllers

import (
	"fmt"

	"github.com/1deep1/deepcraft-backend/models"
	"github.com/gofiber/fiber/v2"
	gomail "gopkg.in/mail.v2"
)

func Mail(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	mail := models.Mail{
		Name:    data["name"],
		Phone:   data["phone"],
		Message: data["message"],
	}

	m := gomail.NewMessage()

	m.SetHeader("From", "no-reply@1deep1.ru")
	m.SetHeader("To", "mail@1deep1.ru")
	m.SetHeader("Subject", "Новый заказ!")
	m.SetBody("text/html", fmt.Sprintf("<h1>Новый заказ!</h1><h3>1deep1.ru</h3><br><p><b>Имя:</b> %s</p><p><b>Телефон:</b> %s</p><p><b>Сообщение:</b> %s</p>", mail.Name, mail.Phone, mail.Message))

	a := gomail.NewDialer("smtp.yandex.ru", 465, "no-reply@1deep1.ru", "Tishina20024524")

	if err := a.DialAndSend(m); err != nil {
		return err
	}

	return c.JSON(fiber.Map{
		"message": "success",
	})
}

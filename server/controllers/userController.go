package controllers

import (
	"context"

	"../models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["name"],
		password: password,
	}

	_, err = database.userCollection.InsertOne(context.Background(), user)

	if err != nil {
		return err
	}

	return c.JSON(user)

}

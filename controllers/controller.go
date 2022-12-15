package controllers

import (
	"encoding/json"
	"log"
	"strings"

	"github.com/senpainikolay/CS-sem5/models"

	"gorm.io/gorm"

	"github.com/pquerna/otp/totp"
)

type AuthController struct {
	DB *gorm.DB
}

func NewAuthController(DB *gorm.DB) AuthController {
	return AuthController{DB}
}

func (ac *AuthController) SignUpUser(ctx models.RegisterUserInput) string {

	newUser := models.User{
		Name:     ctx.Name,
		Email:    strings.ToLower(ctx.Email),
		Password: ctx.Password,
	}
	result := ac.DB.Create(&newUser)

	if result.Error != nil && strings.Contains(result.Error.Error(), "duplicate key value violates unique") {
		return "email already exists!"
	} else if result.Error != nil {
		return result.Error.Error()
	}

	return "Succesfully Registration"
}

func (ac *AuthController) LoginUser(ctx models.LoginUserInput) string {

	var user models.User
	result := ac.DB.First(&user, "email = ?", strings.ToLower(ctx.Email))
	if result.Error != nil {
		return "Invalid password or Email"
	}
	return "Succesful Log In"
}

func (ac *AuthController) GenerateOTP(ctx models.OTPInput) string {

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "MESSAGE_SENT_OTP_TEST:",
		AccountName: "ACCOUNT@mail.com",
		SecretSize:  15,
	})

	if err != nil {
		panic(err)
	}

	var user models.User
	result := ac.DB.First(&user, "id = ?", ctx.UserId)
	if result.Error != nil {
		return "Invalid password or Email"
	}

	dataToUpdate := models.User{
		Otp_secret:   key.Secret(),
		Otp_auth_url: key.URL(),
	}

	ac.DB.Model(&user).Updates(dataToUpdate)

	resp, _ := json.Marshal(user)

	return string(resp)
}

func (ac *AuthController) VerifyOTP(ctx models.OTPInput) string {

	message := "Token is invalid or user doesn't exist"

	var user models.User
	result := ac.DB.First(&user, "id = ?", ctx.UserId)
	if result.Error != nil {
		return result.Error.Error()
	}
	log.Println(ctx.Token)
	log.Println(user.Otp_secret)

	valid := ctx.Token == user.Otp_secret
	if !valid {
		return message
	}

	dataToUpdate := models.User{
		Otp_enabled:  true,
		Otp_verified: true,
	}

	ac.DB.Model(&user).Updates(dataToUpdate)

	resp, _ := json.Marshal(user)

	return string(resp)
}

func (ac *AuthController) ValidateOTP(ctx models.OTPInput) string {

	message := "Token is invalid or user doesn't exist"

	var user models.User
	result := ac.DB.First(&user, "id = ?", ctx.UserId)
	if result.Error != nil {
		return result.Error.Error()
	}

	valid := ctx.Token == user.Otp_secret
	if !valid {
		return message
	}

	resp, _ := json.Marshal(user)
	return string(resp)
}

func (ac *AuthController) DisableOTP(ctx models.OTPInput) string {

	var user models.User
	result := ac.DB.First(&user, "id = ?", ctx.UserId)
	if result.Error != nil {
		return result.Error.Error()
	}

	user.Otp_enabled = false
	ac.DB.Save(&user)

	resp, _ := json.Marshal(user)
	return string(resp)

}

package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zuramai/crypto-price-tracker/internal/config"
	"github.com/zuramai/crypto-price-tracker/internal/model"
	"github.com/zuramai/crypto-price-tracker/internal/utils"
)

var baseUrl string
var registeredUserEmail string
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var token string

func randSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
func init() {
	viper := config.NewViper()
	appPort := viper.GetString("app.port")
	registeredUserEmail = randSeq(10) + "@gmail.com"
	baseUrl = "http://localhost:" + appPort + "/api"
}

func TestRegisterPasswordConfirmationDoesNotMatch(t *testing.T) {
	payload := []byte(fmt.Sprintf(`{"email": "%s", "password": "password", "password_confirmation": "password1"}`, registeredUserEmail))
	res, err := http.Post(baseUrl+"/auth/register", "application/json", bytes.NewBuffer(payload))
	assert.NoError(t, err)
	assert.Equal(t, res.StatusCode, 422)

	body, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	var response utils.ApiResponseMessage
	json.Unmarshal(body, &response)

	assert.Equal(t, response.Message, "Password confirmation doesn't match")
}

func TestRegisterSuccess(t *testing.T) {
	payload := []byte(fmt.Sprintf(`{"email": "%s", "password": "password", "password_confirmation": "password"}`, registeredUserEmail))
	res, err := http.Post(baseUrl+"/auth/register", "application/json", bytes.NewBuffer(payload))
	assert.NoError(t, err)
	assert.Equal(t, res.StatusCode, 200)

	body, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	var response utils.ApiResponseMessage
	json.Unmarshal(body, &response)

	assert.Equal(t, response.Message, "Register Success")
}

func TestLoginFailed(t *testing.T) {
	payload := []byte(`{"email": "randomemail@gmail.com", "password": "password"}`)
	res, err := http.Post(baseUrl+"/auth/login", "application/json", bytes.NewBuffer(payload))
	assert.NoError(t, err)
	assert.Equal(t, res.StatusCode, 401)

	body, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	var response utils.ApiResponseMessage
	json.Unmarshal(body, &response)

	assert.Equal(t, response.Message, "Invalid credentials")
}
func TestLoginSuccess(t *testing.T) {
	payload := []byte(fmt.Sprintf(`{"email": "%s", "password": "password"}`, registeredUserEmail))
	res, err := http.Post(baseUrl+"/auth/login", "application/json", bytes.NewBuffer(payload))
	assert.NoError(t, err)
	assert.Equal(t, res.StatusCode, 200)

	body, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	var response utils.ApiResponse[model.AuthSuccessResponse]
	json.Unmarshal(body, &response)
	token = response.Data.Token

	assert.Equal(t, response.Message, "Login Success")
}
func TestLogoutFailed(t *testing.T) {
	req, err := http.NewRequest("POST", baseUrl+"/auth/logout", nil)
	assert.NoError(t, err)

	client := &http.Client{}
	res, err := client.Do(req)

	assert.NoError(t, err)
	assert.Equal(t, res.StatusCode, 401)

	body, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	var response utils.ApiResponse[model.AuthSuccessResponse]
	json.Unmarshal(body, &response)

	assert.Equal(t, response.Message, "Unauthorized")
}

func TestLogoutSuccess(t *testing.T) {
	req, err := http.NewRequest("POST", baseUrl+"/auth/logout", nil)
	assert.NoError(t, err)

	req.Header.Set("Authorization", "Bearer "+token)
	client := &http.Client{}
	res, err := client.Do(req)

	assert.NoError(t, err)
	assert.Equal(t, res.StatusCode, 200)

	body, err := io.ReadAll(res.Body)
	assert.NoError(t, err)

	var response utils.ApiResponse[model.AuthSuccessResponse]
	json.Unmarshal(body, &response)

	assert.Equal(t, response.Message, "Logout Success")
}

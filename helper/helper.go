package helper

import (
	"math/rand"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/lib/pq"
)

type Response struct {
	Status 		bool 			`json:"status"`
	Message 	string 			`json:"message"`
	Data		interface{}		`json:"data"`
}

func APIResponse(status bool, message string, data interface{}) Response {
	jsonResponse := Response{
		Status: status,
		Message: message,
		Data: data,
	}

	return jsonResponse
}

func FormatValidationError(err error) []string {
	var errors []string

	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}

	return errors
}

func IsDatabaseError(err error) bool {
	if _, ok := err.(*pq.Error); ok {
		return true
	}

	return false
}

func GenerateAccountNumber() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomInt := r.Intn(10000000000000000)
	
	randomIntToString := strconv.Itoa(randomInt)
	return randomIntToString
}
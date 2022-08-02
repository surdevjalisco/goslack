package goslack

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Implementación de llamadas a las API de slack

type userProfile struct {
	Email string `json:"email"`
}

type user struct {
	Profile    userProfile `json:"profile"`
	ValidEmail bool        `json:"is_email_confirmed"`
}

type members struct {
	Users []user `json:"members"`
}

const slackAPI string = "https://slack.com/api"
const slackAuthEnvVar string = "GOSLACK_AUTH_TOKEN"

func getAuthFromEnv() (string, error) {
	value := os.Getenv(slackAuthEnvVar)
	if value == "" {
		return "", fmt.Errorf("no se encontró token de autentificación (%s) en el ambiente ", slackAuthEnvVar)
	}
	return value, nil
}

func GetUsersEmails() ([]string, error) {

	fmt.Println("Obteniendo lista de correos...")
	request, err := http.NewRequest(http.MethodGet,
		fmt.Sprintf("%s/users.list", slackAPI),
		nil)
	if err != nil {
		return nil, fmt.Errorf("creando petición para API (request)  %w", err)
	}

	token, err := getAuthFromEnv()
	if err != nil {
		return nil, err
	}
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("enviando petición: %w", err)
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("leyendo respuesta: %w", err)
	}

	var slackMembers members
	if err = json.Unmarshal(body, &slackMembers); err != nil {
		return nil, fmt.Errorf("creando modelo: %w", err)
	}

	var emailsResult = []string{}
	for _, member := range slackMembers.Users {
		if member.ValidEmail {
			emailsResult = append(emailsResult, member.Profile.Email)
		}
	}

	return emailsResult, nil
}

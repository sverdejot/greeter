package http

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type UserService struct {
	client *http.Client
}

func NewUserService() *UserService {
	return &UserService{
		http.DefaultClient,
	}
}

func (s *UserService) GetUserName(id int) (string, bool) {
	resp, err := s.client.Get(fmt.Sprintf("http://users:8081/users/%d", id))

	if err != nil {
		log.Print(err)
		return "", false
	}

	body := &struct {
		Name string `json:"name"`
	}{}

	bodyStr, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		log.Print(err)
		return "", false
	}

	err = json.Unmarshal(bodyStr, body)
	if err != nil {
		log.Print(err)
		return "", false
	}

	return body.Name, true
}

package jsonresponse

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User map[string]any

type Scraper interface {
	Fetch()
}

type UsersFetcher struct {
	URL      string
	response *http.Response
}

func (u *UsersFetcher) Fetch() error {
	response, err := http.Get(u.URL)
	if err != nil {
		return err
	}

	u.response = response

	return nil
}

func (u *UsersFetcher) GetUsers() ([]User, error) {
	defer u.response.Body.Close()

	body, readErr := io.ReadAll(u.response.Body)
	if readErr != nil {
		return nil, readErr
	}

	var users []User
	jsonErr := json.Unmarshal(body, &users)
	if jsonErr != nil {
		return nil, jsonErr
	}

	fmt.Println(users[0])
	return users, nil
}

func (u *UsersFetcher) GetOneUser() (User, error) {
	defer u.response.Body.Close()

	body, readErr := io.ReadAll(u.response.Body)
	if readErr != nil {
		return nil, readErr
	}

	var user User
	jsonErr := json.Unmarshal(body, &user)
	if jsonErr != nil {
		return nil, jsonErr
	}

	// TODO: Try to get value of inner JSON key (ex: user.address.geo)
	fmt.Println(user["address"])

	return user, nil
}

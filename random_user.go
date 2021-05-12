package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ApiResponce struct {
	Results []User `json:"results"`
	Info    Info   `json:"info"`
	Status  Status `json:"-"`
}

type Status struct {
	Code int
	Text string
}

type Info struct {
	Seed    string `json:"seed"`
	Version string `json:"version"`
}

type User struct {
	Gender  string   `json:"gender"`
	Name    UserName `json:"name"`
	Email   string   `json:"email"`
	Phone   string   `json:"cell"`
	Picture Picture  `json:"picture"`
}

type UserName struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

type Picture struct {
	URL string `json:"large"`
}

func getRandomUser() (*ApiResponce, error) {
	resp, err := http.Get(RandomUserApiUrl)
	if err != nil {
		return nil, err
	}

	apiResp := &ApiResponce{Status: formateStatus(resp.StatusCode)}
	if resp.StatusCode != http.StatusOK {
		return apiResp, nil
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, apiResp)
	if err != nil {
		return nil, err
	}

	return apiResp, nil
}

func formateStatus(code int) Status {
	return Status{
		Code: code,
		Text: http.StatusText(code),
	}
}

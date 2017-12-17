package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

// var host = "https://private-5b97e-agenda29.apiary-mock.com"

var host = "http://localhost:8080"

func Register(username string, password string, email string, phone string) error {
	data := struct {
		Username string `json:"username"`
		Password string `json:"password"`
		Email    string `json:"email"`
		Phone    string `json:"phone"`
	}{username, password, email, phone}
	buf, _ := json.Marshal(data)
	res, err := http.Post(host+"/v1/user/register",
		"application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != 201 {
		fmt.Println(string(body))
		return errors.New("Register fail")
	} else {
		fmt.Println("Register success as: ")
		fmt.Println(string(body))
		return nil
	}
}

func Login(username string, password string) error {
	res, err := http.Get(host + "/v1/user/login?username=" +
		username + "&password=" + password)
	if err != nil {
		return err
	}
	body, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != 204 {
		fmt.Println(string(body))
		return errors.New("Login fail")
	} else {
		fmt.Println("Login: success.")
		return nil
	}
}

func Logout() error {
	res, _ := http.Get(host + "/v1/user/logout")
	body, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != 204 {
		fmt.Println(string(body))
		return errors.New("Logout fail")
	} else {
		fmt.Println("Logout: success.")
		return nil
	}
}

func ListUsers() error {
	res, err := http.Get(host + "/v1/users")
	if err != nil {
		return err
	}
	body, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 {
		fmt.Println(string(body))
		return errors.New("List user fail")
	} else {
		fmt.Printf(string(body))
		return nil
	}
}

func CreateMeeting(title string, participators_str string, startTime string, endTime string) error {
	// members := util.Str2slice(participators_str)
	members := strings.Fields(participators_str)
	data := struct {
		Title         string   `json:"Title"`
		Participators []string `json:"Participators"`
		Starttime     string   `json:"StartTime"`
		Endtime       string   `json:"EndTime"`
	}{title, members, startTime, endTime}
	buf, _ := json.Marshal(data)

	res, err := http.Post(host+"/v1/meeting",
		"application/json", bytes.NewBuffer(buf))
	if err != nil {
		return err
	}
	body, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != 201 {
		fmt.Println(string(body))
		return errors.New("Create meeting fail")
	} else {
		fmt.Println("Creat meeting success as:")
		fmt.Println(string(body))
		return nil
	}
}

func ListMeetings(startTime, endTime string) error {
	res, err := http.Get(host + "/v1/meetings")
	if err != nil {
		return err
	}
	body, _ := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 {
		fmt.Println(string(body))
		return errors.New("List meetings fail")
	} else {
		fmt.Printf(string(body))
		return nil
	}
}

package main

import (
	"errors"
	"fmt"
	"go_lessons/pointers"
	"net/url"
	"time"
)

type accaunt struct {
	login    string
	password string
	url      string
}

type accountWithTime struct {
	createdAt time.Time
	updatedAt time.Time
	acc       accaunt
}

func main() {
	slice := make([]int, 3, 3)
	pointers.ReversedSlice(slice)
	accaunt1 := accaunt{
		login:    "login111",
		password: "password22",
		url:      "url222",
	}

	accaunt1.setPassword("12345")
	accaunt1.printAccount()

	account2, err := newAcccaunt("admin", "1257", "http://a.net")
	if err != nil {
		fmt.Println("Ошибка создания struct = ", err.Error())
	}
	fmt.Println(*account2)

	accountWithTime, err := newAccauntWithTime()
	fmt.Println(accountWithTime)
}

func (acc *accaunt) printAccount() {
	fmt.Println("Account = ", *acc)
}

func (acc *accaunt) setPassword(password string) {
	acc.password = password
}

func newAcccaunt(login, password, urlStr string) (*accaunt, error) {
	_, err := url.ParseRequestURI(urlStr)
	if err != nil {
		return nil, errors.New("Invalid URL")
	}
	return &accaunt{
		login:    login,
		password: password,
		url:      urlStr,
	}, nil
}

func newAccauntWithTime() (*accountWithTime, error) {
	return &accountWithTime{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		acc: accaunt{
			login:    "login123",
			password: "password123",
			url:      "https://ya.org",
		},
	}, nil
}

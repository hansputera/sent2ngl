package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

// TODO: get ig username.

func get_url(dst string) string {
	return fmt.Sprintf("https://ngl.link/%s", dst)
}

func get_device_token(body *[]byte) *string {
	tokens := regexp.MustCompile(`;token=([a-zA-Z0-9]+-[a-zA-Z0-9]+-[a-zA-Z0-9]+-[a-zA-Z0-9]+-([a-zA-Z0-9]+)?)`).FindAllStringSubmatch(string(*body), -1)
	if len(tokens) != 1 {
		return nil
	}

	return &tokens[0][len(tokens[0])/2]
}

func Sent(dst *string, msg *string) (*string, error) {
	res, err := http.Get(get_url(*dst))
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	token := get_device_token(&body)
	if token == nil {
		return nil, errors.New("user not found")
	}

	res, err = http.PostForm(get_url(*dst), url.Values{
		"question": []string{*msg},
		"deviceId": []string{*token},
	})

	if err != nil {
		return nil, err
	}

	res.Body.Close()

	return token, nil
}

package sae

import (
	"net/http"
)

type Client struct {
	ClientId     string
	ClientSecret string
	accessToken  string
	refreshToken string
}

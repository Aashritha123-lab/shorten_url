package models

import (
	"database/sql"
	"fmt"
	"url/config"
)

type ShortenRequest struct {
	Url string
}

type Response struct {
	ResponseUrl    string
	ShortenRequest ShortenRequest
}

func (url *Response) Insert() error {

	query := `INSERT INTO short_url(short_response,url)VALUES($1,$2)RETURNING short_response,url`

	if err := config.DB.QueryRow(query, url.ResponseUrl, url.ShortenRequest.Url).Scan(&url.ResponseUrl, &url.ShortenRequest.Url); err != nil {
		fmt.Printf("Error inserting to database :%v\n", err)
		return err
	}
	return nil
}

func (url *Response) Get() error {
	query := `SELECT short_response,url FROM short_url where short_response = $1`
	fmt.Println("url.ResponseUrl in get method is:", url.ResponseUrl)
	if err := config.DB.QueryRow(query, url.ResponseUrl).Scan(&url.ResponseUrl, &url.ShortenRequest.Url); err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("Error in getting the url :%v\n", err)
			return err
		}
	}
	return nil
}

func (url *Response) UrlValidate() bool {
	// 1. Check if the original URL already exists
	var existingUrl string
	query := `SELECT url FROM short_url WHERE url = $1 LIMIT 1`
	err := config.DB.QueryRow(query, url.ShortenRequest.Url).Scan(&existingUrl)
	if err == nil {
		fmt.Println("URL already exists:", existingUrl)
		return true
	}
	if err != sql.ErrNoRows {
		fmt.Println("DB error while checking URL:", err)
		return true // fail safe
	}
	return false
}
func (url *Response) CodeValidate() bool {
	// 2. Check if the short code already exists
	var existingShort string
	query := `SELECT short_response FROM short_url WHERE short_response = $1 LIMIT 1`
	err := config.DB.QueryRow(query, url.ResponseUrl).Scan(&existingShort)
	if err == nil {
		fmt.Println("Short code already exists:", existingShort)
		return true
	}
	if err != sql.ErrNoRows {
		fmt.Println("DB error while checking short code:", err)
		return true
	}

	// No conflicts → safe to insert
	return false
}

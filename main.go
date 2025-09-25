package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"url/config"
	"url/handlers"
	"url/middleware"
	"url/utils"

	_ "github.com/lib/pq"
)

func main() {
	config.ConnectDB()
	utils.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	mux := http.NewServeMux()
	mux.HandleFunc("/shorten", handlers.ShorthandUrl)
	mux.HandleFunc("/", handlers.Redirect)

	wrappedMux := middleware.Logger(mux)
	wrappedMux = middleware.SecurityHeader(wrappedMux)

	fmt.Println("Http server is listening....")
	http.ListenAndServe(":3051", wrappedMux)

}

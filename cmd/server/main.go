package main

import (
	"github.com/AnhducNA/go-ecommerce/internal/routers"
)

func main() {
	router := routers.NewRouter()
	router.Run(":5000") // listen and serve on 0.0.0.0:8080
}

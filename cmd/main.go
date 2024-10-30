package main

import (
	"adobe-ripper/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.Fire(r)

	r.Run("127.0.0.1:5741")
}

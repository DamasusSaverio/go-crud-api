package main

import (
	"github.com/damasussaverio/go-crud-gin/initializers"
	"github.com/damasussaverio/go-crud-gin/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}

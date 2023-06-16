package controllers

import (
	"github.com/damasussaverio/go-crud-gin/initializers"
	"github.com/damasussaverio/go-crud-gin/models"
	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	//Get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//Create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	//Get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	// Return it
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// Get the id off URL
	id := c.Param("id")

	//Get the single post
	var post models.Post
	initializers.DB.First(&post, id)

	// Return it
	c.JSON(200, gin.H{
		"posts": post,
	})
}

func PostsUpdate(c *gin.Context) {
	// Get the id off URL
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	// Find the post we're updating
	var post models.Post
	initializers.DB.First(&post, id)

	// Update it
	initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	// Return it
	c.JSON(200, gin.H{
		"posts": post,
	})

}

func PostsDelete(c *gin.Context) {
	// get the id off the url
	id := c.Param("id")

	// delete the posts
	initializers.DB.Delete(&models.Post{}, id)

	// return it
	c.Status(200)

}

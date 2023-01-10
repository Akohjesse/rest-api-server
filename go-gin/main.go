package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Album struct {
	ID string  `json:"id"`
	Title string  `json:"title"`
	Artist string  `json:"artist"`
	Price float64   `json:"price"`
}

func getAlbums(c *gin.Context){
    c.IndentedJSON(http.StatusOK, albums)
}

func postAlbums(c *gin.Context){
	var newalbum Album
	err:= c.BindJSON(&newalbum);
	if err!= nil {
		return 
	}
	albums =  append(albums, newalbum)
	c.IndentedJSON(http.StatusCreated, newalbum)
}

func getAlbumById(c *gin.Context){
  id := c.Param("id")
  for _,a := range albums {
	if a.ID == id {
		c.IndentedJSON(http.StatusOK, a)
		return
	} 
  }
  c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbums)
	router.Run("localhost:7000")
	log.Println("Server started at port 7000")
}
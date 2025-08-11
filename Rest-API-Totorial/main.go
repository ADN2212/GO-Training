package main

import (
	"fmt"
    "net/http"
    "github.com/gin-gonic/gin"
	"os"
	"github.com/joho/godotenv"//Si este paquete las variables no cargan.
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`//Estas etiquetas determinan la clave en el JSON.
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float64 `json:"price"`
}

// albums slice to seed record album data.
//Esta es la "Base de Datos" del tutoral.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
    {ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main() {

	err := godotenv.Load(".env")
	
	if err != nil {
		fmt.Println("Envars could not be loeaded")
		return
	}
	
	port := os.Getenv("PORT")
	
	if len(port) == 0 {
		fmt.Println("PORT not found, setting to 8080")
		port = "8080"
	}

    router := gin.Default()
    router.GET("/albums", getAlbums)
    router.POST("/albums", postAlbum)
	router.GET("/albums/:id", getAlbumByID)
	router.Run(fmt.Sprintf("localhost:%s", port))
}

func getAlbumByID(c *gin.Context) {
	id := c.Param("id")//Para obtener url params.

	for _, al := range albums {
		if al.ID == id {
			c.IndentedJSON(http.StatusFound, al)
			return
		}
	}
	
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found :( "})
}

//Como en JS, Go hace hoisting.
// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(c *gin.Context) {

	var newAlbum album

	err := c.BindJSON(&newAlbum)//Esta funcion toma el puntero a newAlbum y 'vuelca' la info del JSON en la struct 
	if err != nil {
		fmt.Println(err)//Pendiente ver los errores que da gin.
		return
	}

	albums = append(albums, newAlbum)
    //Envia la response:
	c.IndentedJSON(http.StatusCreated, newAlbum.ID)
}

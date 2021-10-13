package controllers

import (
	"log"
	"net/http"

	"github.com/gabrielsodre91/api-gin/database"
	"github.com/gabrielsodre91/api-gin/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAllBooks(c *gin.Context) {
	var books []models.Book

	cur, err := database.MongoDatabase.Collection("books").Find(c, bson.M{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error at query Books collection."})
	}

	defer cur.Close(c)

	for cur.Next(c) {
		var book models.Book

		err := cur.Decode(&book)

		if err != nil {
			log.Fatal(err)
		}

		books = append(books, book)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}


	c.JSON(http.StatusOK, books)
}

func GetBook(c *gin.Context) {
	id := c.Param("id")

	var book models.Book

	objectID, _ := primitive.ObjectIDFromHex(id)

	err := database.MongoDatabase.Collection("books").FindOne(c, bson.M{"_id": objectID}).Decode(&book)
	if err != nil {
		log.Println("Book not found!")

		c.JSON(http.StatusBadRequest, gin.H{ "msg": "Book not found!" })

		return
	}

	c.JSON(http.StatusOK, book)
}

func UpdateBook(c *gin.Context) {
	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	filter := bson.M{"_id": book.ID}

	update := bson.D{
		{"$set", bson.D{
			{"title", book.Title},
			{"author", bson.D{
				{"firstname", book.Author.FirstName},
				{"lastname", book.Author.LastName},
			}},
		}},
	}

	errUpdate := database.MongoDatabase.Collection("books").FindOneAndUpdate(c, filter, update).Decode(&book)
	if errUpdate != nil {
		log.Println("Error at update book!")

		c.JSON(http.StatusBadRequest, gin.H{ "msg": "Error at update book!" })

		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Book updated."})
}

func CreateBook(c *gin.Context) {
	var book models.Book

	err := c.ShouldBindJSON(&book)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})
		return
	}

	_, errInsert := database.MongoDatabase.Collection("books").InsertOne(c, book)
	if errInsert != nil {
		log.Println("Error at insert book!")

		c.JSON(http.StatusBadRequest, gin.H{ "msg": "Error at insert book!" })

		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Book insertd."})
}

func DeleteBook(c *gin.Context) {
	id := c.Param("id")

	objectID, _ := primitive.ObjectIDFromHex(id)

	_, err := database.MongoDatabase.Collection("books").DeleteOne(c, bson.M{"_id": objectID})
	if err != nil {
		log.Println("Error at delete book!")

		c.JSON(http.StatusBadRequest, gin.H{ "msg": "Error at delete book!" })

		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Book deleted!"})
}
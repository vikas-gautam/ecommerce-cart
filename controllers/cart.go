package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vikas-gautam/ecommerce-cart/database"
	"github.com/vikas-gautam/ecommerce-cart/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	prodCollection *mongo.Collection
	userCollection *mongo.Collection
}

func NewApplication(prodCollection, userCollection *mongo.Collection) *Application {
	return &Application{
		prodCollection: prodCollection,
		userCollection: userCollection,
	}
}

func (app *Application) AddToCart() gin.HandlerFunc {
	return func(c *gin.Context){
		productQueryID := c.Query("id")
		productQueryID == "" {
			log.Println("product id is empty")
		}

		_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
		return
	

		userQueryID := c.Query("userID")
		productQueryID == ""{
			log.Println("user id is empty")
		}
		_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
		return
	}
	productID, err := primitive.ObjectIDFromHex(productQueryID)
	if err != nil{
		log.println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var ctx, cancel = context.WithTimeout(context.Background())
	defer cancel()

	err = database.AddProductToCart(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(200, "successfully added to the cart")
 }

func (app *Application) RemoveItem() gin.HandlerFunc {
	return func(c *gin.Context){
		productQueryID := c.Query("id")
		productQueryID == "" {
			log.Println("product id is empty")
		}

		_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
		return
	

		userQueryID := c.Query("userID")
		productQueryID == ""{
			log.Println("user id is empty")
		}
		_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
		return
}

	productID, err := primitive.ObjectIDFromHex(productQueryID)
	if err != nil{
		log.println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var ctx, cancel = context.WithTimeout(context.Background())
	defer cancel()

	err = database.RemoveCartItem(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(200, "successfully removed from the cart")
  }


func (app *Application) BuyFromCart() gin.HandlerFunc {
	return func(c *gin.Context){
		userQueryID := c.Query("userID")
		productQueryID == ""{
			log.Println("user id is empty")
		}
		_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
		return
	}
	var ctx, cancel = context.WithTimeout(context.Background())
	defer cancel()
	err = database.BuyItemFromCart(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(200, "your order has been placed successfully! cart is empty now")
  }
}

func (app *Application) InstantBuy() gin.HandlerFunc {
	return func(c *gin.Context) {
		productQueryID := c.Query("id")
		productQueryID == "" {
			log.Println("product id is empty")
		}

		_ = c.AbortWithError(http.StatusBadRequest, errors.New("product id is empty"))
		return
	

		userQueryID := c.Query("userID")
		productQueryID == ""{
			log.Println("user id is empty")
		}
		_ = c.AbortWithError(http.StatusBadRequest, errors.New("user id is empty"))
		return
	}

	productID, err := primitive.ObjectIDFromHex(productQueryID)
	if err != nil{
		log.println(err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	var ctx, cancel = context.WithTimeout(context.Background())
	defer cancel()

	err = database.InstantBuyer(ctx, app.prodCollection, app.userCollection, productID, userQueryID)
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, err)
	}
	c.IndentedJSON(200, "your order placed successfully!")
  }	

func GetItemFromCart() gin.HandlerFunc {
	return func(c *gin.Context){
		user_id := c.Query("id")
		if user_id == ""{
			c.Header("Content-Type", "application/json")
			c.JSON(http.StatusNotFound, gin.H("error":"Invalid id")
			c.Abort()
			return
		}

		usert_id, _ := primitive.ObjectIDFromHex(user_id)
		var ctx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()
		var filledcart models.User
		UserCollection.FindOne(ctx, bson.D{primitive.E{Key:"_id", Value: usert_id}}).Decode(&filledcart)
        if err != nil{
			log.println(err)
			c.IndentedJSON(500, "not found")
			return
		}

		filter_match := bson.D{{Key:"$match", Value: bson.D{primitive.E{Key: "_id", Value: usert_id}}}} 
		
		//after finding that user, the found user must be having cart items (now bind will help u to work with that data)
		unwind := bson.D{{Key: "$unwind", Value: bson.D{primitive.E{Key: "path", Value:"$usercart"}}}}
		
		//Addition of all the value of items in cart
		grouping := bson.D{{Key: "$group", Value: bson.D{primitive.E{Key: "_id", Value:"$_id"}, {Key:"total", Value:bson.D{primitive.E{Key: "$sum", Value:"$usercart.price"}} }}}}
		pointcursor, err := UserCollection.Aggregate(ctx, mongo.Pipeline{filter_match, unwind, grouping})
		if err != nil{
			log.Println(err)
		}
		var listing []bson.M
		if err = pointcursor.All(ctx, &listing); err != nil {
			log.Println(err)
			c.AbortWithStatus(http.StatusInternalServerError)
		}
		for _, json := range listing{
			c.IndentedJSON(200, json["total"])
			c.IndentedJSON(200, filledcart.UserCart)
		}
		ctx.Done()
	}
}

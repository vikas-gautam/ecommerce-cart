package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/vikas-gautam/ecommerce-cart/database"
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

func GetItemFromCart() gin.HandlerFunc {}

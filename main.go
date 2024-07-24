package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func increment(ctx *gin.Context) {

	// Get the value for the key
	val, err := rdb.Get(ctx, "count").Int64()
	if err != nil {
		fmt.Println("key is not found", err)
		// Set a key-value pair
		val = 1
		err = rdb.Set(ctx, "count", 1, 0).Err()
		if err != nil {
			fmt.Println("Error setting value:", err)
			return
		}
	} else {
		// Set a key-value pair
		err = rdb.Set(ctx, "count", val+1, 0).Err()
		if err != nil {
			fmt.Println("Error setting value:", err)
			return
		}
	}
	htmlContent := fmt.Sprintf(`
            <!DOCTYPE html>
            <html lang="en">
            <head>
                <meta charset="UTF-8">
                <meta name="viewport" content="width=device-width, initial-scale=1.0">
                <title>Simple Message</title>
                <style>
                    .box {
                        border: 2px solid #4CAF50;
                        padding: 20px;
                        margin: 20px auto;
                        text-align: center;
                        width: 300px;
                        box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1);
                        font-family: Arial, sans-serif;
                    }
                </style>
            </head>
            <body>
                <div class="box">
                    <h1>Visited %d times</h1>
                </div>
            </body>
            </html>
        `, val)

	ctx.Data(http.StatusOK, "text/html; charset=utf-8", []byte(htmlContent))

}

func main() {
	r := gin.Default()

	// Define a context
	ctx := context.Background()

	// Create a new Redis client
	rdb = redis.NewClient(&redis.Options{
		Addr:     "redis-server:8082", // Redis server address
		Password: "",                  // No password set
		DB:       0,                   // Use default DB
	})
	// Ping the Redis server to check connection
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		fmt.Println("Error connecting to Redis:", err)
		return
	}
	fmt.Println("Redis connected:", pong)

	r.GET("/visit", increment)
	// Initialize Router
	r.Run(":8080")

}

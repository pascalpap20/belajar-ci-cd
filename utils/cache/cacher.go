package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"log"
	"time"
)

func main() {
	// membuat koneksi ke redis
	client := redis.NewClient(&redis.Options{
		Network:            "",
		Addr:               "localhost:6379",
		Dialer:             nil,
		OnConnect:          nil,
		Username:           "",
		Password:           "",
		DB:                 0,
		MaxRetries:         0,
		MinRetryBackoff:    0,
		MaxRetryBackoff:    0,
		DialTimeout:        0,
		ReadTimeout:        0,
		WriteTimeout:       0,
		PoolFIFO:           false,
		PoolSize:           0,
		MinIdleConns:       0,
		MaxConnAge:         0,
		PoolTimeout:        0,
		IdleTimeout:        0,
		IdleCheckFrequency: 0,
		TLSConfig:          nil,
		Limiter:            nil,
	})

	defer func(client *redis.Client) {
		err := client.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(client)

	data := map[string]any{
		"ID":       1,
		"username": "johndoe",
		"email":    "johndoe@gmail.com",
	}

	value, _ := json.Marshal(data)

	// menyimpan data dalam redis
	err := client.Set(context.Background(), "data", value, time.Hour).Err()
	if err != nil {
		log.Fatal(err)
	}

	// mengambil data dari redis
	val, err := client.Get(context.Background(), "data").Result()
	if err != nil {
		log.Fatal(err)
	}
	x := map[string]any{}

	json.Unmarshal([]byte(val), &x)

	fmt.Println(fmt.Sprintf("ID : %.0f", x["ID"]))
	fmt.Println(fmt.Sprintf("username : %s", x["username"]))
	fmt.Println(fmt.Sprintf("email : %s", x["email"]))

	res, err := client.Del(context.Background(), "data").Result()
	if err != nil {
		log.Fatal(err)
	}
	if res == 1 {
		fmt.Println("data berhasil dihapus")
	}
}

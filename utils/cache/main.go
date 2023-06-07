package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"log"
	"time"
)

type Person struct {
	Id    uint
	Name  string
	Email string
}

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // alamat port redis server
		Password: "",               // password redis (jika ada)
		DB:       0,                // index database redis
	})

	//menutup Koneksi Redis Ketika program selesai
	defer func(client *redis.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)
	// menyimpan data dalam redis

	person := Person{
		Id:    1,
		Name:  "John Doe",
		Email: "johndoe@gmail.com",
	}

	encoded, err := json.Marshal(person)
	if err != nil {
		fmt.Println("error encoding JSON:", err)
	}

	err = json.Unmarshal([]byte(encoded), &person)
	if err != nil {
		fmt.Println("Error decoding JSON", err)
		return
	}
	err = client.Set(context.Background(), "key", string(encoded), time.Hour).
		Err()
	if err != nil {
		log.Fatal(err)
	}

	// mengambil data dari redis
	val, err := client.Get(context.Background(), "key").Result()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("data dari redis", val)

	err = client.Del(context.Background(), "key").Err()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(client.Exists(context.Background(), "key"))
}

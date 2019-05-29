package main

import (
	"sync"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	var list [1000000]int
	var setGroup sync.WaitGroup

	setGroup.Add(len(list))

	for _, i := range list {
		go func(idx string, client *redis.Client) {
			client.Set(idx, "wow", 0)

			setGroup.Done()
		}(string(i), client)
	}

	setGroup.Wait()

	client.BgSave()

	var getGroup sync.WaitGroup

	getGroup.Add(len(list))

	for _, i := range list {
		go func(idx string, client *redis.Client) {
			client.Get(idx)

			getGroup.Done()
		}(string(i), client)
	}

	getGroup.Wait()
}

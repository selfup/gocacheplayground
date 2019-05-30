package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-redis/redis"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	http.HandleFunc("/api/v1", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			keys := r.URL.Query()
			key := keys.Get("key")
			value, err := client.Get(key).Result()

			if err != nil {
				http.Error(w, "Not Found", 404)
			}

			response := make(map[string]string)
			response["key"] = key
			response["value"] = value

			outgoing, err := json.Marshal(response)
			if err != nil {
				http.Error(w, "failed to stringify JSON", 500)
				return
			}

			w.Write(outgoing)
		case "POST":
			decoder := json.NewDecoder(r.Body)

			var incoming struct {
				Key     string
				Value   string
				Expires int64
			}

			err := decoder.Decode(&incoming)
			if err != nil {
				http.Error(w, "failed to parse JSON", 500)
				return
			}

			clientErr := client.Set(incoming.Key, incoming.Value, 0).Err()
			if clientErr != nil {
				http.Error(w, "Failed to SET", 500)
			}

			w.Write([]byte("{}"))
		default:
			http.Error(w, "", 405)
		}
	})

	http.ListenAndServe(":8080", nil)
}

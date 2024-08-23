package keyValueStore

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// HandleSet handles the HTTP endpoint for setting nad updating key-value pairs.
func HandleSet(kv *KeyValueStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost && r.Method != http.MethodPut {
			w.WriteHeader(http.StatusMethodNotAllowed)
		} else {
			log.Println("HandleSet")
			log.Printf("BEFORE SET LIST %v", kv.data)
			var task Task

			var bodyBytes []byte
			var errr error
			if r.Body != nil {
				bodyBytes, errr = io.ReadAll(r.Body)
				if errr != nil {
					log.Printf("Body reading error: %v", errr)
					return
				}
				defer r.Body.Close()
			}

			if len(bodyBytes) > 0 {
				var prettyJSON bytes.Buffer
				if errr = json.Indent(&prettyJSON, bodyBytes, "", "\t"); errr != nil {
					log.Printf("JSON parse error: %v", errr)
					return
				}
				log.Println("SET", string(prettyJSON.String()))
			} else {
				log.Printf("Body: No Body Supplied\n")
			}

			err := json.Unmarshal(bodyBytes, &task)
			if err != nil {
				log.Printf("Error: %s\n", err)
				http.Error(w, "Invalid request body", http.StatusBadRequest)
			}

			if r.Method == http.MethodPost {
				kv.Set("", task)
				log.Printf("POST AFTER SET LIST %v", kv.data)
				w.WriteHeader(http.StatusOK)
			} else if r.Method == http.MethodPut {
				kv.Set(task.Id, task)
				log.Printf("PUT AFTER SET LIST %v", kv.data)
				w.WriteHeader(http.StatusOK)
			}
		}
	}
}

// HandleGet handles the HTTP endpoint for retrieving values by key.
func HandleGet(kv *KeyValueStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
		} else {
			log.Println("HandleGet")
			log.Printf("GET LIST %v", kv.data)
			w.Header().Set("Content-Type", "application/json")
			id := r.URL.Query().Get("id")
			if id == "" { //create
				task := Task{Id: "", Title: "", Description: "", Status: "false"}
				json.NewEncoder(w).Encode(task)
			} else {
				task, ok := kv.Get(id)
				log.Printf("GET TASK RESULT: %v, %s", ok, task)
				if !ok {
					http.Error(w, "id not found", http.StatusNotFound)
					return
				} else {
					json.NewEncoder(w).Encode(task)
				}
			}
		}
	}
}

// HandleGet handles the HTTP endpoint for retrieving all values.
func HandleGetAll(kv *KeyValueStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
		} else {
			log.Println("HandleGetAll")
			log.Printf("GETALL LIST %v", kv.data)

			jsonList, err := json.Marshal(kv.data)
			if err != nil {
				log.Printf("Error: %s\n", err)
				http.Error(w, "id not found", http.StatusNotFound)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(jsonList)
		}
	}
}

// HandleGet handles the HTTP endpoint for retrieving all values.
func HandleDelete(kv *KeyValueStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusMethodNotAllowed)
		} else {
			log.Println("HandleDelete")
			log.Printf("BEFORE DELETE %v", kv.data)
			id := r.URL.Query().Get("id")
			log.Printf("BEFORE DELETE ID: %s", id)
			exist := kv.Delete(id)
			if !exist {
				http.Error(w, "id not found", http.StatusNotFound)
				return
			}
			log.Printf("AFTER DELETE %v", kv.data)
		}
	}
}

func Run() {
	// Create a new instance of KeyValueStore
	kv := NewKeyValueStore()

	// Set up HTTP handlers
	http.HandleFunc("/set", HandleSet(kv))
	http.HandleFunc("/get", HandleGet(kv))
	http.HandleFunc("/getAll", HandleGetAll(kv))
	http.HandleFunc("/delete", HandleDelete(kv))

	// Start the HTTP server
	port := 9000
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Starting key-value store on http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Printf("Key-value store Error: %s\n", err)
	}
}

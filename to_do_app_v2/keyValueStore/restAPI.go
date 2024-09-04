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
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		log.Printf("RESTAPI SET: %s", r.Method)
		if r.Method != http.MethodPost && r.Method != http.MethodPut {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		var task Task

		var bodyBytes []byte
		var err error
		bodyBytes, err = io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error: error reading body\n%s\n", err)
			http.Error(w, "error reading body", http.StatusBadRequest)
			return
		}
		defer r.Body.Close()

		if len(bodyBytes) == 0 {
			log.Printf("Error: empty body\n%s\n", err)
			http.Error(w, "empty body", http.StatusBadRequest)
			return
		}

		var prettyJSON bytes.Buffer
		if err = json.Indent(&prettyJSON, bodyBytes, "", "\t"); err != nil {
			log.Printf("Error: error printing indented JSON body\n%s\n", err)
			http.Error(w, "error printing indented JSON body", http.StatusBadRequest)
			return
		}
		log.Println("SET", string(prettyJSON.String()))

		err = json.Unmarshal(bodyBytes, &task)
		if err != nil {
			log.Printf("Error: Invalid request body\n%s\n", err)
			http.Error(w, "invalid request body", http.StatusBadRequest)
			return
		}

		if r.Method == http.MethodPost {
			kv.Set("", task)
			log.Printf("POST AFTER SET LIST %v", kv.data)
			w.WriteHeader(http.StatusOK)
		} else {
			kv.Set(task.Id, task)
			log.Printf("PUT AFTER SET LIST %v", kv.data)
			w.WriteHeader(http.StatusOK)
		}
	}
}

// HandleGet handles the HTTP endpoint for retrieving values by key.
func HandleGet(kv *KeyValueStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
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
			}
			json.NewEncoder(w).Encode(task)
		}
	}
}

// HandleGet handles the HTTP endpoint for retrieving all values.
func HandleGetAll(kv *KeyValueStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		if r.Method != http.MethodGet {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		log.Printf("GETALL LIST %v", kv.data)

		jsonList, err := json.Marshal(kv.data)
		if err != nil {
			log.Printf("Error: id not found\n%s\n", err)
			http.Error(w, "id not found", http.StatusNotFound)
			return
		}
		w.Write(jsonList)
	}
}

// HandleGet handles the HTTP endpoint for retrieving all values.
func HandleDelete(kv *KeyValueStore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		log.Printf("BEFORE DELETE %v", kv.data)
		id := r.URL.Query().Get("id")
		log.Printf("BEFORE DELETE ID: %s", id)
		exist := kv.Delete(id)
		if !exist {
			log.Printf("Error: id not found [%s]", id)
			http.Error(w, "id not found", http.StatusNotFound)
			return
		}
		log.Printf("AFTER DELETE %v", kv.data)
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
		log.Printf("Error starting key-value-store REST API: %s\n", err)
	}
}

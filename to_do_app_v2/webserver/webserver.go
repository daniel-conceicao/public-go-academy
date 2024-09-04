package webserver

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"goAcademy/todoApp/keyValueStore"
	"io"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

type TodoAppHomepageData struct {
	Tasks map[string]keyValueStore.Task
}

type TaskPageData struct {
	PageTitle string
	Task      keyValueStore.Task
	Action    string
}

type ErrorPageData struct {
	ErrorCode string
	Error     error
}

var homePage, taskPage, errorPage *template.Template

func Run() {
	homePage = template.Must(template.ParseFiles("templates/index.html"))
	taskPage = template.Must(template.ParseFiles("templates/task.html"))
	errorPage = template.Must(template.ParseFiles("templates/error.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/getList", http.StatusSeeOther)
	})
	http.HandleFunc("/getList", HandleGetList())
	http.HandleFunc("/editTask", HandleEditTask())
	http.HandleFunc("/updateTask", HandleUpdateTask())
	http.HandleFunc("/addTask", HandleAddTask())
	http.HandleFunc("/createTask", HandleCreateTask())
	http.HandleFunc("/deleteTask", HandleDeleteTask())

	// Start the HTTP server
	port := 11000
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Starting webserver on http://localhost%s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Printf("webserver Error: %s\n", err)
	}
}

func HandleGetList() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("GET LIST ", r.Method)
		var tasks map[string]keyValueStore.Task

		if r.Method != http.MethodGet {
			data := ErrorPageData{
				ErrorCode: strconv.Itoa(http.StatusMethodNotAllowed),
				Error:     errors.New(http.StatusText(http.StatusMethodNotAllowed)),
			}
			w.WriteHeader(http.StatusMethodNotAllowed)
			errorPage.Execute(w, data)
			return
		}

		resp, err := http.Get("http://localhost:9000/getAll")
		if err != nil {
			data := ErrorPageData{
				ErrorCode: "999",
				Error:     err,
			}
			w.WriteHeader(http.StatusBadRequest)
			errorPage.Execute(w, data)
			return
		}
		if resp.StatusCode != http.StatusOK {
			data := ErrorPageData{
				ErrorCode: strconv.Itoa(resp.StatusCode),
				Error:     errors.New(http.StatusText(resp.StatusCode)),
			}
			w.WriteHeader(resp.StatusCode)
			errorPage.Execute(w, data)
			return
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			data := ErrorPageData{
				ErrorCode: "999",
				Error:     err,
			}
			w.WriteHeader(http.StatusBadRequest)
			errorPage.Execute(w, data)
			return
		}
		defer resp.Body.Close()
		if len(bodyBytes) == 0 {
			data := ErrorPageData{
				ErrorCode: "001",
				Error:     errors.New("empty response body"),
			}
			w.WriteHeader(http.StatusBadRequest)
			errorPage.Execute(w, data)
			return
		}

		var prettyJSON bytes.Buffer
		if err = json.Indent(&prettyJSON, bodyBytes, "", "\t"); err != nil {
			data := ErrorPageData{
				ErrorCode: "001",
				Error:     errors.New("json indent error"),
			}
			w.WriteHeader(http.StatusBadRequest)
			errorPage.Execute(w, data)
			return
		}
		log.Println("GET LIST REST API", string(prettyJSON.String()))

		err = json.Unmarshal(bodyBytes, &tasks)
		if err != nil {
			data := ErrorPageData{
				ErrorCode: "999",
				Error:     err,
			}
			w.WriteHeader(http.StatusBadRequest)
			errorPage.Execute(w, data)
			return
		}
		log.Printf("GET LIST: %v", tasks)
		data := TodoAppHomepageData{
			Tasks: tasks,
		}
		homePage.Execute(w, data)
	}
}

func HandleEditTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("EDIT TASK", r.Method)

		if r.Method != http.MethodGet {
			data := ErrorPageData{
				ErrorCode: strconv.Itoa(http.StatusMethodNotAllowed),
				Error:     errors.New(http.StatusText(http.StatusMethodNotAllowed)),
			}
			w.WriteHeader(http.StatusMethodNotAllowed)
			errorPage.Execute(w, data)
			return
		}

		taskId := r.URL.Query().Get("taskId")
		if taskId == "" {
			data := ErrorPageData{
				ErrorCode: "999",
				Error:     errors.New("invalid task id"),
			}
			w.WriteHeader(http.StatusBadRequest)
			errorPage.Execute(w, data)
			return
		}

		var task keyValueStore.Task
		resp, err := http.Get(fmt.Sprintf("http://localhost:9000/get?id=%s", taskId))
		if err != nil {
			data := ErrorPageData{
				ErrorCode: "001",
				Error:     errors.New("error getting task"),
			}
			errorPage.Execute(w, data)
			return
		}

		if resp.StatusCode != http.StatusOK {
			data := ErrorPageData{
				ErrorCode: strconv.Itoa(resp.StatusCode),
				Error:     errors.New(http.StatusText(resp.StatusCode)),
			}
			w.WriteHeader(resp.StatusCode)
			errorPage.Execute(w, data)
			return
		}

		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			data := ErrorPageData{
				ErrorCode: "999",
				Error:     err,
			}
			w.WriteHeader(http.StatusBadRequest)
			errorPage.Execute(w, data)
			return
		}
		defer resp.Body.Close()
		if len(bodyBytes) == 0 {
			data := ErrorPageData{
				ErrorCode: "001",
				Error:     errors.New("empty response body"),
			}
			w.WriteHeader(http.StatusBadRequest)
			errorPage.Execute(w, data)
			return
		}

		var prettyJSON bytes.Buffer
		if err = json.Indent(&prettyJSON, bodyBytes, "", "\t"); err != nil {
			data := ErrorPageData{
				ErrorCode: "001",
				Error:     errors.New("json indent error"),
			}
			w.WriteHeader(http.StatusBadRequest)
			errorPage.Execute(w, data)
			return
		}
		log.Println("EDIT REST API", string(prettyJSON.String()))

		err = json.Unmarshal(bodyBytes, &task)
		if err != nil {
			data := ErrorPageData{
				ErrorCode: "999",
				Error:     err,
			}
			w.WriteHeader(http.StatusBadRequest)
			errorPage.Execute(w, data)
			return
		}
		log.Println("TASK_EDIT:", task)
		data := TaskPageData{
			PageTitle: "Edit task",
			Task:      task,
			Action:    "update",
		}
		taskPage.Execute(w, data)
	}
}

func HandleUpdateTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("UPDATE", r.Method)
		if r.Method != http.MethodPut {
			data := ErrorPageData{
				ErrorCode: strconv.Itoa(http.StatusMethodNotAllowed),
				Error:     errors.New(http.StatusText(http.StatusMethodNotAllowed)),
			}
			w.WriteHeader(http.StatusMethodNotAllowed)
			errorPage.Execute(w, data)
			return
		}
		decoder := json.NewDecoder(r.Body)
		var task keyValueStore.Task
		err := decoder.Decode(&task)
		if err != nil {
			data := ErrorPageData{
				ErrorCode: "999",
				Error:     err,
			}
			w.WriteHeader(http.StatusBadRequest)
			errorPage.Execute(w, data)
			return
		}
		log.Printf("UPDATE TASK: %v", task)
		// create a new HTTP client
		client := &http.Client{}
		jsonReq, err := json.Marshal(task)
		if err != nil {
			data := ErrorPageData{
				ErrorCode: "999",
				Error:     err,
			}
			w.WriteHeader(http.StatusBadRequest)
			errorPage.Execute(w, data)
			return
		}

		req, err := http.NewRequest(http.MethodPut, "http://localhost:9000/set", bytes.NewBuffer(jsonReq))
		if err != nil {
			data := ErrorPageData{
				ErrorCode: "999",
				Error:     err,
			}
			w.WriteHeader(http.StatusBadRequest)
			errorPage.Execute(w, data)
			return
		}
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		// send the request
		resp, err := client.Do(req)
		if err != nil {
			data := ErrorPageData{
				ErrorCode: "999",
				Error:     err,
			}
			w.WriteHeader(http.StatusBadRequest)
			errorPage.Execute(w, data)
			return
		}
		if resp.StatusCode != http.StatusOK {
			data := ErrorPageData{
				ErrorCode: strconv.Itoa(resp.StatusCode),
				Error:     errors.New(http.StatusText(resp.StatusCode)),
			}
			w.WriteHeader(resp.StatusCode)
			errorPage.Execute(w, data)
			return
		}
	}
}

func HandleAddTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			data := ErrorPageData{
				ErrorCode: strconv.Itoa(http.StatusMethodNotAllowed),
				Error:     errors.New(http.StatusText(http.StatusMethodNotAllowed)),
			}
			w.WriteHeader(http.StatusMethodNotAllowed)
			errorPage.Execute(w, data)
			return
		}

		task := keyValueStore.Task{Id: "", Title: "", Description: "", Status: "false"}

		log.Println("TASK_ADD:", task)
		data := TaskPageData{
			PageTitle: "New task",
			Task:      task,
			Action:    "create",
		}
		taskPage.Execute(w, data)
	}
}

func HandleCreateTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("CREATE ", r.Method)
		if r.Method != http.MethodPost {
			data := ErrorPageData{
				ErrorCode: strconv.Itoa(http.StatusMethodNotAllowed),
				Error:     errors.New(http.StatusText(http.StatusMethodNotAllowed)),
			}
			w.WriteHeader(http.StatusMethodNotAllowed)
			errorPage.Execute(w, data)
			return
		}
		decoder := json.NewDecoder(r.Body)
		var task keyValueStore.Task
		err := decoder.Decode(&task)
		if err != nil {
			data := ErrorPageData{
				ErrorCode: "999",
				Error:     err,
			}
			w.WriteHeader(http.StatusBadRequest)
			errorPage.Execute(w, data)
			return
		}
		if task.Title == "" {
			data := ErrorPageData{
				ErrorCode: "001",
				Error:     errors.New("title field is mandatory"),
			}
			w.WriteHeader(http.StatusBadRequest)
			errorPage.Execute(w, data)
			return
		}
		// create a new HTTP client
		client := &http.Client{}
		jsonReq, err := json.Marshal(task)
		if err != nil {
			data := ErrorPageData{
				ErrorCode: "999",
				Error:     err,
			}
			w.WriteHeader(http.StatusBadRequest)
			errorPage.Execute(w, data)
			return
		}
		// create a new POST request
		req, err := http.NewRequest(http.MethodPost, "http://localhost:9000/set", bytes.NewBuffer(jsonReq))

		if err != nil {
			data := ErrorPageData{
				ErrorCode: "999",
				Error:     err,
			}
			w.WriteHeader(http.StatusBadRequest)
			errorPage.Execute(w, data)
			return
		}
		req.Header.Set("Content-Type", "application/json; charset=utf-8")
		// send the request
		resp, err := client.Do(req)
		if err != nil {
			data := ErrorPageData{
				ErrorCode: "001",
				Error:     err,
			}
			w.WriteHeader(http.StatusBadRequest)
			errorPage.Execute(w, data)
		}
		if resp.StatusCode != http.StatusOK {
			data := ErrorPageData{
				ErrorCode: strconv.Itoa(resp.StatusCode),
				Error:     errors.New(http.StatusText(resp.StatusCode)),
			}
			w.WriteHeader(resp.StatusCode)
			errorPage.Execute(w, data)
		}
	}
}

func HandleDeleteTask() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("DELETE ", r.Method)
		if r.Method != http.MethodDelete {
			data := ErrorPageData{
				ErrorCode: strconv.Itoa(http.StatusMethodNotAllowed),
				Error:     errors.New(http.StatusText(http.StatusMethodNotAllowed)),
			}
			w.WriteHeader(http.StatusMethodNotAllowed)
			errorPage.Execute(w, data)
			return
		}
		taskId := r.URL.Query().Get("taskId")
		if taskId == "" {
			data := ErrorPageData{
				ErrorCode: "999",
				Error:     errors.New("invalid task id"),
			}
			w.WriteHeader(http.StatusBadRequest)
			errorPage.Execute(w, data)
			return
		}
		client := &http.Client{}
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("http://localhost:9000/delete?id=%s", taskId), nil)
		if err != nil {
			data := ErrorPageData{
				ErrorCode: strconv.Itoa(http.StatusBadRequest),
				Error:     errors.New(http.StatusText(http.StatusBadRequest)),
			}
			w.WriteHeader(http.StatusBadRequest)
			errorPage.Execute(w, data)
			return
		}
		// send the request
		resp, err := client.Do(req)
		if err != nil {
			data := ErrorPageData{
				ErrorCode: "001",
				Error:     err,
			}
			w.WriteHeader(http.StatusBadRequest)
			errorPage.Execute(w, data)
			return
		}
		if resp.StatusCode != http.StatusOK {
			data := ErrorPageData{
				ErrorCode: strconv.Itoa(resp.StatusCode),
				Error:     errors.New(http.StatusText(resp.StatusCode)),
			}
			w.WriteHeader(http.StatusBadRequest)
			errorPage.Execute(w, data)
			return
		}
	}
}

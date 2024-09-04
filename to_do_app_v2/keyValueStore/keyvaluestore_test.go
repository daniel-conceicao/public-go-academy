package keyValueStore

import (
	"testing"
)

func TestNewKeyValueStore(t *testing.T) {
	keyValueStore := NewKeyValueStore()
	if len(keyValueStore.data) != 0 {
		t.Errorf("got %q want %q", len(keyValueStore.data), 0)
	}
}

func TestSetTableDriven(t *testing.T) {
	testKey := "testKey"
	emptyKey := ""
	testTitle := "testTitle"
	testDescription := "testDescription"
	testStatus := "false"
	var task Task
	var exist bool

	var scenarios = []struct {
		key string
	}{
		{testKey},
		{emptyKey},
	}

	for _, tt := range scenarios {
		t.Run(tt.key, func(t *testing.T) {
			keyValueStore := NewKeyValueStore()
			keyValueStore.Set(tt.key, Task{Id: tt.key, Title: testTitle, Description: testDescription, Status: testStatus})
			if len(keyValueStore.data) != 1 {
				t.Errorf("got %q want %q", len(keyValueStore.data), 0)
			}

			var key string

			if tt.key == emptyKey {
				for key = range keyValueStore.data {
					if key == "" {
						t.Errorf("got %q want %q", key, "<GUUID>")
					}
					task, exist = keyValueStore.data[key]
				}
				if task.Id == emptyKey {
					t.Errorf("got %q want %q", task.Id, "<GUUID>")
				}

			} else {
				task, exist = keyValueStore.data[tt.key]
				if task.Id != tt.key {
					t.Errorf("got %q want %q", task.Id, tt.key)
				}
			}

			if !exist {
				t.Errorf("got %t want %t", exist, true)
			}

			if task.Title != testTitle {
				t.Errorf("got %q want %q", task.Title, testTitle)
			}
			if task.Description != testDescription {
				t.Errorf("got %q want %q", task.Description, testDescription)
			}
			if task.Status != testStatus {
				t.Errorf("got %q want %q", task.Status, testStatus)
			}
		})
	}
}

func TestGetTableDriven(t *testing.T) {
	testKey := "testKey"
	emptyKey := ""
	testTitle := "testTitle"
	testDescription := "testDescription"
	testStatus := "false"
	notExistingKey := "qwerty"
	var task Task
	var exist bool

	var scenarios = []struct {
		key string
	}{
		{testKey},
		{emptyKey},
		{notExistingKey},
	}

	for _, tt := range scenarios {
		t.Run(tt.key, func(t *testing.T) {
			keyValueStore := NewKeyValueStore()
			keyValueStore.Set(testKey, Task{Id: testKey, Title: testTitle, Description: testDescription, Status: testStatus})
			task, exist = keyValueStore.Get(tt.key)
			switch tt.key {
			case testKey:
				if !exist {
					t.Errorf("got %t want %t", exist, true)
				}
				if task.Id != testKey {
					t.Errorf("got %q want %q", task.Id, testKey)
				}
				if task.Title != testTitle {
					t.Errorf("got %q want %q", task.Title, testTitle)
				}
				if task.Description != testDescription {
					t.Errorf("got %q want %q", task.Description, testDescription)
				}
				if task.Status != testStatus {
					t.Errorf("got %q want %q", task.Status, testStatus)
				}
			case emptyKey, notExistingKey:
				if exist {
					t.Errorf("got %t want %t", exist, false)
				}
			}

		})
	}
}

func TestDEleteTableDriven(t *testing.T) {
	testKey := "testKey"
	emptyKey := ""
	testTitle := "testTitle"
	testDescription := "testDescription"
	testStatus := "false"
	notExistingKey := "qwerty"

	var scenarios = []struct {
		key string
	}{
		{emptyKey},
		{testKey},
		{notExistingKey},
	}

	for _, tt := range scenarios {
		t.Run(tt.key, func(t *testing.T) {
			keyValueStore := NewKeyValueStore()
			if tt.key != emptyKey {
				keyValueStore.Set(testKey, Task{Id: testKey, Title: testTitle, Description: testDescription, Status: testStatus})
			}

			switch tt.key {
			case emptyKey:
				exist := keyValueStore.Delete(testKey)
				if exist {
					t.Errorf("got %t want %t", exist, false)
				}
			case testKey:
				exist := keyValueStore.Delete(testKey)
				if !exist {
					t.Errorf("got %t want %t", exist, true)
				}
			case notExistingKey:
				exist := keyValueStore.Delete(notExistingKey)
				if exist {
					t.Errorf("got %t want %t", exist, false)
				}
			}
		})
	}
}

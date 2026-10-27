package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateAndGet(t *testing.T) {
	app := &App{store: NewStore()}
	srv := httptest.NewServer(setupRoutes(app))
	defer srv.Close()

	resp, err := http.Post(srv.URL+"/tasks", "application/json",
		bytes.NewBufferString(`{"title":"buy milk"}`))
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 201 {
		t.Fatalf("create: got %d, want 201", resp.StatusCode)
	}

	resp, err = http.Get(srv.URL + "/tasks")
	if err != nil {
		t.Fatal(err)
	}
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	var tasks []Task
	json.Unmarshal(body, &tasks)
	if len(tasks) != 1 || tasks[0].Title != "buy milk" {
		t.Fatalf("list: got %+v", tasks)
	}

	resp, err = http.Get(srv.URL + "/tasks/1")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != 200 {
		t.Fatalf("get: got %d, want 200", resp.StatusCode)
	}
}

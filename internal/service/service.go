package service

import (
	"fmt"
	"net/http"
)

type MyService struct {
	Port int
}

func NewMyService(port int) *MyService {
	return &MyService{Port: port}
}

func (s *MyService) Start() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	addr := fmt.Sprintf(":%d", s.Port)
	http.ListenAndServe(addr, nil)
}

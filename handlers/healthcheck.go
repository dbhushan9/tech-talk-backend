package handlers

import (
	"fmt"
	"log"
	"net/http"
)

type Healthcheck struct {
	logger *log.Logger
}

func NewHealthcheck(l *log.Logger) *Healthcheck {
	return &Healthcheck{l}
}

func (h *Healthcheck) Get(res http.ResponseWriter, req *http.Request) {
	fmt.Fprint(res, "{\"message\":\"Healthcheck endpoint\"}")
}

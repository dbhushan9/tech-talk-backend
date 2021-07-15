package handlers

import (
	"dbhushan9/tech-talk-backend/models"
	"dbhushan9/tech-talk-backend/services"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
)

type TechTalk struct {
	logger  *log.Logger
	service *services.TechTalkService
}

func New(l *log.Logger, db *mongo.Database) *TechTalk {
	t := new(TechTalk)
	t.initialize(l, db)
	return t
}

func (t *TechTalk) initialize(l *log.Logger, db *mongo.Database) {
	t.logger = l
	t.service = services.NewTechTalkService(l, db)

}

func (t *TechTalk) Get(res http.ResponseWriter, req *http.Request) {
	data, _ := t.service.GetAll()
	msg := fmt.Sprintf("found %d tech talks", len(data))
	ResponseWriter(res, http.StatusOK, msg, data)
}

func (t *TechTalk) GetByID(res http.ResponseWriter, req *http.Request) {

}

func (t *TechTalk) Create(res http.ResponseWriter, req *http.Request) {
	techTalk := &models.TechTalk{}
	techTalk.FromJSON((req.Body))
	data, _ := t.service.Create(techTalk)
	ResponseWriter(res, http.StatusCreated, "created tech talk", data)
}

func (t *TechTalk) Update(res http.ResponseWriter, req *http.Request) {

	techTalk := &models.TechTalk{}
	techTalk.FromJSON((req.Body))
	vars := mux.Vars(req)
	id := vars["id"]
	techTalk.Id = id
	data, _ := t.service.Update(id, techTalk)
	ResponseWriter(res, http.StatusOK, "update tech talks", data)

}

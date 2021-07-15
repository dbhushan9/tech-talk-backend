package domain

import (
	"context"
	"dbhushan9/tech-talk-backend/models"
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type TechTalkDAO struct {
	logger     *log.Logger
	collection *mongo.Collection
}

const TechTalkCollectionName = "tech-talks"

func NewTechTalkDAO(logger *log.Logger, db *mongo.Database) *TechTalkDAO {
	return &TechTalkDAO{logger, db.Collection(TechTalkCollectionName)}
}

func (t *TechTalkDAO) Save(techTalk *models.TechTalk) (*models.TechTalk, error) {
	techTalk.Id = uuid.New().String()
	result, err := t.collection.InsertOne(context.TODO(), techTalk)
	if err != nil {
		return nil, err
	}
	techTalk.Id = result.InsertedID.(string)
	return techTalk, nil
}

func (t *TechTalkDAO) GetAll() ([]*models.TechTalk, error) {
	var techTalkList []*models.TechTalk
	findOptions := options.Find()

	cur, err := t.collection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Printf("Error while quering collection: %v\n", err)
		return nil, err
	}

	for cur.Next(context.TODO()) {
		// create a value into which the single document can be decoded
		var elem models.TechTalk
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		techTalkList = append(techTalkList, &elem)
	}
	return techTalkList, err
}

func (t *TechTalkDAO) Update(id string, updateData *models.TechTalk) (*models.TechTalk, error) {
	update := bson.M{
		"$set": updateData,
	}
	filter := bson.M{"_id": id}
	result, err := t.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	if result.MatchedCount == 1 {
		return updateData, nil
	}
	return nil, errors.New(fmt.Sprintf("no tech talk found with ID: %s", id))
}

func (t *TechTalkDAO) Delete() {

}

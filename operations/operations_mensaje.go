package operations

import (
	"../helper"
	"../models"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"context"
	"fmt"
)



func getClient2() *mongo.Collection {
	return helper.ConnectDB("prod", "mensajes")
}


func CreaMensaje(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	defer r.Body.Close()
	var msg models.TableroMsg
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		ResponseWriter(w, http.StatusBadRequest, "body json request have issues!!!", nil)
		return
	}
	msg.ID = primitive.NewObjectID()
	result, err := getClient2().InsertOne(nil, msg)
	if err != nil {
		switch err.(type) {
		case mongo.WriteException:
			ResponseWriter(w, http.StatusNotAcceptable, "username or email already exists in database.", nil)
		default:
			ResponseWriter(w, http.StatusInternalServerError, "Error while inserting data.", nil)
		}
		return
	}
	msg.ID = result.InsertedID.(primitive.ObjectID)
	ResponseWriter(w, http.StatusCreated, "", msg)
}


func GetMensaje(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	client := getClient2()
	id, _ := primitive.ObjectIDFromHex(string(mux.Vars(r)["id"]))
	filter := bson.M{"_id": id}
	var msg models.TableroMsg
	err := client.FindOne(context.TODO(), filter).Decode(&msg)
	if err != nil {
		fmt.Println(err)
	}
	ResponseWriter(w, http.StatusOK, "", msg)
}

func GetMensajes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client := getClient2()
	cursor, err := client.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal("pendejo " +
			"", err)
		panic(err)

	}

	defer cursor.Close(context.TODO())
	var msgs []*models.TableroMsg
	for cursor.Next(context.TODO()){
		var msg models.TableroMsg
		if err = cursor.Decode(&msg); err != nil{
			log.Fatal("la muerte")
		}
		msgs = append(msgs, &msg)
	}

	ResponseWriter(w, http.StatusOK, "", &msgs)

}

func ActualizaMensajes(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	id, _ := primitive.ObjectIDFromHex(string(mux.Vars(r)["id"]))
	client := getClient2()
	var updateData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updateData)

	update := bson.M{
		"$set":
		bson.M{
			"habilitado": 0,
			//"nombre": "dsadas",
		},
	}
	result, err := client.UpdateOne(context.Background(), bson.M{"_id": id}, update)
	if err != nil {
		log.Printf("Error while updateing document: %v", err)
		ResponseWriter(w, http.StatusInternalServerError, "error in updating document!!!", nil)
		return
	}
	if result.MatchedCount == 1 {
		ResponseWriter(w, http.StatusOK, "", &updateData)
	} else {
		ResponseWriter(w, http.StatusNotFound, "Message not found", nil)
	}
}
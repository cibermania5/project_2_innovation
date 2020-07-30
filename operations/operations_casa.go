package operations

import (
	"../helper"
	"../models"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)



func getClient() *mongo.Collection {
	return helper.ConnectDB("prod", "casas3")
}


func GetCasa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client := getClient()
	id, _ := primitive.ObjectIDFromHex(string(mux.Vars(r)["id"]))
	filter := bson.M{"_id": id}
	var casa models.Casa
	err := client.FindOne(context.TODO(), filter).Decode(&casa)
	if err != nil {
		fmt.Println(err)
		ResponseWriter(w, http.StatusInternalServerError, "error in adding document!!!", nil)
	}
	ResponseWriter(w, http.StatusOK, "", casa)
}

func CreaCasa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	defer r.Body.Close()
	var casa models.Casa
	err := json.NewDecoder(r.Body).Decode(&casa)
	if err != nil {
		ResponseWriter(w, http.StatusBadRequest, "body json request have issues!!!", nil)
		return
	}
	casa.ID = primitive.NewObjectID()
	result, err := getClient().InsertOne(nil, casa)
	if err != nil {
		switch err.(type) {
		case mongo.WriteException:
			ResponseWriter(w, http.StatusNotAcceptable, "username or email already exists in database.", nil)
		default:
			ResponseWriter(w, http.StatusInternalServerError, "Error while inserting data.", nil)
		}
		return
	}
	casa.ID = result.InsertedID.(primitive.ObjectID)
	ResponseWriter(w, http.StatusCreated, "", casa)
}

func AniadeMultas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var updateData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updateData)
	if err != nil {
		ResponseWriter(w, http.StatusBadRequest, "json body is incorrect", nil)
		return
	}
	// we dont handle the json decode return error because all our fields have the omitempty tag.
	var params = mux.Vars(r)
	oid, err := primitive.ObjectIDFromHex(params["id"])
	fmt.Println("oid ", oid)
	if err != nil {
		ResponseWriter(w, http.StatusBadRequest, "id that you sent is wrong!!!", nil)
		return
	}
	update := bson.M{
		//"$addToSet": updateData,
		//"$addToSet": bson.M{"cobro": updateData},
		"$push": bson.M{"cobros": updateData},
		"$set": bson.M{ "debe": 1 } ,
	}
	result, err := getClient().UpdateOne(context.Background(), bson.M{"_id": oid}, update)
	if err != nil {
		log.Printf("Error while updateing document: %v", err)
		ResponseWriter(w, http.StatusInternalServerError, "error in updating document!!!", nil)
		return
	}
	if result.MatchedCount == 1 {
		ResponseWriter(w, http.StatusOK, "", &updateData)
	} else {
		ResponseWriter(w, http.StatusNotFound, "Casa not found", nil)
	}
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	client := getClient()
	cursor, err := client.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal("pendejo " +
			"", err)
		//panic(err)
		ResponseWriter(w, http.StatusInternalServerError, "error in getting document!!!", nil)

	}

	defer cursor.Close(context.TODO())
	var todasLasCasas []*models.Casa
	for cursor.Next(context.TODO()){
		var casa models.Casa
		if err = cursor.Decode(&casa); err != nil{
			log.Fatal("la muerte")
			ResponseWriter(w, http.StatusInternalServerError, "error in getting document!!!", nil)
		}
		todasLasCasas = append(todasLasCasas, &casa)
	}
	//jsonRes, err := json.Marshal(&todasLasCasas)
	//fmt.Println(string(jsonRes))
	//json.NewEncoder(w).Encode(&todasLasCasas)
	ResponseWriter(w, http.StatusOK, "", &todasLasCasas)

}
func CalculaTotalCasa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id, _ := primitive.ObjectIDFromHex(string(mux.Vars(r)["id"]))

	pipeline := []bson.M{
		bson.M{ "$match":
		bson.M{ "_id":  id} },
		bson.M{"$unwind": "$cobros"},
		bson.M{ "$group": bson.M{ "_id": "null", "cantidad": bson.M{"$sum": "$cobros.monto"}}},
		bson.M{"$project": bson.M{"_id": 0} },

	}
	client := getClient()
	data, err :=  client.Aggregate(context.TODO(), pipeline)

	if err != nil {
		fmt.Println("data ", data)
		log.Println(err.Error())
		fmt.Errorf("failed to execute aggregation %s", err.Error())
		ResponseWriter(w, http.StatusInternalServerError, "error in updating document!!!", nil)
		return
	}

	var pipelineResult  []bson.M

	//fmt.Println("da ", data.ID())
	err = data.All(context.TODO(), &pipelineResult)

	if err != nil {
		log.Println(err.Error())
		fmt.Errorf("failed to decode results", err.Error())
		ResponseWriter(w, http.StatusInternalServerError, "error in updating document!!!", nil)
		return
	}

	/*value, err2 := pipelineResult[0]["cantidad"]
	if err 2!= nil {
		log.Println(err.Error())
		fmt.Errorf("failed to decode results", err.Error())
		return
	}*/

	//ResponseWriter(w, http.StatusOK, "", pipelineResult[0])

	if  len(pipelineResult) > 0  {
		ResponseWriter(w, http.StatusOK, "", pipelineResult[0])
	} else {
		ResponseWriter(w, http.StatusNotFound, "La persona no tiene adeudo", nil)
	}
}

func Pagar(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	id, _ := primitive.ObjectIDFromHex(string(mux.Vars(r)["id"]))
	client := getClient()
	var updateData map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&updateData)
	/*update := bson.M{
		//"$addToSet": updateData,
		//"$addToSet": bson.M{"cobro": updateData},
		//"$push": bson.M{"cobros": updateData},
		"$unset": bson.M{ "cobros": "" } ,
		"$set": bson.M{ "debe": false } ,
	}*/
	//val := false
	update := bson.M{
		"$set":
			bson.M{
			"debe": 0,
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
		ResponseWriter(w, http.StatusNotFound, "person not found", nil)
	}
}


package operations

import (
	"../helper"
	"../models"
	"context"
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"

	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type clientOptions *options.ClientOptions
var clOpt clientOptions

func getClient() *mongo.Collection {
	return helper.ConnectDB("testt", "tres")
}


func GetCasa(w http.ResponseWriter, r *http.Request) {
	client := getClient()
	id, _ := primitive.ObjectIDFromHex(string(mux.Vars(r)["id"])) //primitive.ObjectIDFromHex("5f1b610820d9c27b71722e42")
	//id, _ := primitive.ObjectIDFromHex(mux.Vars(r)) //primitive.ObjectIDFromHex("5f1b610820d9c27b71722e42")
	filter := bson.M{"_id": id}

	var rr models.Casa

	err := client.FindOne(context.TODO(), filter).Decode(&rr)

	if err != nil {
		fmt.Println(err)
	}


	json.NewEncoder(w).Encode(rr)

}

func CreaCasa(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("kdaspokjas``fkas`")
	now := time.Now()

	/*var casa models.Casa

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&casa)

	result, err := getClient().InsertOne(context.TODO(), casa)*/

	val := true
	test := models.Casa{
		ID: primitive.NewObjectID(),
		Casa: "casa03", Nombre: "Byron", Debe: val,
		Cobros: []models.Cobro{
			//{2000,"perros cagando", now},
			{800, "mantenimiento", now}},

	}

	result, err := getClient().InsertOne(context.TODO(), test)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func AniadeMultas(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	/*for k, v := range mux.Vars(r) {
		fmt.Printf("key[%s] value[%s]\n", k, v)

	}*/
	cobro := models.Cobro{}

	//fmt.Println(mux.Vars(r)["ms"])

	d := json.Unmarshal([]byte(mux.Vars(r)["ms"]), &cobro)

	fmt.Println("que ", d)
	/*now := time.Now()
	client := getClient()
	id, _ := primitive.ObjectIDFromHex(string(mux.Vars(r)["id"]))
	t := models.Cobro{400, "por cholo", now}


	var res, err = client.UpdateOne(context.TODO(),
		bson.M{"_id": id},
		bson.M{
			"$addToSet": bson.M{"cobro": t},
		})

	fmt.Println("count ", res.MatchedCount)
	if err != nil {
		log.Fatal("32132 ", err)
	}

	json.NewEncoder(w).Encode(res)*/
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	client := getClient()
	cursor, err := client.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal("pendejo " +
			"", err)
		panic(err)

	}

	defer cursor.Close(context.TODO())

	var todasLasCasas []*models.Casa
	for cursor.Next(context.TODO()){
		var casa models.Casa
		if err = cursor.Decode(&casa); err != nil{
			log.Fatal("la muerte")
		}
		todasLasCasas = append(todasLasCasas, &casa)
	}

	//jsonRes, err := json.Marshal(&todasLasCasas)
	//fmt.Println(string(jsonRes))


	json.NewEncoder(w).Encode(&todasLasCasas)

}
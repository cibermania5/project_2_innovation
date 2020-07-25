/*
	Author:  Mike Motta
	Purpose: This is the main file
*/

package main

import (
	"./operations"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


//type clientOptions *options.ClientOptions
//var clOpt clientOptions



func main() {
	r:= mux.NewRouter()

	r.HandleFunc("/casas/consulta/{id}", operations.GetCasa).Methods("GET")
	r.HandleFunc("/casas/nuevo", operations.CreaCasa).Methods("POST")
	r.HandleFunc("/casas/multas/{id}/{ms}", operations.AniadeMultas).Methods("PATCH")
	r.HandleFunc("/casas/consulta", operations.GetTodos).Methods("GET")


	//now := time.Now()


	log.Fatal(http.ListenAndServe(":8000", r))

	/*id, _ := primitive.ObjectIDFromHex("5f1b610820d9c27b71722e42")
	filter := bson.M{"_id": id}

	var rr models.Casa

	err := client.FindOne(context.TODO(), filter).Decode(&rr)

	if err != nil {
		fmt.Println(err)
	}


	fmt.Println("objeto: ", rr)*/
	//fmt.Println(res.InsertedID.(primitive.ObjectID).Hex())



/*
	var casa []bson.M
	if err = cursor.All(context.TODO(), &casa); err != nil {
		log.Fatal(err)
		panic(err)
	}
	fmt.Println(casa)*/
}

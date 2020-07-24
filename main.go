/*
	Author:  Mike Motta
	Purpose: This is the main file
*/

package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"./helper"
	"./models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
type clientOptions *options.ClientOptions
var clOpt clientOptions





func main() {
	client := helper.ConnectDB("testt", "abc")


	now := time.Now()
	t := models.Cobro{
		Monto:  323132131.0,
		Fecha:  now,
	}

	res, err := client.InsertOne(context.TODO(), bson.D{
		{"monto", t.Monto},
		{"createdAt", primitive.DateTime(timeToMillis(t.Fecha))},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res.InsertedID.(primitive.ObjectID).Hex())

}

func timeToMillis(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

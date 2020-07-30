package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Casa struct {
	 ID     		primitive.ObjectID `json:"_id" bson:"_id"`
	 Casa   		string             `json:"casa,omitempty" bson:"casa,omitempty"`
	 Nombre 		string             `json:"nombre,omitempty" bson:"nombre,omitempty"`
	 Debe   		int8               `json:"debe,omitempty" bson:"debe,omitempty"`
	 Cobros[] 		Cobro            	`json:"cobros,omitempty" bson:"cobros,omitempty"`
}

type Cobro struct {
	 Monto 			float64   			`json:"monto,omitempty" bson:"monto,omitempty"`
	 Causa 			string    			`json:"causa,omitempty" bson:"causa,omitempty"`
	 Fecha 			time.Time 			`json:"fecha,omitempty" bson:"fecha,omitempty"`
}

type TableroMsg struct{
	 ID      		primitive.ObjectID 	`json:"_id,omitempty" bson:"_id.omitempty"`
	 Mensaje 		string 				`json:"mensaje,omitempty" bson:"mensaje,omitempty"`
	 Habilitado    	int8               	`json:"habilitado,omitempty" bson:"habilitado,omitempty"`
}


package transaction

import (
	"fmt"
	"github.com/aadi756/cardflow/app/pkg/platform/db/mongo"
	"github.com/globalsign/mgo/bson"
)

func PinChangeRequest()  ([]map[string]interface{}, error){
	query := bson.M{}

	mongoSession, mongoCollection := mongo.Mongo.GetSessionCollection("ratings")
	defer mongoSession.Close()

	result := make([]map[string]interface{}, 0)
	err := mongoCollection.Find(query).All(&result)
	if err != nil {
		fmt.Println(err)
	}
	return result, nil
}





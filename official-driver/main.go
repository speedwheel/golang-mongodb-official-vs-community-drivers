package main

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/kataras/iris"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/primitive"
	"github.com/mongodb/mongo-go-driver/mongo"
)

type Car struct {
	ID               primitive.ObjectID `json:"_id" bson:"_id"`
	TestID           int                `json:"test_id" bson:"test_id"`
	VehicleID        int                `json:"vehicle_id" bson:"vehicle_id"`
	TestDate         time.Time          `json:"test_date" bson:"test_date"`
	TestClassID      int                `json:"test_class_id" bson:"test_class_id"`
	TestType         string             `json:"test_type" bson:"test_type"`
	TestResult       string             `json:"test_result" bson:"test_result"`
	TestMileage      string             `json:"test_mileage" bson:"test_mileage"`
	PostcodeArea     string             `json:"postcode_area" bson:"postcode_area"`
	Make             string             `json:"make" bson:"make"`
	Model            string             `json:"model" bson:"model"`
	Colour           string             `json:"colour" bson:"colour"`
	FuelType         string             `json:"fuel_type" bson:"fuel_type"`
	CylinderCapacity int                `json:"cylinder_capacity" bson:"cylinder_capacity"`
	FirstUseDate     time.Time          `json:"first_use_date" bson:"first_use_date"`
}

var (
	dns            = "mongodb://localhost:27017"
	databaseName   = "test"
	collectionName = "test_result_2017"
	// Collections.
	carCollection *mongo.Collection

	ErrNotFound = errors.New("not found")
)

func main() {
	client, err := mongo.Connect(context.Background(), dns)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	db := client.Database(databaseName)
	carCollection = db.Collection(collectionName)

	app := iris.New()
	storeAPI := app.Party("/api/store")
	{
		storeAPI.Get("/cars/{id}", Get)
	}
	app.Run(iris.Addr(":8080"), iris.WithOptimizations)
}

func Get(ctx iris.Context) {
	id := ctx.Params().Get("id")

	m, err := GetByID(nil, id)
	if err != nil {
		if err == ErrNotFound {
			ctx.NotFound()
		} else {
			ctx.JSON(err)
		}
		return
	}

	ctx.JSON(m)
}

func GetByID(ctx context.Context, id string) (Car, error) {
	var car Car
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return car, err
	}
	err = carCollection.FindOne(nil, bson.D{{Key: "_id", Value: objectID}}).Decode(&car)
	if err == mongo.ErrNoDocuments {
		return car, ErrNotFound
	}
	return car, err
}

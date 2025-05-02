package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/anuj123upadhyay/mongoapi/model"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://akumarupadhyay8445:szCZwEomjkU0mKAy@cluster0.24mn3xe.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"

const dbName = "netflix"

const colName = "watchlist"

// Most Imp
var collection *mongo.Collection

//connect with mongodb

func init() {
	//client options
	clientOptions := options.Client().ApplyURI(connectionString)

	//	connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connection Success")

	collection = client.Database(dbName).Collection(colName)

	//collection instances
	fmt.Println("Collection instance is ready")
}

// MongoDB Helpers - usually are in the separate file but this keep it hereonly
// insert one record
func insertOneMovie(movie model.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	fmt.Println("Inserted one movie record", inserted)
	if err != nil {
		log.Fatal(err)
	}
	//inserted.InsertedID is a bson.ObjectId type
	fmt.Println("Inserted Document with ID", inserted.InsertedID)
}

//update one record

func updateOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Matched and modified count", result.MatchedCount, result.ModifiedCount)

}

// delete one record
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie deleted with delete count", deleteCount.DeletedCount)

}

// delete all records from mongoDB
func deleteAllMovies() int64 {
	// filter := bson.D{{}}
	// collection.DeleteMany(context.Background(), filter)
	deleteCount, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("All movies deleted with delete count", deleteCount.DeletedCount)
	return deleteCount.DeletedCount
}

// get all movies from db
func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var movies []primitive.M
	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)

	}
	defer cur.Close(context.Background())
	return movies
}

//Actual Controllers - file

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)

}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	insertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)

}

func MarkAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	params := mux.Vars(r)
	updateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	deleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	deleteCount := deleteAllMovies()
	json.NewEncoder(w).Encode(deleteCount)
}

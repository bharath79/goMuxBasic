package services

import (
	"context"
	"fmt"
	"log"

	"github.com/bharath79/golang/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb+srv://bharath:nV4Ma5TWuSjC0Uy6@cluster0.x73qz8d.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init()  {
	clientOption := options.Client().ApplyURI(connectionString)

	client,err := mongo.Connect(context.TODO(),clientOption)

	if(err != nil){
		log.Fatal(err)
	}
	fmt.Println("Mongo connected successfully")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Mongo table/collection successful")
}

func InsertMovie(movie model.Netflix) string{
	savedMovie,err := collection.InsertOne(context.Background(),movie)

	if(err!=nil){
		log.Fatal(err)
	}

	fmt.Println("movied saved with ID: ",savedMovie.InsertedID)

	mongoId := savedMovie.InsertedID
	stringObjectID := mongoId.(primitive.ObjectID).Hex()
	return stringObjectID

}

func UpdateMovie(movieId string)  {
	id,_ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}
	update := bson.M{"$set":bson.M{"watched":true}}
	updated,err :=collection.UpdateOne(context.Background(),filter,update)

	if(err !=nil){
		log.Fatal(err)
	}

	fmt.Println("updated successfully ",updated)
}

func DeleteMovie(movieId string)  {
	id,_ :=primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id":id}

	deleted,err:= collection.DeleteOne(context.Background(),filter)
	
	if(err!=nil){
		log.Fatal(err)
	}

	fmt.Println("movie deleted successfully ",deleted)
}

func DeleteMany()  {
	allDeleted,err := collection.DeleteMany(context.Background(),bson.D{{}})

	if(err!=nil){
		log.Fatal(err)
	}

	fmt.Println("All movies deleted ",allDeleted)
}

func GetAllMovies() []primitive.M {
	cursor,err := collection.Find(context.Background(),bson.M{})
	checkNilError(err)

	var movies []primitive.M
	for cursor.Next(context.Background()){
		var movie primitive.M
		err:=cursor.Decode(&movie)
		checkNilError(err)
		movies = append(movies, movie)

	}

	defer cursor.Close(context.Background())
	return movies

}

func checkNilError(err error)  {
	if(err!=nil){
		log.Fatal(err)
	}
}
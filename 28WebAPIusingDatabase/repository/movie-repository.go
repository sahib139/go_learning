package repository

import (
	"context"
	"log"

	config "github.com/sahib139/go_Basic_API_using_mongo/config"
	model "github.com/sahib139/go_Basic_API_using_mongo/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var movieCollection = config.Db.Collection("movie")

func InsertOne(movie model.Movie) {
	_, err := movieCollection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
}

func DeleteById(movieId string) int64 {
	_id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": _id}

	result, err := movieCollection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	return result.DeletedCount
}

func UpdateById(movieId string, updateMovie model.Movie) int64 {
	_id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": _id}

	updateFilter := bson.M{
		"$set": updateMovie,
	}

	result, err := movieCollection.UpdateOne(context.Background(), filter, updateFilter)

	if err != nil {
		log.Fatal(err)
	}
	return result.ModifiedCount
}

func GetById(movieId string) model.Movie {
	var movie model.Movie

	_id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": _id}

	err = movieCollection.FindOne(context.Background(), filter).Decode(&movie)
	if err != nil {
		log.Fatal(err)
	}

	return movie
}

func GetAll() []model.Movie {
	var movies []model.Movie

	cursor, err := movieCollection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var movie model.Movie
		if err := cursor.Decode(&movie); err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	return movies
}

func MarkAsWatched(movieId string) int64 {
	_id, err := primitive.ObjectIDFromHex(movieId)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": _id}
	updateFilter := bson.M{"$set": bson.M{"isWatched": true}}

	result, err := movieCollection.UpdateOne(context.Background(), filter, updateFilter)

	if err != nil {
		log.Fatal(err)
	}

	return result.ModifiedCount
}

func DeleteAll() int64 {
	result, err := movieCollection.DeleteMany(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	return result.DeletedCount
}

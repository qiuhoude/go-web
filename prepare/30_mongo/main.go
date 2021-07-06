package main

import (
	"context"
	"fmt"
	"github.com/qiuhoude/go-web/prepare/30_mongo/m"
	"github.com/qiuhoude/go-web/prepare/30_mongo/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {
	// Replace the uri string with your MongoDB deployment's connection string.

	uri := fmt.Sprintf("mongodb://%s:%s@%s",
		"admin",
		"123456",
		"localhost:27017",
	)
	m.Connect(uri)
	defer m.Disconnect()

	//example1()
	//createDocuments()
	//readDocuments()
	//updateDocuments()
	//findStructData()
	//insertStructData()
	aggregatePipeline()
}

func aggregatePipeline() {

	database := m.GetClient().Database("blog")
	//podcastsCollection := database.Collection("podcasts")
	episodesCollection := database.Collection("episodes")
	ctx, _ := context.WithTimeout(context.TODO(), 3*time.Second)

	id, _ := primitive.ObjectIDFromHex("60e27ce22b8b40b9ab46f2f3")

	matchStage := bson.D{{"$match", bson.D{{"podcast", id}}}}
	groupStage := bson.D{{"$group", bson.D{{"_id", "$podcast"}, {"total", bson.D{{"$sum", "$duration"}}}}}}

	showInfoCursor, err := episodesCollection.Aggregate(ctx, mongo.Pipeline{matchStage, groupStage})
	if err != nil {
		panic(err)
	}
	var showsWithInfo []bson.M
	if err = showInfoCursor.All(ctx, &showsWithInfo); err != nil {
		panic(err)
	}
	fmt.Println(showsWithInfo)

	lookupStage := bson.D{{"$lookup", bson.D{{"from", "podcasts"}, {"localField", "podcast"}, {"foreignField", "_id"}, {"as", "podcast"}}}}
	//unwindStage := bson.D{{"$unwind", bson.D{{"path", "$podcast"}, {"preserveNullAndEmptyArrays", false}}}}

	showLoadedCursor, err := episodesCollection.Aggregate(ctx, mongo.Pipeline{lookupStage})
	if err != nil {
		panic(err)
	}
	var showsLoaded []bson.M
	if err = showLoadedCursor.All(ctx, &showsLoaded); err != nil {
		panic(err)
	}
	fmt.Println(showsLoaded)
}

func insertStructData() {
	database := m.GetClient().Database("blog")
	podcastsCollection := database.Collection("podcasts")
	//episodesCollection := database.Collection("episodes")
	ctx, _ := context.WithTimeout(context.TODO(), 3*time.Second)

	podcast := model.Podcast{
		Title:  "The Polyglot Developer",
		Author: "Nic Raboy",
		Tags:   []string{"development", "programming", "coding"},
	}
	insertResult, err := podcastsCollection.InsertOne(ctx, podcast)
	if err != nil {
		panic(err)
	}
	fmt.Println(insertResult.InsertedID)
}

func findStructData() {
	database := m.GetClient().Database("blog")
	//podcastsCollection := database.Collection("podcasts")
	episodesCollection := database.Collection("episodes")
	ctx, _ := context.WithTimeout(context.TODO(), 3*time.Second)

	var episodes []model.Episode
	cursor, err := episodesCollection.Find(ctx, bson.M{"duration": bson.D{{"$gt", 25}}})
	if err != nil {
		panic(err)
	}

	if err = cursor.All(ctx, &episodes); err != nil {
		panic(err)
	}
	fmt.Println(episodes)

}

func updateDocuments() {
	database := m.GetClient().Database("blog")
	podcastsCollection := database.Collection("podcasts")
	//episodesCollection := database.Collection("episodes")
	ctx, _ := context.WithTimeout(context.TODO(), 3*time.Second)

	//id, _ := primitive.ObjectIDFromHex("60e27ce22b8b40b9ab46f2f3")
	//result, err := podcastsCollection.UpdateOne(
	//	ctx,
	//	//bson.M{"_id": id},
	//	bson.M{"title": "The Polyglot Developer Podcast"},
	//	bson.D{
	//		{"$set", bson.D{{"author", "houde qiu_"}}},
	//	},
	//)

	result, err := podcastsCollection.ReplaceOne(
		ctx,
		bson.M{"author": "houde qiu_"},
		bson.M{
			"title":  "The Nic Raboy Show",
			"author": "Nicolas Raboy",
		},
	)

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Updated %v Documents!\n", result.ModifiedCount)
}

func readDocuments() {
	database := m.GetClient().Database("blog")
	podcastsCollection := database.Collection("podcasts")
	episodesCollection := database.Collection("episodes")
	ctx, _ := context.WithTimeout(context.TODO(), 3*time.Second)

	opts := options.Find()
	opts.SetSort(bson.D{{"duration", 1}})
	cursor, err := episodesCollection.Find(ctx, bson.M{}, opts)
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var episode bson.M
		if err = cursor.Decode(&episode); err != nil {
			log.Fatal(err)
		}
		log.Println(episode)
	}

	var podcast bson.M
	if err = podcastsCollection.FindOne(ctx, bson.M{}).Decode(&podcast); err != nil {
		log.Fatal(err)
	}
	log.Println(podcast)
}

func createDocuments() {
	database := m.GetClient().Database("blog")
	podcastsCollection := database.Collection("podcasts")
	episodesCollection := database.Collection("episodes")
	ctx, _ := context.WithTimeout(context.TODO(), 3*time.Second)

	podcastResult, err := podcastsCollection.InsertOne(ctx, bson.D{
		{Key: "title", Value: "The Polyglot Developer Podcast"},
		{Key: "author", Value: "Nic Raboy"},
		{Key: "tags", Value: bson.A{"development", "programming", "coding"}},
	})
	if err != nil {
		log.Fatal(err)
	}

	episodeResult, err := episodesCollection.InsertMany(ctx, []interface{}{
		bson.D{
			{"podcast", podcastResult.InsertedID},
			{"title", "GraphQL for API Development"},
			{"description", "Learn about GraphQL from the co-creator of GraphQL, Lee Byron."},
			{"duration", 25},
		},
		bson.D{
			{"podcast", podcastResult.InsertedID},
			{"title", "Progressive Web Application Development"},
			{"description", "Learn about PWA development with Tara Manicsic."},
			{"duration", 32},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Inserted %v documents into episode collection!\n", len(episodeResult.InsertedIDs))

	log.Printf("%v\n", podcastResult)
}

func example1() {
	type Post struct {
		Title string `bson:"title,omitempty"`
		Body  string `bson:"body,omitempty"`
	}

	collection := m.GetClient().Database("blog").Collection("posts")

	/*
	  Insert documents
	*/
	docs := []interface{}{
		bson.D{{"title", "World"}, {"body", "Hello World"}},
		bson.D{{"title", "Mars"}, {"body", "Hello Mars"}},
		bson.D{{"title", "Pluto"}, {"body", "Hello Pluto"}},
	}

	res, insertErr := collection.InsertMany(context.Background(), docs)
	if insertErr != nil {
		log.Fatal(insertErr)
	}
	fmt.Println(res)
	/*
		Iterate a cursor and print it
	*/
	cur, currErr := collection.Find(context.Background(), bson.D{})

	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(context.Background())

	var posts []Post
	if err := cur.All(context.Background(), &posts); err != nil {
		log.Fatal(err)
	}
	fmt.Println(posts)
}

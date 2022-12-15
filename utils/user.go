package utils

import (
	"errors"

	"github.com/jkarage/alinker/env"
	"github.com/jkarage/alinker/models/db"
	"github.com/jkarage/alinker/models/user"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Userservice is for handling user db relation queries
type Userservice struct{}

func (userservice Userservice) Create(user *(user.User)) (primitive.ObjectID, error) {
	client, ctx, cancel, _ := db.Connect()
	defer cancel()

	db, col := useDb()
	database := client.Database(db)
	usersCollection := database.Collection(col)
	result := usersCollection.FindOne(ctx, bson.M{"email": user.Email})
	nameResult := usersCollection.FindOne(ctx, bson.M{"username": user.Username})
	if result.Err() == nil {
		return primitive.NilObjectID, errors.New("account with email already exists")
	} else if nameResult.Err() == nil {
		return primitive.NilObjectID, errors.New("username already taken")
	} else {
		insertResult, err := usersCollection.InsertOne(ctx, user)
		return (insertResult.InsertedID).(primitive.ObjectID), err
	}
}

func (userservice Userservice) Get(id primitive.ObjectID) (bson.Raw, error) {
	client, ctx, cancel, _ := db.Connect()
	defer cancel()

	db, col := useDb()
	database := client.Database(db)
	usersCollection := database.Collection(col)
	userResult := usersCollection.FindOne(ctx, bson.M{
		"_id": id,
	})
	user, err := userResult.DecodeBytes()
	if err != nil {
		return nil, errors.New(err.Error())
	}
	return user, nil
}
func (userservice Userservice) Delete(id primitive.ObjectID) error {
	client, ctx, cancel, _ := db.Connect()
	defer cancel()

	db, col := useDb()
	database := client.Database(db)
	userCollection := database.Collection(col)
	result := userCollection.FindOneAndDelete(ctx, bson.M{
		"_id": id,
	})

	if result.Err() != nil {
		return errors.New("no Documents with such ID")
	} else {
		return nil
	}
}

// Find user from email
func (userservice Userservice) FindByEmail(email string) (*user.User, error) {
	client, ctx, cancel, _ := db.Connect()
	defer cancel()

	db, col := useDb()
	database := client.Database(db)
	usersCollection := database.Collection(col)
	user := new(user.User)
	result := usersCollection.FindOne(ctx, bson.M{
		"email": email,
	})

	err := result.Decode(user)
	if err != nil {
		return nil, errors.New(err.Error())
	}

	return user, nil
}

// Update the Number of Apps
func (userservice Userservice) UpdateApps(email string) error {
	client, ctx, cancel, _ := db.Connect()
	defer cancel()

	db, col := useDb()
	database := client.Database(db)
	usersCollection := database.Collection(col)

	user := new(user.User)
	result := usersCollection.FindOne(ctx, bson.M{
		"email": email,
	})
	err := result.Decode(&user)
	if err != nil {
		return errors.New(err.Error())
	}
	_, err = usersCollection.UpdateOne(ctx, bson.M{
		"email": email,
	}, bson.M{
		"$set": bson.M{"apps": user.Apps + 1},
	})
	if err != nil {
		return errors.New(err.Error())
	}

	return nil

}

func useDb() (string, string) {
	db, err := env.Env("DB_NAME", "alinker")
	if err != nil {
		// log the err
	}
	collection, err := env.Env("COLLECTION_NAME", "users")
	if err != nil {
		// log the error
	}
	return db, collection
}

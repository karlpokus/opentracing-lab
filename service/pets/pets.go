package main

import (
	"time"
	"context"
	"encoding/json"
	"io"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/bson"
)

type Pet struct {
	Name string
	Type string
	Born time.Time
}

func findOnePet(c *mongo.Collection, petName string) ([]byte, error) {
	var p Pet
	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	filter := bson.M{"name": petName}
	err := c.FindOne(ctx, filter).Decode(&p)
	if err != nil {
		return nil, err
	}

	b, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func findAllPets(c *mongo.Collection) ([]byte, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
 	cur, err := c.Find(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var pets []Pet
	for cur.Next(ctx) {
		var p Pet
		err := cur.Decode(&p)
		if err != nil {
			return nil, err
		}
		pets = append(pets, p)
	}
	if err := cur.Err(); err != nil {
	  return nil, err
	}
	b, err := json.Marshal(pets)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func addOnePet(c *mongo.Collection, rBody io.ReadCloser) error {
	var p Pet
	if err := json.NewDecoder(rBody).Decode(&p); err != nil {
		return err
	}

	ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
	_, err := c.InsertOne(ctx, p)
	if err != nil {
		return err
	}
	return nil
}

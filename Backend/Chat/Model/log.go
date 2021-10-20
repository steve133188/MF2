package Model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Log struct {
	id string `json:id`
	username string `json:username`
	userId string `json:user_id`
	description string `json:description`
	date primitive.DateTime `json::date`
}



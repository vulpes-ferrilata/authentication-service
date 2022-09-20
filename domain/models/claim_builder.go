package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ClaimBuilder struct {
	id     primitive.ObjectID
	userID primitive.ObjectID
}

func (c ClaimBuilder) SetID(id primitive.ObjectID) ClaimBuilder {
	c.id = id

	return c
}

func (c ClaimBuilder) SetUserID(userID primitive.ObjectID) ClaimBuilder {
	c.userID = userID

	return c
}

func (c ClaimBuilder) Create() *Claim {
	return &Claim{
		aggregateRoot: aggregateRoot{
			aggregate: aggregate{
				id: c.id,
			},
		},
		userID: c.userID,
	}
}

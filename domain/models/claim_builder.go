package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ClaimBuilder interface {
	SetID(id primitive.ObjectID) ClaimBuilder
	SetUserID(userID primitive.ObjectID) ClaimBuilder
	Create() *Claim
}

func NewClaimBuilder() ClaimBuilder {
	return &claimBuilder{}
}

type claimBuilder struct {
	id     primitive.ObjectID
	userID primitive.ObjectID
}

func (c *claimBuilder) SetID(id primitive.ObjectID) ClaimBuilder {
	c.id = id

	return c
}

func (c *claimBuilder) SetUserID(userID primitive.ObjectID) ClaimBuilder {
	c.userID = userID

	return c
}

func (c claimBuilder) Create() *Claim {
	return &Claim{
		aggregateRoot: aggregateRoot{
			aggregate: aggregate{
				id: c.id,
			},
		},
		userID: c.userID,
	}
}

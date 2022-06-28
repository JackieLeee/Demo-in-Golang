package objid

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"Demo-in-Golang/demo/docker_mongo/id"
)

/**
 * @Author  Flagship
 * @Date  2022/4/10 14:56
 * @Description
 */

// FromID id转obj id
func FromID(id fmt.Stringer) (primitive.ObjectID, error) {
	return primitive.ObjectIDFromHex(id.String())
}

// MustFromID id转obj id（失败panic）
func MustFromID(id fmt.Stringer) primitive.ObjectID {
	oid, err := FromID(id)
	if err != nil {
		panic(err)
	}
	return oid
}

// ToUserID obj id转用户id
func ToUserID(oid primitive.ObjectID) id.UserID {
	return id.UserID(oid.Hex())
}

package mongo

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"Demo-in-Golang/demo/docker_mongo/mongo/objid"
)

/**
 * @Author  Flagship
 * @Date  2022/4/10 0:01
 * @Description
 */

// 通用字段名
const (
	IDFieldName       = "_id"
	UpdateAtFieldName = "update_at"
)

// IDField 定义id字段
type IDField struct {
	ID primitive.ObjectID `bson:"_id"`
}

// UpdateAtField 定义更新时间字段
type UpdateAtField struct {
	UpdateAt int64 `bson:"update_at"`
}

// NewObjID 生成一个新的obj id的函数
var NewObjID = primitive.NewObjectID

// NewObjIDWithValue 指定值生成新obj id的函数
func NewObjIDWithValue(id fmt.Stringer) {
	NewObjID = func() primitive.ObjectID {
		return objid.MustFromID(id)
	}
}

// UpdateAt 返回更新时间的函数
var UpdateAt = func() int64 {
	return time.Now().UnixNano()
}

// Set 返回一个set更新的document
func Set(v interface{}) bson.M {
	return bson.M{
		"$set": v,
	}
}

// SetOnInsert 返回一个setOnInsert更新或插入的document
func SetOnInsert(v interface{}) bson.M {
	return bson.M{
		"$setOnInsert": v,
	}
}

package docker_mongo

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"Demo-in-Golang/demo/docker_mongo/id"
	mgutil "Demo-in-Golang/demo/docker_mongo/mongo"
	"Demo-in-Golang/demo/docker_mongo/mongo/objid"
)

/**
 * @Author  Flagship
 * @Date  2022/6/28 19:37
 * @Description
 */

const openIDField = "open_id"

// Mongo 定义mongo数据访问对象
type Mongo struct {
	col *mongo.Collection
}

// NewMongo 创建一个mongo集合user
func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		col: db.Collection("user"),
	}
}

// ResolveUserID 从userid解析
func (m *Mongo) ResolveUserID(c context.Context, openID string) (id.UserID, error) {
	insertID := mgutil.NewObjID()
	res := m.col.FindOneAndUpdate(c, bson.M{
		openIDField: openID,
	}, mgutil.SetOnInsert(bson.M{
		mgutil.IDFieldName: insertID,
		openIDField:        openID,
	}), options.FindOneAndUpdate().
		SetUpsert(true).
		SetReturnDocument(options.After))

	if err := res.Err(); err != nil {
		return "", fmt.Errorf("cannot findOneAndUpdate: %s", err.Error())
	}

	var row mgutil.IDField
	if err := res.Decode(&row); err != nil {
		return "", fmt.Errorf("cannot decode result: %s", err.Error())
	}

	return objid.ToUserID(row.ID), nil
}

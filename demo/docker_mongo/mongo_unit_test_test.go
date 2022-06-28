package docker_mongo

import (
	"context"
	"os"
	"testing"

	"go.mongodb.org/mongo-driver/bson"

	"Demo-in-Golang/demo/docker_mongo/id"
	mgutil "Demo-in-Golang/demo/docker_mongo/mongo"
	"Demo-in-Golang/demo/docker_mongo/mongo/objid"
	"Demo-in-Golang/demo/docker_mongo/mongo/testing"
)

/**
 * @Author  Flagship
 * @Date  2022/6/28 19:23
 * @Description
 */

func TestResolveAccountID(t *testing.T) {
	// 初始化连接
	c := context.Background()
	mc, err := mongotesting.NewClient(c)
	if err != nil {
		t.Fatalf("connot connect mongodb: %s", err.Error())
		return
	}
	m := NewMongo(mc.Database("test"))

	// 执行预处理，先插入两条数据
	_, err = m.col.InsertMany(c, []interface{}{
		bson.M{
			mgutil.IDFieldName: objid.MustFromID(id.UserID("6251a48da2ff5a521f785bd6")),
			openIDField:        "openid_1",
		},
		bson.M{
			mgutil.IDFieldName: objid.MustFromID(id.UserID("6251a48da2ff5a521f785bd7")),
			openIDField:        "openid_2",
		},
	})
	if err != nil {
		t.Fatalf("cannot insert initial values: %s", err.Error())
	}

	// 生成指定obj id
	mgutil.NewObjIDWithValue(id.UserID("6251a48da2ff5a521f785bd8"))

	// 测试用例
	cases := []struct {
		name   string
		openID string
		want   string
	}{
		{
			// 用例1，已存在的用户
			name:   "existing_user",
			openID: "openid_1",
			want:   "6251a48da2ff5a521f785bd6",
		},
		{
			// 用例2，另一个已存在的用户
			name:   "another_existing_user",
			openID: "openid_2",
			want:   "6251a48da2ff5a521f785bd7",
		},
		{
			// 用例3，新用户
			name:   "new_user",
			openID: "openid_3",
			want:   "6251a48da2ff5a521f785bd8",
		},
	}

	// 执行测试用例
	for _, cc := range cases {
		t.Run(cc.name, func(t *testing.T) {
			userID, err := m.ResolveUserID(
				context.Background(), cc.openID)
			if err != nil {
				t.Errorf("failed resolve userID for %q: %s", cc.openID, err.Error())
			} else {
				if userID.String() != cc.want {
					t.Errorf("resolve userID: want: %q, got: %q", cc.want, userID)
				}
			}
		})
	}
}

func TestMain(m *testing.M) {
	os.Exit(mongotesting.RunWithMongoInDocker(m))
}

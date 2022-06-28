package mongotesting

import (
	"context"
	"fmt"
	"testing"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/docker/go-connections/nat"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/**
 * @Author  Flagship
 * @Date  2022/4/10 1:47
 * @Description
 */
const (
	// 镜像版本
	image = "mongo:4.4"
	// 容器端口
	containerPort = "27017/tcp"
)

// 全局mongoURI
var mongoURI string

// 默认MongoURI
const defaultMongoURI = "mongodb://localhost:27017"

// RunWithMongoInDocker 在docker容器中运行一个mongo实例
func RunWithMongoInDocker(m *testing.M) int {
	// 创建docker客户端
	c, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		panic(err)
	}

	// 创建docker容器
	ctx := context.Background()
	resp, err := c.ContainerCreate(ctx, &container.Config{
		// 容器镜像
		Image: image,
		// 容器对外暴露端口
		ExposedPorts: nat.PortSet{
			containerPort: {},
		},
	}, &container.HostConfig{
		// 端口绑定
		PortBindings: nat.PortMap{
			containerPort: []nat.PortBinding{
				{
					// 容器绑定的地址
					HostIP: "127.0.0.1",
					// 容器绑定的端口
					HostPort: "0",
				},
			},
		},
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}
	// 执行结束删除容器
	containerID := resp.ID
	defer func() {
		if err := c.ContainerRemove(ctx, containerID, types.ContainerRemoveOptions{
			Force: true,
		}); err != nil {
			panic(err)
		}
	}()

	// 启动容器
	if err := c.ContainerStart(ctx, containerID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	// 返回容器信息
	inspRes, err := c.ContainerInspect(ctx, containerID)
	if err != nil {
		panic(err)
	}
	// 容器ip端口
	hostPort := inspRes.NetworkSettings.Ports[containerPort][0]
	// 容器的MongoURI
	mongoURI = fmt.Sprintf("mongodb://%s:%s", hostPort.HostIP, hostPort.HostPort)

	return m.Run()
}

// NewClient 创建一个mongo client
func NewClient(c context.Context) (*mongo.Client, error) {
	if mongoURI == "" {
		return nil, fmt.Errorf("mongo uri not set. Please run RunWithMongoInDocker in TestMain")
	}
	return mongo.Connect(c, options.Client().ApplyURI(mongoURI))
}

// NewDefaultClient 创建一个默认连接到127.0.0.1:27017的mongo client
func NewDefaultClient(c context.Context) (*mongo.Client, error) {
	return mongo.Connect(c, options.Client().ApplyURI(defaultMongoURI))
}

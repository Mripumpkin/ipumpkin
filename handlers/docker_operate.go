package handlers

import (
	"context"
	"time"

	"ipumpkin/config"
	"ipumpkin/models"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// 操作容器信息
func DockerOperate(cfg config.Provider, mongodb *mongo.Database, logger *logrus.Logger) error {
	collection := mongodb.Collection(cfg.GetString("mongodb.docker"))
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		return err
	}
	containers, err := cli.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		return err
	}
	for _, container := range containers {
		Newcontainer := new(models.DockerContainer)
		Newcontainer.Name = container.Names[0][1:]
		Newcontainer.CreateAt = container.Created
		Newcontainer.ID = container.ID
		Newcontainer.ImageID = container.ImageID
		Newcontainer.Port = int16(container.Ports[0].PrivatePort)
		name := Newcontainer.Name
		_ = collection.FindOneAndDelete(context.Background(), bson.M{"name": name})
		_, err := collection.InsertOne(context.Background(), Newcontainer)
		if err != nil {
			logger.Info(err)
			return err
		}
	}
	logger.Info("Success to get docker information")
	return nil
}

// 重启容器
func RestartContainer(containerID string) error {
	var durationMinute *time.Duration
	*durationMinute = 1 * time.Minute
	ctx := context.Background()
	cli, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	err := cli.ContainerRestart(ctx, containerID, durationMinute)
	if err != nil {
		return err
	}
	return nil
}

// 启动容器
func StartContainer(containerID string) error {
	ctx := context.Background()
	cli, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	err := cli.ContainerStart(ctx, containerID, types.ContainerStartOptions{})
	if err != nil {
		return err
	}
	return nil
}

// 停止容器
func StopContainer(containerID string) error {
	var durationMinute *time.Duration
	*durationMinute = 1 * time.Minute
	ctx := context.Background()
	cli, _ := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	err := cli.ContainerStop(ctx, containerID, durationMinute)
	if err != nil {
		return err
	}
	return nil
}

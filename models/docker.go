package models

type DockerContainer struct {
	Name     string
	ImageID  string
	ID       string
	Port     int16
	CreateAt int64
}

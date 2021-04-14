package services

import (
	"fmt"
	"shop/demo/repositories"
)

//业务逻辑
type MovieService interface {
	ShowMovieName() string
}

type MovieServiceManager struct {
	repo repositories.MovieRepository
}

func NewMovieServiceManager(repo repositories.MovieRepository) *MovieServiceManager {
	return &MovieServiceManager{repo: repo}
}

func (m *MovieServiceManager)ShowMovieName() string {
	fmt.Println("我们获取到的视频名称:"+m.repo.GetMovieName())
	return fmt.Sprintf("我们获取到的视频名称: %s",m.repo.GetMovieName())
}

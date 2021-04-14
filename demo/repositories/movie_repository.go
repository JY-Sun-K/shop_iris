package repositories

import "shop/demo/datamodels"


//数据库操作


type MovieRepository interface {
	GetMovieName() string
}

type MovieManager struct {

}

func NewMovieManager () MovieRepository {
	return &MovieManager{}
}

func (m *MovieManager)GetMovieName()string  {
	movie := datamodels.Movie{Name: "大话西游"}
	return movie.Name
}

package Cache

import "GoWithGin/entity"

type MovieCache interface {
	Set(key string, value *entity.Movie)
	Get(key string) *entity.Movie
}

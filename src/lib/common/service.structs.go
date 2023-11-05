package common

import "github.com/spf13/viper"

type Pageable struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

func PageableFrom(page int, size int) Pageable {
	if page < 1 {
		page = 1
	}
	if size < 1 {
		size = viper.GetInt("services.pagination.size.default")
	}
	if size > viper.GetInt("services.pagination.size.max") {
		size = viper.GetInt("services.pagination.size.default")
	}

	return Pageable{
		Page: page,
		Size: size,
	}
}

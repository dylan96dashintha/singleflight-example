package main

import (
	"context"
	"github.com/singleflight-example/database"
	"github.com/singleflight-example/distributed_cache"
	"github.com/singleflight-example/usecase"
	"time"
)

func main() {
	ctx := context.Background()
	// database initialization
	templateRepo := database.NewMockTemplateRepository()

	// redis initialization
	cache := distributed_cache.NewMockCache()

	// use case initialization
	templateUsecase := usecase.NewTemplate(templateRepo, cache)

	for i := 0; i < 100; i++ {
		go func(i int) {
			_ = templateUsecase.GetTemplateNameById(ctx, "key1")

		}(i)
	}
	time.Sleep(5 * time.Second)
}

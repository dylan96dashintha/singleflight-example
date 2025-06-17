package usecase

import (
	"context"
	"fmt"
	"github.com/singleflight-example/database"
	"github.com/singleflight-example/distributed_cache"
	"golang.org/x/sync/singleflight"
)

type templateObj struct {
	templateRepo      database.TemplateRepository
	cache             distributed_cache.Cache
	singleFlightGroup *singleflight.Group
}

type Template interface {
	GetTemplateNameById(ctx context.Context, key string) string
}

func NewTemplate(templateRepo database.TemplateRepository,
	cache distributed_cache.Cache) Template {
	return &templateObj{
		templateRepo:      templateRepo,
		cache:             cache,
		singleFlightGroup: &singleflight.Group{},
	}
}

func (t templateObj) GetTemplateNameById(ctx context.Context,
	key string) string {
	name, err := t.cache.Get(ctx, key)
	if err != nil {
		fmt.Println("error in getting template name in cache, err ", err)

		result, err, _ := t.singleFlightGroup.Do(key, func() (interface{}, error) {
			tmpl, err := t.templateRepo.GetTemplateByID(key)
			if err != nil {
				fmt.Println("error in getting template name in database: ", err)
				return tmpl, err
			}
			//err = t.cache.Set(ctx, key, tmpl.Name, 1)
			//if err != nil {
			//	return nil, err
			//}
			return tmpl.Name, nil
		})
		if err != nil {
			return ""
		}

		return (result).(string)

	}
	return name
}

package business

import (
	"log"
	"sync"
	"time"

	"github.com/pavelerokhin/go-and-scrape/business/modules"
	"github.com/pavelerokhin/go-and-scrape/models/configs"
	"github.com/pavelerokhin/go-and-scrape/models/entities"
	"github.com/pavelerokhin/go-and-scrape/storage"
)

type Business struct {
	scrapper *modules.Scrapper
	storage  storage.Storage
	logger   *log.Logger
}

func GetBusinessLogic(logger *log.Logger, storage storage.Storage) Business {
	return Business{scrapper: modules.NewScrapper(logger), storage: storage, logger: logger}
}

func (b *Business) ScheduleOnce(mediaConfig configs.MediaConfig) {
	b.logger.Printf("scheduling scrape and persist of %d media at %v", len(mediaConfig),
		time.Now())
	var wg sync.WaitGroup
	for _, medium := range mediaConfig {
		wg.Add(1)
		go b.scrapeAndPersist(medium.MediumConfig, &wg)
	}
	wg.Wait()
}

func (b *Business) ScheduleWithInterval(mediaConfig configs.MediaConfig,
	persistenceConfig *configs.PersistenceConfig) {
	b.logger.Printf("scheduling scrape and persist work every %d sec",
		persistenceConfig.Interval)
	for {
		go b.ScheduleOnce(mediaConfig)
		time.Sleep(persistenceConfig.Interval * time.Second)
	}
}

func (b *Business) scrapeAndPersist(mediaConfig configs.MediumConfig, wg *sync.WaitGroup) {
	defer wg.Done()

	// scrape
	medium := b.storage.GetMediumByURL(mediaConfig.URL)
	articlePreviewsWithArticle := b.scrapper.Scrap(&mediaConfig)

	// persist
	if len(articlePreviewsWithArticle) > 0 {
		articlePreviewsWithArticle = modules.NormalizeText(articlePreviewsWithArticle)

		if medium == nil || medium.URL == "" {
			b.storage.SaveMedium(&entities.Medium{
				Name:            mediaConfig.Name,
				URL:             mediaConfig.URL,
				ArticlePreviews: articlePreviewsWithArticle,
			})
		} else {
			for _, articlePreview := range articlePreviewsWithArticle {
				a := articlePreview
				a.MediumID = medium.ID
				b.storage.SaveArticle(&a)
			}
		}
	}
}

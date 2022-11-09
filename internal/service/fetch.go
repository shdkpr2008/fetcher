package service

import (
	"context"
	"errors"
	"fetcher/internal/argument"
	"fetcher/internal/config"
	"fetcher/internal/model"
	"fetcher/internal/repository"
	"fetcher/internal/utility"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

const runtimeErrorText = "runtime error"
const htmlExtension = ".html"

type FetchService struct {
	argument           argument.Argument
	config             config.Config
	networkRepository  *repository.NetworkRepository
	storageRepository  *repository.StorageRepository
	metadataRepository *repository.MetadataRepository
}

func NewFetchService(argument argument.Argument, config config.Config,
	networkRepository *repository.NetworkRepository,
	storageRepository *repository.StorageRepository,
	metadataRepository *repository.MetadataRepository) *FetchService {
	return &FetchService{
		argument:           argument,
		config:             config,
		networkRepository:  networkRepository,
		storageRepository:  storageRepository,
		metadataRepository: metadataRepository,
	}
}

func (fS *FetchService) fetchProcessStoreUrl(url string) {
	if !utility.IsValidURL(url) {
		log.Println(fmt.Errorf("invalid url %s", url))
		return
	}

	urlSource, err := fS.networkRepository.Source(url)
	if err != nil {
		log.Println(err.Error())
		return
	}

	hostname, err := utility.HostNameFromEndpoint(url)
	if err != nil {
		log.Println(err.Error())
		return
	}

	filename := hostname + htmlExtension
	err = fS.storageRepository.Write(filename, strings.NewReader(urlSource))
	if err != nil {
		log.Println(err.Error())
		return
	}

	numLinks, images, err := utility.NumLinksImagesFromDocument(strings.NewReader(urlSource))
	if err != nil {
		log.Println(err.Error())
		return
	}

	err = fS.metadataRepository.Store(&model.Metadata{Site: hostname, NumLinks: numLinks,
		Images: images, LastFetch: time.Now().String()})
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Printf("Fetched url %s\n", url)
}

func (fS *FetchService) urlsMetadata(urls []string) {
	var hostnames []string

	for _, url := range urls {
		hostname, err := utility.HostNameFromEndpoint(url)
		if err != nil {
			log.Println(err.Error())
			continue
		}

		hostnames = append(hostnames, hostname)
	}

	mds, err := fS.metadataRepository.Metadata(hostnames)
	if err != nil {
		log.Println(err.Error())
		return
	}

	for _, metadata := range mds {
		fmt.Printf("%+v\n", *metadata)
	}
}

func (fS *FetchService) fetchProcessStoreUrls(urls []string) {
	maxUrlAtATime := fS.config.MaxThread()

	var wg sync.WaitGroup
	var ch = make(chan string)

	wg.Add(maxUrlAtATime)
	for i := 0; i < maxUrlAtATime; i++ {
		go func() {
			for {
				url, exists := <-ch
				if !exists {
					wg.Done()
					return
				}

				fS.fetchProcessStoreUrl(url)
			}
		}()
	}

	for i := 0; i < len(urls); i++ {
		ch <- urls[i]
	}

	close(ch)
	wg.Wait()
}

func (fS *FetchService) Run(ctx context.Context) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = errors.New(runtimeErrorText)
		}
	}()

	urls := fS.argument.Urls()
	if len(urls) == 0 {
		return nil
	}

	if fS.argument.IsShowMetadata() {
		fS.urlsMetadata(urls)
		return nil
	}

	fS.fetchProcessStoreUrls(urls)
	return nil
}

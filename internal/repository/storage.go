package repository

import (
	"fetcher/internal/config"
	"io"
	"os"
	"path/filepath"
)

type StorageRepository struct {
	config config.Config
}

func NewStorageRepository(config config.Config) *StorageRepository {
	return &StorageRepository{config: config}
}

func (s *StorageRepository) Write(filename string, content io.Reader) (err error) {
	_filepath := filepath.Join(s.config.WorkingDirectory(), filename)
	f, err := os.Create(_filepath)

	if err != nil {
		return err
	}
	defer f.Close()

	_, err = io.Copy(f, content)
	if err != nil {
		return err
	}

	return nil
}

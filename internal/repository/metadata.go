package repository

import (
	"fetcher/internal/model"
	"fetcher/internal/sqlite"
	"fmt"
	"strings"
)

type MetadataRepository struct {
	sqlite *sqlite.SQLite
}

func NewMetadataRepository(sqlite *sqlite.SQLite) *MetadataRepository {
	return &MetadataRepository{sqlite: sqlite}
}

func (m *MetadataRepository) Metadata(sites []string) ([]*model.Metadata, error) {
	var mds []*model.Metadata

	query := fmt.Sprintf("SELECT * FROM metadata where site in (%s)",
		"'"+strings.Join(sites, "', '")+"'")

	row, err := m.sqlite.Query(query)
	if err != nil {
		return mds, err
	}

	defer row.Close()
	for row.Next() {
		var site string
		var numLinks int
		var images int
		var lastFetch string

		err := row.Scan(&site, &numLinks, &images, &lastFetch)
		if err != nil {
			return mds, err
		}

		mds = append(mds, &model.Metadata{Site: site, NumLinks: numLinks, Images: images, LastFetch: lastFetch})
	}

	return mds, err
}

func (m *MetadataRepository) Store(metadata *model.Metadata) error {
	query := fmt.Sprintf(`INSERT INTO metadata (site, num_links, images, last_fetch)
						 VALUES ("%s",%v,%v,"%s") ON CONFLICT(site)
						 DO UPDATE SET num_links = "%v", images = "%v", last_fetch = "%s";`,
		metadata.Site, metadata.NumLinks, metadata.Images, metadata.LastFetch,
		metadata.NumLinks, metadata.Images, metadata.LastFetch)

	if _, err := m.sqlite.Prepare(query); err != nil {
		return err
	}

	if _, err := m.sqlite.Exec(query); err != nil {
		return err
	}

	return nil
}

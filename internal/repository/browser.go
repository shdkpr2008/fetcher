package repository

import (
	"context"
	"fetcher/internal/config"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
)

type BrowserRepository struct {
	config config.Config
}

func NewBrowserRepository(config config.Config) *BrowserRepository {
	return &BrowserRepository{config: config}
}

func (b *BrowserRepository) Source(endpoint string) (source string, err error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	ctx, cancel = context.WithTimeout(ctx, b.config.RequestTimeout())
	defer cancel()

	err = chromedp.Run(ctx,
		chromedp.Navigate(endpoint),
		chromedp.ActionFunc(func(ctx context.Context) error {
			node, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			source, err = dom.GetOuterHTML().WithNodeID(node.NodeID).Do(ctx)
			return err
		}),
	)

	return source, err
}

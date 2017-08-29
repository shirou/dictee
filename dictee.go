package dictee

import (
	"context"

	"github.com/apex/log"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"

	"github.com/shirou/dictee/dictionary"
)

type Dictee struct {
	ConfigRoot   string                  `json:"-"`
	Dictionaries []dictionary.Dictionary `json:"dictionaries"`
}

func (d *Dictee) SearchDictonary() {

}

func (d *Dictee) Search(ctx context.Context, word string, ch chan dictionary.Title, targets []string) error {
	defer close(ch)

	eg, ctx := errgroup.WithContext(ctx)

	for _, dict := range d.Dictionaries {
		if !isTarget(dict.GetName(), targets) {
			continue
		}

		dict := dict
		eg.Go(func() error {
			t, err := dict.Search(ctx, word)
			if err != nil {
				return errors.Wrapf(err, "search word:%s, dict: %s", word, dict.GetName())
			}
			for _, tt := range t {
				ch <- tt
			}
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return errors.Wrapf(err, "search wait")
	}

	return nil
}

func (d *Dictee) SearchAll(ctx context.Context, word string, ch chan dictionary.Title) error {
	return d.Search(ctx, word, ch, []string{})
}

func (d *Dictee) MakeIndex() {
	for _, dict := range d.Dictionaries {
		log.Infof("create index start: %v", dict)
		dict.MakeIndex(d.ConfigRoot)
	}
}

func isTarget(name string, targets []string) bool {
	if len(targets) == 0 {
		return true
	}
	for _, target := range targets {
		if name == target {
			return true
		}
	}
	return false
}

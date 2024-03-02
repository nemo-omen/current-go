package model

import (
	"encoding/json"
	"time"

	"github.com/mmcdole/gofeed"
)

type Newsfeed struct {
	ID            int64      `db:"id"`
	Title         string     `json:"title,omitempty" db:"title"`
	SiteUrl       string     `json:"link,omitempty" db:"site_url"`
	FeedUrl       string     `json:"feedLink,omitempty" db:"feed_url"`
	Description   string     `json:"description,omitempty" db:"description"`
	Image         *Image     `json:"image" db:"-"`
	Updated       string     `json:"updated,omitempty" db:"updated"`
	UpdatedParsed *time.Time `json:"updatedParsed,omitempty" db:"updated_parsed"`
	Copyright     string     `json:"copyright,omitempty" db:"copyright"`
	Articles      []*Article `json:"articles" db:"-"`
	Author        *Person    `json:"author" db:"-"`
	Language      string     `json:"language,omitempty" db:"language"`
	Categories    []string   `json:"categories,omitempty" db:"categories"`
	FeedType      string     `json:"feedType,omitempty" db:"feed_type"`
	Slug          string     `json:"slug" db:"slug"`
}

func (f Newsfeed) String() string {
	json, _ := json.MarshalIndent(f, "", "    ")
	return string(json)
}

func (f Newsfeed) Len() int {
	return len(f.Articles)
}

func (f Newsfeed) Less(i, k int) bool {
	return f.Articles[i].PublishedParsed.Before(
		f.Articles[k].PublishedParsed,
	)
}

func (f Newsfeed) Swap(i, k int) {
	f.Articles[i], f.Articles[k] = f.Articles[k], f.Articles[i]
}

func FeedFromRemote(rf *gofeed.Feed) (*Newsfeed, error) {
	m, err := json.Marshal(rf)
	nf := new(Newsfeed)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(m, nf)
	if err != nil {
		return nil, err
	}

	return nf, nil
}

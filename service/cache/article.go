package cache

import (
	"strconv"
	"strings"

	"github.com/yrjkqq/tiny-website/pkg/e"
)

// Article cache service article struct
type Article struct {
	ID    string
	TagID string
	State int

	PageNum  int
	PageSize int
}

// GetArticleKey get article cache key
func (a *Article) GetArticleKey() string {
	return e.CacheArticle + "_" + a.ID
}

// GetArticlesKey get article list cache key
func (a *Article) GetArticlesKey() string {
	keys := []string{
		e.CacheArticle,
		"LIST",
	}

	if a.ID != "" {
		keys = append(keys, a.ID)
	}
	if a.TagID != "" {
		keys = append(keys, a.TagID)
	}
	if a.State >= 0 {
		keys = append(keys, strconv.Itoa(a.State))
	}
	if a.PageNum > 0 {
		keys = append(keys, strconv.Itoa(a.PageNum))
	}
	if a.PageSize > 0 {
		keys = append(keys, strconv.Itoa(a.PageSize))
	}

	return strings.Join(keys, "_")
}

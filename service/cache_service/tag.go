package cache_service

import (
	"strconv"
	"strings"

	"github.com/yrjkqq/tiny-website/pkg/e"
)

// Tag cache service tag struct
type Tag struct {
	ID    string
	Name  string
	State int

	PageNum  int
	PageSize int
}

// GetTagsKey get tag list cache key
func (t *Tag) GetTagsKey() string {
	keys := []string{
		e.CacheTag,
		"LIST",
	}

	if t.Name != "" {
		keys = append(keys, t.Name)
	}
	if t.State >= 0 {
		keys = append(keys, strconv.Itoa(t.State))
	}
	if t.PageNum > 0 {
		keys = append(keys, strconv.Itoa(t.PageNum))
	}
	if t.PageSize > 0 {
		keys = append(keys, strconv.Itoa(t.PageSize))
	}

	return strings.Join(keys, "_")
}

package app

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"github.com/yrjkqq/tiny-website/models"
)

// ServeBackground 后台运行的任务
func ServeBackground() {
	c := cron.New(cron.WithSeconds())
	c.AddFunc("0 0 0 1 1 *", func() {
		fmt.Println("Start clear all soft deleted tags and articles")
		models.ClearAllSoftDeletedArticle()
		models.ClearAllSoftDeletedTag()
	})
	c.Start()
}

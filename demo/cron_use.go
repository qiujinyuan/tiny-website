package demo

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

// CronUseDemo ...
func CronUseDemo() {
	c := cron.New(cron.WithSeconds())
	// c := cron.New()
	// c.AddFunc("30-40/2 * * * * *", func() {
	// 	fmt.Println(time.Now().String())
	// })
	// c.AddFunc("50/2 * * * * *", func() {
	// 	fmt.Println(time.Now().String())
	// })
	// c.AddFunc("1,3,5,7,47 * * * * *", func() {
	// 	fmt.Println(time.Now().String())
	// })
	c.AddFunc("@every 1m3s", func() {
		fmt.Println(time.Now().String())
	})
	// c.AddFunc("30 3-6,20-23 * * *", func() { fmt.Println(".. in the range 3-6am, 8-11pm") })
	// c.AddFunc("CRON_TZ=Asia/Tokyo 30 04 * * *", func() { fmt.Println("Runs at 04:30 Tokyo time every day") })
	// c.AddFunc("@hourly", func() { fmt.Println("Every hour, starting an hour from now") })
	// c.AddFunc("@every 1h30m", func() { fmt.Println("Every hour thirty, starting an hour thirty from now") })
	c.Start()
	// Funcs are invoked in their own goroutine, asynchronously.
	// Funcs may also be added to a running Cron
	// c.AddFunc("@daily", func() { fmt.Println("Every day") })
	// Inspect the cron job entries' next and previous run times.
	// inspect(c.Entries())
	time.Sleep(30 * time.Minute)
	// c.Stop() // Stop the scheduler (does not stop any jobs already running).
}

func inspect(entries []cron.Entry) {
	for v, vv := range entries {
		fmt.Println(v, vv)
	}
}

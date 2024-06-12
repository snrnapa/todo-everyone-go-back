package main

import (
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/go-co-op/gocron/v2"
	"github.com/joho/godotenv"
)

func task() {
	fmt.Println("task!!!!")
}

func main() {
	err := godotenv.Load(filepath.Join("path", "to", "dotenv", ".env"))

	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
		return
	}

	ns, err := gocron.NewScheduler(gocron.WithLocation(jst))
	if err != nil {
		log.Fatal(err)
		return
	}

	ns.Start()

	nj, err := ns.NewJob(
		// gocron.CronJob("0 0 9 * * 1-5", true),
		gocron.CronJob("0 * * * * *", true),
		gocron.NewTask(task),
	)
	if err != nil {
		log.Fatal(err)
		return
	}

	// ぶっちゃけここのprintは特に意味はない
	fmt.Printf("job ID: %v\n", nj.ID)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/Dima-Vasilenko/downloader/internal/downloader"
)

func GetInt(message string) (res int) {
	fmt.Println(message)
	fmt.Scan(&res)
	return
}

func main() {
	var path string
	fmt.Println("Enter file path with urls to download:")
	fmt.Scan(&path)

	workersAmount := GetInt("Enter amount of workers who will be download files:")
	timeout := GetInt("Enter timeout:")

	dwL, err := downloader.NewDownloader(path, workersAmount, time.Duration(timeout))
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout*workersAmount)*time.Second)
	defer cancel()

	res, err := dwL.Run(ctx)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, val := range res {
		fmt.Printf("Status code: %v, time spend %v, duration %v\n", val.StatusCode, val.Written, val.Duration)
	}
}

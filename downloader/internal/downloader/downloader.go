package downloader

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

var (
	ErrIncorrectStatusCode = errors.New("incorrect status code")
	ErrIncorrectData       = errors.New("invalid input data")
)

func GetInt(message string) (res int) {
	fmt.Println(message)
	fmt.Scan(&res)
	return
}

type Result struct {
	StatusCode int
	Written    int64
	Duration   time.Duration
	Err        error
}

type Downloader struct {
	filePath      string
	workersAmount int
	timeout       time.Duration
}

func NewDownloader(path string, workers int, t time.Duration) (*Downloader, error) {
	if workers <= 0 || t <= 0 {
		return nil, ErrIncorrectData
	}
	return &Downloader{
		filePath:      path,
		workersAmount: workers,
		timeout:       t,
	}, nil
}

func download(ctx context.Context, client *http.Client, url string) Result {
	start := time.Now()
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return Result{StatusCode: -1, Written: -1, Duration: 0, Err: err}
	}
	resp, err := client.Do(req)
	if err != nil {
		return Result{StatusCode: -1, Written: -1, Duration: 0, Err: err}
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return Result{StatusCode: resp.StatusCode, Written: -1, Duration: 0, Err: ErrIncorrectStatusCode}
	}

	written, err := io.Copy(io.Discard, resp.Body)
	if err != nil {
		return Result{StatusCode: -1, Written: -1, Duration: 0, Err: err}
	}

	dur := time.Since(start)
	return Result{StatusCode: resp.StatusCode, Written: written, Duration: dur, Err: nil}
}

func (d *Downloader) Run(pctx context.Context) ([]Result, error) {

	var results []Result
	var mu sync.Mutex

	g, ctx := errgroup.WithContext(pctx)
	g.SetLimit(d.workersAmount)

	client := &http.Client{
		Timeout: d.timeout * time.Second,
	}

	file, err := os.Open(d.filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		url := scanner.Text()
		g.Go(func() error {
			res := download(ctx, client, url)

			mu.Lock()
			results = append(results, res)
			mu.Unlock()
			return nil
		})
	}

	if err := scanner.Err(); err != nil {
		return results, err
	}

	if err := g.Wait(); err != nil {
		return results, err
	}

	return results, nil
}

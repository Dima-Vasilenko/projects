package downloader_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/Dima-Vasilenko/downloader/internal/downloader"
	"github.com/stretchr/testify/assert"
)

func writeURLs(t *testing.T, urls ...string) string {
	t.Helper()
	path := filepath.Join(t.TempDir(), "urls.txt")
	if err := os.WriteFile(path, []byte(strings.Join(urls, "\n")), 0o644); err != nil {
		t.Fatal(err)
	}
	return path
}

func TestSucces(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello")
	}))
	defer srv.Close()

	path := writeURLs(t, srv.URL, srv.URL)
	d, _ := downloader.NewDownloader(path, 2, 5)
	results, err := d.Run(context.Background())

	assert.NoError(t, err)
	assert.Len(t, results, 2)
	assert.Equal(t, 200, results[0].StatusCode)
	assert.Equal(t, int64(5), results[0].Written)
	assert.Nil(t, results[0].Err)
}

func TestErrorStatusCode(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "hello")
	}))
	defer srv.Close()

	path := writeURLs(t, srv.URL, srv.URL)
	d, _ := downloader.NewDownloader(path, 2, 5)
	results, _ := d.Run(context.Background())

	assert.ErrorIs(t, results[0].Err, downloader.ErrIncorrectStatusCode)
}

func TestErrorTimeout(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
	}))
	defer srv.Close()

	path := writeURLs(t, srv.URL, srv.URL)
	d, _ := downloader.NewDownloader(path, 1, 50*time.Millisecond)

	results, _ := d.Run(context.Background())

	errors.Is(results[0].Err, context.DeadlineExceeded)
}

func TestErrorCancelContext(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
	}))
	defer srv.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	path := writeURLs(t, srv.URL, srv.URL)
	d, _ := downloader.NewDownloader(path, 1, 50*time.Millisecond)

	_, err := d.Run(ctx)

	assert.Nil(t, err)
}

func TestFileNotFind(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(200 * time.Millisecond)
	}))
	defer srv.Close()

	d, _ := downloader.NewDownloader("/nonexistent", 1, 1)
	_, err := d.Run(context.Background())

	assert.Error(t, err)
}

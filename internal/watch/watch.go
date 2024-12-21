package watch

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

type CudairWatch struct {
	W            *fsnotify.Watcher
	RootDir      string
	excludedDirs []string
}

func NewCudairWatch(root string, exclude []string) (*CudairWatch, error) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return nil, err
	}

	return &CudairWatch{
		W:            watcher,
		excludedDirs: exclude,
		RootDir:      root,
	}, nil
}

func (cw *CudairWatch) AddWatcherRecursively() error {
	err := filepath.Walk(cw.RootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			if isExcluded(path, cw.excludedDirs) {
				fmt.Printf("Excluding: %s\n", path)
				return filepath.SkipDir
			}

			err := cw.W.Add(path)
			if err != nil {
				return err
			}
			fmt.Printf("Added watching dir: %s\n", path)
		}
		return nil
	})
	return err
}

func isExcluded(path string, excludedDirs []string) bool {
	for _, excluded := range excludedDirs {

		if strings.HasPrefix(filepath.Clean(path), excluded) {
			return true
		}
	}
	return false
}

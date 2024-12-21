package cmd

import (
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"sync"
	"syscall"
	"time"

	"github.com/ei-sugimoto/cudair/internal/builder"
	"github.com/ei-sugimoto/cudair/internal/config"
	"github.com/ei-sugimoto/cudair/internal/executor"
	"github.com/ei-sugimoto/cudair/internal/watch"
	"github.com/fsnotify/fsnotify"
)

func Run(configFilePath string) error {
	log.Println("starting...")

	config, err := config.NewCudairConfig(configFilePath)
	if err != nil {
		log.Fatalln("cause Error while parsing toml:", err)
		return err
	}

	watcher, err := watch.NewCudairWatch(config.Root, config.Build.ExcludeDir)
	if err != nil {
		return err
	}

	err = watcher.AddWatcherRecursively()
	if err != nil {
		return err
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGTERM, syscall.SIGINT)

	var mu sync.Mutex
	var lastEventTime time.Time
	var lastEventFile string

	for {
		select {
		case e, ok := <-watcher.W.Events:
			if !ok {
				log.Fatal("watcher event is not ok")
			}
			if ((e.Op&fsnotify.Write == fsnotify.Write) || (e.Op&fsnotify.Remove == fsnotify.Remove) || (e.Op&fsnotify.Create == fsnotify.Create) || (e.Op&fsnotify.Rename == fsnotify.Rename)) && (filepath.Ext(e.Name) == ".cu" || filepath.Ext(e.Name) == ".cuh" || isDir(e.Name)) {
				mu.Lock()
				//　imp event debounce.
				if e.Name == lastEventFile && time.Since(lastEventTime) < 3*time.Second {
					mu.Unlock()
					continue
				}
				lastEventTime = time.Now()
				lastEventFile = e.Name
				if isDir(e.Name) {
					watcher.AddWatcherRecursively()
					mu.Unlock()
					continue

				}
				mu.Unlock()

				log.Printf("Changing %#v\n", e)
				if err := builder.Build(config.Build.Cmd, config.TmpDir); err != nil {
					log.Println("build error:", err)
					continue
				}
				if err := executor.Execute(config.Build.Bin); err != nil {
					log.Println("execution error:", err)
					continue

				}
			}
		case err := <-watcher.W.Errors:
			if err != nil {
				log.Fatalln("cause error while running:", err)
			}
		case <-sig:
			log.Println("Received termination signal, shutting down...")
			return nil
		}
	}
}

func isDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return info.IsDir()
}

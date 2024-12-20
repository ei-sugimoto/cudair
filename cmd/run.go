package cmd

import (
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"github.com/ei-sugimoto/cudair/internal/builder"
	"github.com/ei-sugimoto/cudair/internal/config"
	"github.com/ei-sugimoto/cudair/internal/executor"
	"github.com/fsnotify/fsnotify"
)

func Run(configFilePath string) error {
	log.Println("starting...")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		return err
	}
	defer watcher.Close()

	config, err := config.NewCudairConfig(configFilePath)
	if err != nil {
		return err
	}

	err = watcher.Add(config.Root)
	if err != nil {
		return err
	}

	go func() {
		for {
			select {
			case e, ok := <-watcher.Events:
				if !ok {
					log.Fatal("watcher event is not ok")
					return
				}
				if ((e.Op&fsnotify.Write == fsnotify.Write) || (e.Op&fsnotify.Remove == fsnotify.Remove) || (e.Op&fsnotify.Create == fsnotify.Create) || (e.Op&fsnotify.Rename == fsnotify.Rename)) && (filepath.Ext(e.Name) == ".cu" || filepath.Ext(e.Name) == ".cuh") {
					log.Printf("Changing %#v\n", e)
					if err := builder.Build(config.Build.Cmd); err != nil {
						log.Fatalln("build error:", err)
						return
					}
					if err := executor.Execute(config.Build.Cmd); err != nil {
						log.Fatalln("exec error", err)
						return
					}

				}
			case err := <-watcher.Errors:
				log.Fatalln("cause error:", err)
			}
		}
	}()
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c
	log.Println("getting SIGTERM...")

	return nil
}

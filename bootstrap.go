package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/enderian/directrd/api"
	"github.com/enderian/directrd/database"
	"github.com/enderian/directrd/delegation"
	"github.com/enderian/directrd/radius"
	"github.com/enderian/directrd/sessions"
	"github.com/enderian/directrd/types"
	"github.com/enderian/directrd/users"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"path"
	"sync"
	"syscall"
	"time"
)

var configFile = flag.String("config", "config.yml", "The configuration file.")

func main() {

	flag.Parse()
	configuration := &types.Configuration{}
	confBytes, err := ioutil.ReadFile(*configFile)

	if err != nil {
		log.Panic(err)
	}
	if err := yaml.Unmarshal(confBytes, configuration); err != nil {
		log.Panic(err)
	}

	//Setup the logger
	logger := &directrdLogWriter{folder: configuration.Logs}
	log.SetOutput(logger)
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	//Start all the things
	ctx := types.NewContext(context.Background(), configuration, logger)
	ctx = database.SetupDatabase(ctx)

	users.Setup(ctx)
	delegation.Setup(ctx)
	api.Setup(ctx)
	radius.Setup(ctx)
	sessions.Setup(ctx)

	//If shutdown signal received, exit gracefully.
	sign := make(chan os.Signal)
	signal.Notify(sign)
	for {
		switch <-sign {
		case syscall.SIGKILL:
		case syscall.SIGTERM:
		case syscall.SIGINT:
			log.Println("Server exiting...")
			database.CloseDatabase(ctx)
			return
		}
	}
}

/**
	Writes log files directly to the file without keeping it open.
**/
type directrdLogWriter struct {
	folder string
	file   *os.File
	mx     sync.Mutex
}

func (d *directrdLogWriter) Write(p []byte) (n int, err error) {
	d.mx.Lock()
	defer d.mx.Unlock()
	d.checkDate()
	if _, err := d.file.Write(p); err != nil {
		panic(fmt.Sprintf("Fatal error while writing log file: %s", err.Error()))
	}
	if _, err := os.Stdout.Write(p); err != nil {
		panic(fmt.Sprintf("Fatal error while writing log to console: %s", err.Error()))
	}
	return len(p), nil
}

func (d *directrdLogWriter) checkDate() {
	expectedName := path.Join(d.folder, fmt.Sprintf("log-directd-%s.txt", time.Now().Format("2006-01-02")))
	if d.file == nil || d.file.Name() != expectedName {
		var err error
		d.file, err = os.OpenFile(expectedName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
		if err != nil {
			panic(fmt.Sprintf("Fatal error while creating log file: %s", err.Error()))
		}
	}
}

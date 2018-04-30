// 觀察貓

package main

import (
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/radovskyb/watcher"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

var env envStruct

type envStruct struct {
	ObservedFile string `mapstructure:"OBSERVED_FILE" json:"OBSERVED_FILE"`
	EventAllExec string `mapstructure:"EVENT_ALL_EXEC" json:"EVENT_ALL_EXEC"`
}

func main() {
	app := cli.NewApp()
	app.Name = "OserverdCat"
	app.Version = "0.0.1"
	app.Authors = []cli.Author{
		cli.Author{
			Name: "Ken",
		},
	}
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "file",
			Value:  "ObservedFile",
			Usage:  "ObservedFile",
			EnvVar: "OBSERVED_FILE",
		},
		cli.StringFlag{
			Name:   "eventall",
			Value:  "echo 'EventAll'",
			Usage:  "EventAllExec",
			EnvVar: "EVENT_ALL_EXEC",
		},
	}

	app.Action = run

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func run(c *cli.Context) {
	viper.SetDefault("OBSERVED_FILE", c.String("file"))
	viper.SetDefault("EVENT_ALL_EXEC", c.String("eventall"))
	viper.AutomaticEnv()
	viper.SetEnvPrefix("OC")

	envTmp := &envStruct{}
	err := viper.Unmarshal(envTmp)

	if err != nil {
		panic(err)
	}

	env = *envTmp

	if env.ObservedFile == "ObservedFile" {
		log.Fatal("ObsevedCat 缺少監控檔案")
	}

	log.Println("ObsevedCat 設定成功")
	log.Println("EventAllExec", env.EventAllExec)

	w := watcher.New()
	w.SetMaxEvents(1)
	w.FilterOps(watcher.Create, watcher.Write, watcher.Remove, watcher.Rename, watcher.Chmod, watcher.Move)

	go func() {
		for {
			select {
			case event := <-w.Event:
				cmd := strings.Fields(env.EventAllExec)
				out, err := exec.Command(cmd[0], cmd[1:]...).Output()

				if err != nil {
					log.Fatal(err)
				}

				log.Println(event)
				log.Println(string(out))
			case err := <-w.Error:
				log.Fatalln(err)
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.Add(env.ObservedFile); err != nil {
		log.Fatalln(err)
	}

	if err := w.Start(time.Second * 1); err != nil {
		log.Fatalln(err)
	}
}

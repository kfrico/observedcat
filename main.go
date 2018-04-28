// 觀察貓

package main

import (
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

var env envStruct

type envStruct struct {
	ObservedFile    string `mapstructure:"OBSERVED_FILE" json:"OBSERVED_FILE"`
	EventAllExec    string `mapstructure:"EVENT_ALL_EXEC" json:"EVENT_ALL_EXEC"`
	EventCreateExec string `mapstructure:"EVENT_CREATE_EXEC" json:"EVENT_CREATE_EXEC"`
	EventWriteExec  string `mapstructure:"EVENT_WRITE_EXEC" json:"EVENT_WRITE_EXEC"`
	EventRemoveExec string `mapstructure:"EVENT_REMOVE_EXEC" json:"EVENT_REMOVE_EXEC"`
	EventRenameExec string `mapstructure:"EVENT_RENAME_EXEC" json:"EVENT_RENAME_EXEC"`
	EventChmodExec  string `mapstructure:"EVENT_CHMOD_EXEC" json:"EVENT_CHMOD_EXEC"`
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
		cli.StringFlag{
			Name:   "eventcreate",
			Value:  "echo 'EventCreateExec'",
			Usage:  "EventCreateExec",
			EnvVar: "EVENT_CREATE_EXEC",
		},
		cli.StringFlag{
			Name:   "eventwrite",
			Value:  "echo 'EventWriteExec'",
			Usage:  "EventWriteExec",
			EnvVar: "EVENT_WRITE_EXEC",
		},
		cli.StringFlag{
			Name:   "eventremove",
			Value:  "echo 'EventRemoveExec'",
			Usage:  "EventRemoveExec",
			EnvVar: "EVENT_REMOVE_EXEC",
		},
		cli.StringFlag{
			Name:   "eventrename",
			Value:  "echo 'EventRenameExec'",
			Usage:  "EventRenameExec",
			EnvVar: "EVENT_RENAME_EXEC",
		},
		cli.StringFlag{
			Name:   "eventchmod",
			Value:  "echo 'EventChmodExec'",
			Usage:  "EventChmodExec",
			EnvVar: "EVENT_CHMOD_EXEC",
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
	viper.SetDefault("EVENT_All_EXEC", c.String("eventall"))
	viper.SetDefault("EVENT_CREATE_EXEC", c.String("eventcreate"))
	viper.SetDefault("EVENT_WRITE_EXEC", c.String("eventwrite"))
	viper.SetDefault("EVENT_REMOVE_EXEC", c.String("eventremove"))
	viper.SetDefault("EVENT_RENAME_EXEC", c.String("eventrename"))
	viper.SetDefault("EVENT_CHMOD_EXEC", c.String("eventchmod"))
	viper.AutomaticEnv()
	viper.SetEnvPrefix("OC")

	envTmp := &envStruct{}
	err := viper.Unmarshal(envTmp)

	if err != nil {
		panic(err)
	}

	log.Println("ObsevedCat 設定成功")

	env = *envTmp

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)

	go func() {
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)

				cmd := strings.Fields(env.EventAllExec)
				out, err := exec.Command(cmd[0], cmd[1:]...).Output()

				if err != nil {
					log.Fatal(err)
				}

				log.Println(string(out))

				if event.Op&fsnotify.Create == fsnotify.Create {
					cmd := strings.Fields(env.EventCreateExec)
					out, err := exec.Command(cmd[0], cmd[1:]...).Output()

					if err != nil {
						log.Fatal(err)
					}

					log.Println("創建文件 : ", event.Name)
					log.Println(string(out))
				}
				if event.Op&fsnotify.Write == fsnotify.Write {
					cmd := strings.Fields(env.EventWriteExec)
					out, err := exec.Command(cmd[0], cmd[1:]...).Output()

					if err != nil {
						log.Fatal(err)
					}

					log.Println("寫入文件 : ", event.Name)
					log.Println(string(out))
				}
				if event.Op&fsnotify.Remove == fsnotify.Remove {
					cmd := strings.Fields(env.EventRemoveExec)
					out, err := exec.Command(cmd[0], cmd[1:]...).Output()

					if err != nil {
						log.Fatal(err)
					}

					log.Println("刪除文件 : ", event.Name)
					log.Println(string(out))
				}
				if event.Op&fsnotify.Rename == fsnotify.Rename {
					cmd := strings.Fields(env.EventRenameExec)
					out, err := exec.Command(cmd[0], cmd[1:]...).Output()

					if err != nil {
						log.Fatal(err)
					}

					log.Println("重命名文件 : ", event.Name)
					log.Println(string(out))
				}
				if event.Op&fsnotify.Chmod == fsnotify.Chmod {
					cmd := strings.Fields(env.EventChmodExec)
					out, err := exec.Command(cmd[0], cmd[1:]...).Output()

					if err != nil {
						log.Fatal(err)
					}

					log.Println("修改權限 : ", event.Name)
					log.Println(string(out))
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}

		}
	}()

	err = watcher.Add(env.ObservedFile)

	if err != nil {
		log.Fatal(err)
	}
	<-done

}

package main

import (
	"errors"
	"os/exec"

	"github.com/gin-gonic/gin"
	"github.com/tkanos/gonfig"
)

type Configuration struct {
	DataDevice       string
	DataMountpoint   string
	BackupDevice     string
	BackupMountpoint string
	Port             int
}

func unlockDrive(device string, mountpoint string, password string) (int, error) {
	unlock_command := exec.Command("cryptsetup", "luksOpen", device, mountpoint, "-d -")
	unlock_stdin, err := unlock_command.StdinPipe()
	if nil != err {
		return 1, err
	}
	unlock_stdin.Write([]byte(password))
	if err := unlock_command.Run(); nil != err {
		return 1, err
	}
	if err := unlock_command.Wait(); nil != err {
		if exiterr, ok := err.(*exec.ExitError); ok {
			return exiterr.ExitCode(), errors.New("Couldn't unlock drive")
		}
	}
	return 0, nil
}

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.Static("/static", "./static")

	r.GET("/", func(c *gin.Context) {
		c.Redirect(301, "/static/index.html")
	})

	r.POST("/unlock", func(c *gin.Context) {
		data_password := c.PostForm("data_password")
		backup_password := c.PostForm("backup_password")

		data_unlock_exit, err := unlockDrive("/dev/sda1", "data", data_password)
		if err != nil {
			panic("Couldn't unlock data drive, bad password?")
		}
		backup_unlock_exit, err := unlockDrive("/dev/sdb1", "backup", backup_password)
		if err != nil {
			panic("Couldn't unlock backup drive, bad password?")
		}
	})

	return r
}

func main() {
	configuration := Configuration{}
	err := gonfig.GetConf("/portadisk/config/web_unlock.json", &configuration)
	if err != nil {
		panic(err)
	}

	r := setupRouter()
	r.Run(":" + string(configuration.Port))
}

/*
* ________ _________  ________  ________  ________
*|\   ____\\___   ___\\   __  \|\   ___ \|\   __  \
*\ \  \___\|___ \  \_\ \  \|\  \ \  \_|\ \ \  \|\  \
* \ \  \  ___  \ \  \ \ \  \\\  \ \  \ \\ \ \  \\\  \
*  \ \  \|\  \  \ \  \ \ \  \\\  \ \  \_\\ \ \  \\\  \
*   \ \_______\  \ \__\ \ \_______\ \_______\ \_______\
*    \|_______|   \|__|  \|_______|\|_______|\|_______|
*
 */

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/shiniao/gtodo/config"
	"github.com/shiniao/gtodo/internal/model"
	"github.com/shiniao/gtodo/pkg/log"
	v "github.com/shiniao/gtodo/pkg/version"
	"github.com/shiniao/gtodo/router"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"net/http"
	"os"
	"time"
)

// * 读取命令行配置（config文件）
var (
	cfg     = pflag.StringP("config", "c", "", "apiserver config file path.")
	version = pflag.BoolP("version", "v", false, "show version info.")
)

// @title A todos application API
// @version 1.0
// @description This is a todos application server.
// @termsOfService http://me.shiniao.fun/

// @contact.name shiniao
// @contact.url http://me.shiniao.fun/
// @contact.email zhuzhezhe5@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host todo.shiniao.fun
// @BasePath /v1
func main() {

	// * 初始化配置
	pflag.Parse()
	// version
	if *version{
		v := v.Get()
		marshaled, err := json.MarshalIndent(&v, "", " ")
		if err != nil{
			fmt.Printf("%v\n", err)
			os.Exit(1)
		}
		fmt.Println(marshaled)
		return
	}

	if err := config.Init(*cfg); err != nil {
		panic(err)
	}
	log.Info("The config inited.")

	// * 初始化数据库
	model.DB.Init()
	defer model.DB.Close()
	log.Info("The database inited.")

	// * 设置 run mode
	gin.SetMode(viper.GetString("run_mode"))
	log.Info("Gin run mode set to: %s", viper.GetString("run_mode"))

	// * gin engine.
	g := gin.New()
	router.Load(g)
	log.Info("The gin engine started, and the router loaded.")

	go func() {
		if err := pingServer(); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	log.Info("Start to listening on: %s", viper.GetString("addr"))
	g.Run(viper.GetString("addr")).Error()
}

// * api 状态自检
func pingServer() error {
	for i := 0; i < viper.GetInt("max_ping_count"); i++ {
		resp, err := http.Get(viper.GetString("url") + "/sd/health")
		if err == nil && resp.StatusCode == 200 {
			return nil
		}

		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}

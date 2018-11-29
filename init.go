package DocusServer

import (
	"flag"
	"fmt"
	"git.docus.tech/kdl12138/DocusServer/Api"
	"git.docus.tech/kdl12138/DocusServer/Core"
	"git.docus.tech/kdl12138/DocusServer/Template"
	"gopkg.in/mgo.v2"
	"net/http"
	"strings"
)

func main() {
	Template.MongoDialInfo = make(map[string]mgo.DialInfo)
	Template.Session = make(map[string]*mgo.Session)
	configFilePath := flag.String("config", "", "配置文件路径，必须指定，无默认值！")
	flag.Parse()

	if *configFilePath == "" {
		fmt.Println("请指定配置文件位置，格式为[--config 配置文件路径.json！]")
		return
	}

	err := Api.Read("config.json")

	if err != nil {
		//添加log与输出
		return
	}
	//添加操作log

	mux := http.NewServeMux()
	if strings.EqualFold(Template.Setting.Role, "Storage") {
		go func() {
			mux.HandleFunc("/", Core.Get)
			http.ListenAndServe(":12345", mux)
		}()

	} else {
		go Core.GetAllServers()
		go func() {
			//mux.Handle("/",)
		}()
	}
}

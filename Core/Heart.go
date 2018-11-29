package Core

import (
	"bytes"
	"encoding/json"
	"fmt"
	"git.docus.tech/kdl12138/DocusServer/Api"
	"git.docus.tech/kdl12138/DocusServer/Template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type Heart struct {
	Address string
	Port    int
	Rand    int
}

func Get(w http.ResponseWriter, r *http.Request) {
	if !Api.CheckWhite(r.RemoteAddr) {
		//TODO 异常处理
	} else {
		body, _ := ioutil.ReadAll(r.Body)
		var heart Heart
		if err := json.Unmarshal(body, &heart); err == nil {
			heart.Rand += heart.Port
			if ret, err := json.Marshal(heart); err != nil {
				//TODO 异常处理
			} else {
				fmt.Fprint(w, string(ret))
			}
		} else {
			//TODO 异常处理
		}
	}
}

/**
轮询所有存储服务器
一秒询问一台
如果地址，端口号，和产生的随机值不与处理后的值一样
或是没有返回
则失联或状态错误
*/
func GetAllServers() {
	for {
		for _, i := range Template.Setting.Nodes.Node {

			rand.Seed(time.Now().Unix())
			rnd := rand.Intn(100)
			port, _ := strconv.Atoi(i.Port)
			temp := Heart{i.Address, port, rnd}
			if bs, err := json.Marshal(temp); err != nil {
				//TODO 异常处理
			} else {
				req := bytes.NewBuffer([]byte(bs))
				body_type := "application/json;charset=utf-8"
				resp, err := http.Post(i.Address+":"+i.Port+"/Core/Heart/Get", body_type, req)
				if err != nil {
					//TODO 异常处理
				}
				if body, err := ioutil.ReadAll(resp.Body); err != nil {
					//TODO 异常处理
				} else {
					var heart Heart
					if err := json.Unmarshal(body, &heart); err == nil {
						if !strings.EqualFold(heart.Address, temp.Address) || temp.Port != port || temp.Rand+temp.Port != heart.Rand {
							//TODO 异常处理
						}
					} else {
						//TODO 异常处理
					}
				}
			}
			time.Sleep(time.Second)
		}
	}
}

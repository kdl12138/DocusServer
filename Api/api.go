package Api

import (
	"encoding/json"
	"git.docus.tech/kdl12138/DocusServer/Template"
	"github.com/TokenTeam/Token-Static-Center/util"
	"net/http"
	"strings"
)

func JsonReturn(w http.ResponseWriter, r *http.Request, module string, trace string, errNumber int, message string) {

	// 返回Json数据格式
	type Json struct {
		Errno   int    `json:"error_code"`
		Message string `json:"message"`
	}

	data := Json{errNumber, message}

	jsonDataByte, err := json.Marshal(data)

	if err != nil {
		util.ErrorLog("JsonReturn", "生成Json数据时出现错误："+err.Error(), "app->JsonReturn")
		return
	}

	// 记录错误日志
	if errNumber < 0 {
		util.ErrorLog(module, message, trace)
	}

	w.Write(jsonDataByte)

	return
}

func CheckWhite(r string) bool {
	for _, i := range Template.Setting.White_list.S {
		if strings.EqualFold(r, i) {
			return true
		}
	}
	return false
}

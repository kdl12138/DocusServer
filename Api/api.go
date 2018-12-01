package Api

import (
	"encoding/json"
	"git.docus.tech/kdl12138/DocusServer/Storage"
	"git.docus.tech/kdl12138/DocusServer/Template"
	"github.com/TokenTeam/Token-Static-Center/util"
	"net/http"
	"sort"
	"strings"
	"sync"
)
type Storages []Storage.StorageStruct

func (a Storages) Len() int      { return len(a) }
func (a Storages) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a Storages) Less(i, j int) bool { return a[i].RestMax > a[j].RestMax } // 从大到小排序
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

func Find(size int64 ) (node, block string, flag int, err error){
	var mux sync.Mutex
	mux.Lock()
	if size <= Storage.Storages[0].RestMax{
		node = Storage.Storages[0].Node
		block = Storage.Storages[0].Block
		flag = Template.NewBlock_FALSE
		Storage.Storages[0].RestMax -= size
		Storage.Storages[0].RestOffset += size + 1
		Storage.StorageMap[node][block] = Storage.Storage{
			RestMax: Storage.Storages[0].RestMax,
			RestOffset: Storage.Storages[0].RestOffset,
		}
		sort.Sort(Storages(Storage.Storages))
		mux.Unlock()
		return node, block, flag, err
	} else {
		flag = Template.NewBlock_TRUE
		return "", "", flag, nil
	}



}

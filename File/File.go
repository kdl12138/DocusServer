package File

import (
	"encoding/json"
	"fmt"
	"git.docus.tech/kdl12138/DocusServer/Model"
	"git.docus.tech/kdl12138/DocusServer/Template"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

// 数据文件的接口类型。
type DataFile interface {
	// 读取一个数据块。
	Read() (rsn int64, d Data, err error)
	// 写入一个数据块。
	Write(d Data) (wsn int64, err error)
	// 获取最后读取的数据块的序列号。
	Rsn() int64
	// 获取最后写入的数据块的序列号。
	Wsn() int64
	// 获取数据块的长度
	DataLen() uint32
}

type myDataFile struct {
	f       *os.File     // 文件。
	fmutex  sync.RWMutex // 被用于文件的读写锁。
	woffset int64        // 写操作需要用到的偏移量。
	roffset int64        // 读操作需要用到的偏移量。
	wmutex  sync.Mutex   // 写操作需要用到的互斥锁。
	rmutex  sync.Mutex   // 读操作需要用到的互斥锁。
	dataLen uint32       // 数据块长度。
}
type Data []byte

func Delete(delData Template.Data) {
	var fmutex sync.RWMutex
	fmutex.Lock()
	defer fmutex.Unlock()
	str := delData.Block[0:2]
	f, _ := os.Open(Template.DocusURL + str + "/" + delData.Block)
	defer f.Close()
	//从头开始，文件指针偏移100
	f.Seek(delData.OffsetStart, 0)
	len := delData.OffsetEnd - delData.OffsetStart
	buffer := make([]byte, len)
	// Read 后文件指针也会偏移
	data, err := f.Read(buffer)
	if err != nil {
		fmt.Println(nil)
		return
	}
	// 获取文件指针当前位置
	cur_offset, _ := f.Seek(0, os.SEEK_CUR)
}
func Update(w http.ResponseWriter, r *http.Request) {
	var updateMessage Template.UpdataMessage
	con, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// TODO log
	}
	str := string(con)
	if err := json.Unmarshal([]byte(str), &updateMessage); err != nil {
		// TODO log
	}
	delData, err := Model.Find(updateMessage.Uuid)
	if err != nil {
		// TODO log
	}
	delData.Backup = 0
	if err := Model.Update(delData); err != nil {
		// TODO log
	}
	fmt.Fprint(w, "Update complete")
}

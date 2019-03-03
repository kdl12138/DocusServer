package File

import (
	"git.docus.tech/kdl12138/DocusServer/Storage"
	"git.docus.tech/kdl12138/DocusServer/Template"
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
// TODO return
// all blocks update to zero
func Delete(delData Template.StorageData) {
	var fmutex sync.RWMutex
	var zero string
	fmutex.Lock()
	defer fmutex.Unlock()
	str := delData.Block[0:2]
	f, _ := os.Open(Template.DocusURL + str + "/" + delData.Block)
	defer f.Close()
	for i := 0; i < 64; i++ {
		zero += "\0"
	}
	length := delData.OffsetEnd - delData.OffsetStart
	var ix int64 = 0
	var j int64 = 0
	for {
		if length/64 == 0 {
			for j = 0; j < length; j++ {
				f.WriteAt([]byte("\0"), delData.OffsetStart + ix + j)
			}
		} else {
			ix += 64
			f.WriteAt([]byte(zero), delData.OffsetStart+ix-64)
		}
	}
	// TODO update Storage
}
// first if should new block, will find one, then delete, finally update
func Update(updateData Template.StorageData) {
	var fmutex sync.RWMutex
	fmutex.Lock()
	defer fmutex.Unlock()
	if updateData.Flag == Template.NewBlock_TRUE {
		Storage.NewBlock(updateData.Block)
	}
	str := updateData.Block[0:2]
	Delete(updateData)
	f, _ := os.Open(Template.DocusURL + str + "/" + updateData.Block)
	defer f.Close()
	data := []byte(updateData.Data)
	length := updateData.OffsetEnd - updateData.OffsetStart
	var ix int64 = 0
	for {
		if length/64 == 0 {
			f.WriteAt(data[ix:], updateData.OffsetStart+ix)
		} else {
			ix += 64
			f.WriteAt(data[ix-64:ix], updateData.OffsetStart+ix-64)
		}
	}
}
// get the data from one block
func Read(data Template.StorageData) []byte {
	var fmutex sync.RWMutex
	fmutex.Lock()
	defer fmutex.Unlock()
	str := data.Block[0:2]
	f, _ := os.Open(Template.DocusURL + str + "/" + data.Block)
	//从头开始，文件指针偏移
	f.Seek(data.OffsetStart, 0)
	length := data.OffsetEnd - data.OffsetStart
	buffer := make([]byte, 1024)
	result := make([]byte, 0)
	//循环读取
	for {
		if length < int64(1024) {
			buffer = make([]byte, length)
		}
		len, _ := f.Read(buffer)
		result = append(result, buffer...)
		//读取字节数为0时跳出循环
		if len == 0 {
			break
		}
	}
	return result
}
// first if should new block ,will find one, then add
func Add(addData Template.StorageData) {
	var fmutex sync.RWMutex
	fmutex.Lock()
	defer fmutex.Unlock()
	if addData.Flag == Template.NewBlock_TRUE {
		Storage.NewBlock(addData.Block)
	}
	str := addData.Block[0:2]
	Delete(addData)
	f, _ := os.Open(Template.DocusURL + str + "/" + addData.Block)
	defer f.Close()
	data := []byte(addData.Data)
	length := addData.OffsetEnd - addData.OffsetStart
	var ix int64 = 0
	for {
		if length/64 == 0 {
			f.WriteAt(data[ix:], addData.OffsetStart+ix)
		} else {
			ix += 64
			f.WriteAt(data[ix-64:ix], addData.OffsetStart+ix-64)
		}
	}
}

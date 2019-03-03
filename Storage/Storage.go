package Storage

import (
	"fmt"
	"git.docus.tech/kdl12138/DocusServer/Template"
	"os"
	"sync"
)

// TODO new a block
func NewBlock(block string) (err error) {
	temp_dir := block[0:2]
	_, err = os.Stat(temp_dir)
	if err != nil {
		// TODO log
		if os.IsNotExist(err) {
			fmt.Println("temp dir is not exist")
			err := os.Mkdir(temp_dir, os.ModePerm)
			if err != nil {
				// TODO log
			}
			NewFile(temp_dir, block)
			return
		}
		// TODO log
		return
	}
	return nil
}

func NewFile(dir, fileName string) {
	var fmutex sync.RWMutex
	fmutex.Lock()
	defer fmutex.Unlock()
	f, _ := os.Open(Template.DocusURL + dir + "/" + fileName)
	defer f.Close()
	zero := ""
	for i := 0; i < 64; i++ {
		zero += "\0"
	}

	for length := int64(64 * 1024 * 1024); length > 0; length /= 64 {
		f.Write([]byte(zero))
	}
}

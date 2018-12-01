package Storage

import "git.docus.tech/kdl12138/DocusServer/Model"

type Storage struct {
	RestMax int64
	RestOffset int64
}
type StorageStruct struct {
	Node string
	Block string
	RestMax int64
	RestOffset int64
}
var StorageMap map[string]map[string]Storage
var Storages []StorageStruct
func New()  {
	StorageMap = make(map[string]map[string]Storage)
	StorageData, err := Model.FindAll()
	if err != nil{
		// TODO log
	}
	for _, i := range StorageData{
		StorageMap[i.Node][i.Block] = Storage{
			RestMax: StorageMap[i.Node][i.Block].RestMax - i.Size,
			RestOffset: i.OffsetEnd + 1,
		}
	}
	for k, v := range StorageMap {
		for k1, v1 := range v {
			Storages = append(Storages,StorageStruct{
				k,
				k1,
				v1.RestMax,
				v1.RestOffset,
			})
		}
	}
}


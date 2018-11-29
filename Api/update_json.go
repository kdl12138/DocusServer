package Api

import (
	"git.docus.tech/kdl12138/DocusServer/Template"
	"sync"
)

var (
	Config     *Template.Server
	configLock = new(sync.RWMutex)
)

func loadConfig() (err error) {
	Read("config.json")
	return nil
}

//func GetConfig() *Template.Storage {
//	configLock.RLock()
//	defer configLock.RUnlock()
//	return Config
//}

//func init() {
//	if !loadConfig() {
//		os.Exit(1)
//	}
//
//	s := make(chan os.Signal, 1)
//	signal.Notify(s, syscall.SIGUSR1)
//	go func() {
//		for {
//			<-s
//			log.Println("Reloaded config:", loadConfig())
//		}
//	}()
//}

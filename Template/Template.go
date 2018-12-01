package Template

type White_List struct {
	S []string `json:"s"` //白名单
}
type Servicer struct {
	Name    string `json:"name"` //
	Address string `json:"address"`
	Port    string `json:"port"`
	Status  bool   `json:"status"`
}
type Nodes struct {
	Node []Servicer `json:"node"` // 节点
}
type Server struct {
	Listen_address string     `json:"listen_address"` // 目录服务器监听地址
	Listen_port    string     `json:"listen_port"`    // 目录服务器监听端口
	White_list     White_List `json:"white_list"`     // 白名单内的服务器才允许访问目录服务器
	Gzip           string     `json:"gzip"`           // 是否开启gzip
	Log_lever      string     `json:"log_level"`      //  0-不记录日志 1-记录错误日志 2-记录操作日志 3-记录所有日志
	Log_dir        string     `json:"log_dir"`        // 自动按照log格式保存时间
	Nodes          Nodes      `json:"nodes"`          // 节点
	Role           string     `json:"role"`           // 服务器充当的角色
}
type Serverslice struct {
	Servers []Server `json:"servers"` // 服务器
}

var Setting = Server{}

type GetMessage struct {
	Identity  string `json:"identity"`  // 身份
	SessionId string `json:"session"`   // session号
	Action    string `json:"action"`    // 方式
	Uuid      string `json:"uuid"`      // 全局序号
	Timestamp int    `json:"timestamp"` // 时间戳
}
type UpdataMessage struct {
	Identity  string `json:"identity"`  // 身份
	SessionId string `json:"session"`   // session号
	Action    string `json:"action"`    // 方式
	Uuid      string `json:"uuid"`      // 全局序号
	Data      string `json:"data"`      // 数据
	Timestamp int    `json:"timestamp"` // 时间戳
	Md5       string `json:"md5"`       // md5
}
type DelMessage struct {
	Identity  string `json:"identity"`  // 身份
	SessionId string `json:"session"`   // session号
	Action    string `json:"action"`    // 方式
	Uuid      string `json:"uuid"`      // 全局序号
	Timestamp int    `json:"timestamp"` // 时间戳
}
type AddMessage struct {
	Identity  string `json:"identity"`  // 身份
	SessionId string `json:"session"`   // session号
	Action    string `json:"action"`    // 方式
	Uuid      string `json:"uuid"`      // 全局序号
	Data      string `json:"data"`      // 数据
	Timestamp int    `json:"timestamp"` // 时间戳
	Md5       string `json:"md5"`       // md5
}
type Data struct {
	Identity    string `json:"identity"`     // 身份
	SessionId   string `json:"session"`      // session号
	//Action      string `json:"action"`       // 方式
	Uuid        string `json:"uuid"`         // 全局序号
	Backup      int    `json:"backup"`       // 备份号
	BlockStatus string `json:"block_status"` // 区块状态
	Node        string `json:"node"`         // 所在节点
	Block       string `json:"block"`        // 所在区块
	OffsetStart int64  `json:"offset_start"` // 偏移起点
	OffsetEnd   int64  `json:"offset_end"`   // 偏移终点
	Gzip        string `json:"gzip"`         // gzip开启状态
	Timestamp   int    `json:"timestamp"`    // 时间戳
	Md5         string `json:"md5"`          // md5
}
type MasterData struct {
	Uuid           string `json:"uuid"`             // 全局序号
	Backup         int    `json:"backup"`           // 备份号
	Node           string `json:"node"`             // 所在节点
	Block          string `json:"block"`            // 所在区块
	OffsetStart    int64  `json:"offset_start"`     // 偏移起点
	OffsetEnd      int64  `json:"offset_end"`       // 偏移终点
	Size           int64  `json:"size"`             // 文件大小
	Gzip           string `json:"gzip"`             // gzip开启状态
	EstablishTime  int    `json:"establish_time"`   // 创建时间戳
	LastAccessTime int    `json:"last_access_time"` // 创建时间戳
	LastModifyTime int    `json:"last_modify_time"` // 上一次修改时间戳
	AccessCount    int    `json:"access_count"`     // 访问计数
	ModifyCount    int    `json:"modify_count"`     // 修改计数
	Md5            string `json:"md5"`              // md5
}
type StorageData struct {
	Block       string `json:"block"`        // 所在区块
	OffsetStart int64  `json:"offset_start"` // 偏移起点
	OffsetEnd   int64  `json:"offset_end"`   // 偏移终点
	flag        int    `json:"flag"`         // 是否开辟新的区块
}

const Database string = "data"
const DocusURL string = "var/docus/file/"
const NewBlock_TRUE = 1
const NewBlock_FALSE = 0

package config

// config.jsonファイルのルートレベルに対応する構造体
type Config struct {
	ConnectServer ConnectServer `json:"connect_server"`
	WebServer     WebServer     `json:"web_server"`
	Log           Log           `json:"log"`
}

type ConnectServer struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
}

type WebServer struct {
	Ip   string `json:"ip"`
	Port int    `json:"port"`
	Dir  string `json:"dir"`
}

type Log struct {
	Dir      string `json:"dir"`
	FileName string `json:"file_name"`
}

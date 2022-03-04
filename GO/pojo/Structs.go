package pojo

// UserData 用户数据项结构体
type UserData struct {
	UserId      string
	Password    string
	Service     string
	QueryString string
}

// ConfigData 配置文件结构体
type ConfigData struct {
	UserId       string `yaml:"UserId"`
	Password     string `yaml:"Password"`
	Service      string `yaml:"Service"`
	TimeInterval int    `yaml:"TimeInterval"`
	LogPath      string `yaml:"LogPath"`
	LogSaveDay   int    `yaml:"LogSaveDay"`
	LogClearDay  int    `yaml:"LogClearDay"`
}

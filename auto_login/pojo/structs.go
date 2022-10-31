package pojo

// UserData 用户数据项结构体
type UserData struct {
	UserId      string
	Password    string
	Server      string
	QueryString string
}

// ConfigData 配置文件结构体
type ConfigData struct {
	UserId       string `yaml:"UserId"`
	Password     string `yaml:"Password"`
	Server       string `yaml:"Server"`
	TimeInterval int    `yaml:"TimeInterval"`
	LogPath      string `yaml:"LogPath"`
	LogClearDay  int    `yaml:"LogClearDay"`
}

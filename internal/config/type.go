package config

import "fmt"

// HTTPType http 配置信息
type HTTPType struct {
	Port             string `yaml:"Port"`
	MaxContentLength int    `yaml:"MaxContentLength"`
	ShutdownTimeout  int    `yaml:"ShutdownTimeout"`
	MaxLoggerLength  int    `yaml:"MaxLoggerLength"`
}

// CORSType 跨域设置
type CORSType struct {
	Enable           bool     `yaml:"Enable"`
	AllowOrigins     []string `yaml:"AllowOrigins"`
	AllowMethods     []string `yaml:"AllowMethods"`
	AllowHeaders     []string `yaml:"AllowHeaders"`
	AllowCredentials bool     `yaml:"AllowCredentials"`
	MaxAge           int      `yaml:"MaxAge"`
}

// GORMType gorm 配置信息
type GORMType struct {
	Debug             bool `yaml:"Debug"`
	MaxLifetime       int  `yaml:"MaxLifetime"`
	MaxOpenConns      int  `yaml:"MaxOpenConns"`
	MaxIdleConns      int  `yaml:"MaxIdleConns"`
	EnableAutoMigrate bool `yaml:"EnableAutoMigrate"`
}

// DBType 数据库配置定义
type DBType struct {
	Host       string `yaml:"Host"`
	Port       int    `yaml:"Port"`
	User       string `yaml:"User"`
	Password   string `yaml:"Password"`
	DBName     string `yaml:"DBName"`
	Parameters string `yaml:"Parameters"`
}

// DSN 得到数据库连接
func (d *DBType) DSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%s",
		d.User,
		d.Password,
		d.Host,
		d.Port,
		d.DBName,
		d.Parameters,
	)
}

// LogType 日志配置类型定义
type LogType struct {
	Level  int8   `yaml:"Level"`
	Output string `yaml:"Output"`
}

// LogFileHookType 文件归档钩子配置
type LogFileHookType struct {
	Filename   string `yaml:"Filename"`
	MaxSize    int    `yaml:"Maxsize"`
	MaxBackups int    `yaml:"MaxBackups"`
	MaxAge     int    `yaml:"Maxage"`
	Compress   bool   `yaml:"Compress"`
}

// JWTType 文件归档钩子配置
type JWTType struct {
	Enable  bool   `yaml:"Enable"`
	Secret  string `yaml:"Secret"`
	Expires int    `yaml:"Expires"`
	Issuer  string `yaml:"Issuer"`
}

// CType 配置文件类型定义
type CType struct {
	Mode        string          `yaml:"Mode"`
	HTTP        HTTPType        `yaml:"HTTP"`
	CORS        CORSType        `yaml:"CORS"`
	GORM        GORMType        `yaml:"GORM"`
	DB          DBType          `yaml:"DB"`
	Log         LogType         `yaml:"Log"`
	LogFileHook LogFileHookType `yaml:"LogFileHook"`
	JWT         JWTType         `yaml:"JWT"`
}

package global

import (
	"time"
)

var (
	ServerSetting *ServerSettingS
	AppSetting      *AppSettingS
	DatabaseSetting *DatabaseSettingS
	MasscanSetting *MasscanSettingS
)


type ServerSettingS struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

type AppSettingS struct {
	DefaultPageSize       int
	MaxPageSize           int
	DefaultContextTimeout time.Duration
	JwtSecret string
	PortUserDict          string
	PortPassDict          string
}

type DatabaseSettingS struct {
	DBType       string
	UserName     string
	Password     string
	Host         string
	DBName       string
	TablePrefix  string
	Charset      string
	ParseTime    bool
	MaxIdleConns int
	MaxOpenConns int
}

type MasscanSettingS struct {
	Rate      string
	IpFile    string
	IpNotScan string
	Port      string
}
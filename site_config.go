package config

import (
	"fmt"
	"github.com/go-ini/ini"
	"github.com/gorilla/websocket"
	"net/http"
	"os"
)

var ConfigPath string

//原始文件，如果有个性配置，在此读取
var OriginFile *ini.File

var (
	DbHost string
	DbPort string
	DbName string
	DbUser string
	DbPass string
	DbType string
)

var (
	SiteRootPath            string
	SiteUrl                 string
	SiteStaticUrl           string
	SiteTitle               string
	SiteListenIp            string
	SiteListenPort          string
	SiteUploadDir           string
	SiteSecretKey           string
	SitePublicDir           string
	SiteReportDir           string
	SiteMaxFileUploadSizeMb int64
	SiteWssUrl              string
	AdminInitialPassword    string
)

var (
	SmsSid   string
	SmsToken string
	SmsUrl   string
	SmsAppid string
)

var (
	WxAccount   string
	WxAppid     string
	WXAppsecret string
	WxVerifyUrl string
	WxApptoken  string
)

//redis配置

var (
	RedisIp       string
	RedisPort     string
	RedisDb       int
	RedisPassword string
	RedisMaxIdle  int
)

//Email 配置

var (
	EmailSmtpHost string
	EmailSmtpPort string
	EmailSmtpPass string
	EmailPop3Host string
	EmailPop3Port string
	EmailPop3Pass string
	EmailName     string
	EmailAddress  string
)

//thrift
var (
	ThriftListenPort   string
	ThriftSslServerKey string
	ThriftSslServerCrt string
	ThriftSslCaCrt     string
	ThriftSslClientCrt string
)

// RPC
var (
	RPCProtocol   string
	RPCListenHost string
	RPCListenPort string
)

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func LoadConfig(appPath, configDirectory, configName string) {

	ConfigPath = appPath + "/" + configDirectory + "/"
	var err error
	OriginFile, err = ini.Load(ConfigPath + configName)
	if err != nil {
		panic("Load config error:" + err.Error())
	}

	SiteRootPath = appPath

	//数据库配置
	key, err := OriginFile.GetSection("database")
	if err == nil {
		dbHostKey, err := key.GetKey("db_host")
		dbPortKey, err := key.GetKey("db_port")
		dbNameKey, err := key.GetKey("db_name")
		dbUserKey, err := key.GetKey("db_user")
		dbPassKey, err := key.GetKey("db_pass")
		dbTypeKey, err := key.GetKey("db_type")
		if err == nil {
			DbHost = dbHostKey.String()
			DbPort = dbPortKey.String()
			DbName = dbNameKey.String()
			DbUser = dbUserKey.String()
			DbPass = dbPassKey.String()
			DbType = dbTypeKey.String()
		}

	}

	siteKey, err := OriginFile.GetSection("site")
	if err == nil {
		UrlKey, err := siteKey.GetKey("url")
		StaticUrlKey, err := siteKey.GetKey("static_url")
		TitleKey, err := siteKey.GetKey("title")
		ListenIpKey, err := siteKey.GetKey("listen_ip")
		ListenPortKey, err := siteKey.GetKey("listen_port")
		UploadDirKey, err := siteKey.GetKey("upload_dir")
		SecretKey, err := siteKey.GetKey("secret_string")
		PublicDirKey, err := siteKey.GetKey("public_dir")
		ReportDirKey, err := siteKey.GetKey("report_dir")
		maxFileSizeKey, err := siteKey.GetKey("max_upload_file_size_mb")
		wssKey, err := siteKey.GetKey("wss_url")
		adminInitialPasswordKey, err := siteKey.GetKey("admin_initial_password")
		if err == nil {
			SiteUrl = UrlKey.String()
			SiteListenIp = ListenIpKey.String()
			SiteListenPort = ListenPortKey.String()
			SiteUploadDir = UploadDirKey.String()
			SiteSecretKey = SecretKey.String()
			SiteTitle = TitleKey.String()
			SiteStaticUrl = StaticUrlKey.String()
			SitePublicDir = PublicDirKey.String()
			SiteReportDir = ReportDirKey.String()
			AdminInitialPassword = adminInitialPasswordKey.String()
			SiteMaxFileUploadSizeMb, _ = maxFileSizeKey.Int64()
			if SiteMaxFileUploadSizeMb == 0 {
				SiteMaxFileUploadSizeMb = 5
			}
			SiteWssUrl = wssKey.String()
		}

	}

	//短信配置

	smsKey, err := OriginFile.GetSection("sms")
	if err == nil {
		SmsSidKey, err := smsKey.GetKey("account_sid")
		SmsTokenKey, err := smsKey.GetKey("auth_token")
		SmsUrlKey, err := smsKey.GetKey("rest_url")
		SmsAppIdKey, err := smsKey.GetKey("app_id")
		if err == nil {
			SmsSid = SmsSidKey.String()
			SmsToken = SmsTokenKey.String()
			SmsUrl = SmsUrlKey.String()
			SmsAppid = SmsAppIdKey.String()
		}
	}

	//RPC配置

	rpcKey, err := OriginFile.GetSection("rpc")
	if err == nil {
		RPCListenHostKey, err := rpcKey.GetKey("listen_host")
		RPCListenPortKey, err := rpcKey.GetKey("listen_port")
		RPCProtocolKey, err := rpcKey.GetKey("protocol")
		if err == nil {
			RPCListenHost = RPCListenHostKey.String()
			RPCListenPort = RPCListenPortKey.String()
			RPCProtocol = RPCProtocolKey.String()
		}
	}

	wxKey, err := OriginFile.GetSection("wechat")
	if err == nil {
		WxAccountKey, err := wxKey.GetKey("wxaccount")
		WxAppidKey, err := wxKey.GetKey("appid")
		WXAppsecretKey, err := wxKey.GetKey("appsecret")
		WxVerifyUrlKey, err := wxKey.GetKey("verifyurl")
		WxApptokenKey, err := wxKey.GetKey("apptoken")
		if err == nil {
			WxAccount = WxAccountKey.String()
			WxAppid = WxAppidKey.String()
			WXAppsecret = WXAppsecretKey.String()
			WxVerifyUrl = WxVerifyUrlKey.String()
			WxApptoken = WxApptokenKey.String()
		}

	}

	redisKey, err := OriginFile.GetSection("redis")
	if err == nil {
		RedisIpKey, err := redisKey.GetKey("ip")
		RedisportKey, err := redisKey.GetKey("port")
		RedisDbKey, err := redisKey.GetKey("db")
		RedisPasswordKey, err := redisKey.GetKey("password")
		RedisMaxIdleKey, err := redisKey.GetKey("maxidle")
		if err == nil {

			RedisIp = RedisIpKey.String()
			RedisPort = RedisportKey.String()
			redisdb, err := RedisDbKey.Int()
			if err == nil || redisdb > 20 || redisdb < 0 {
				RedisDb = redisdb
			}

			RedisPassword = RedisPasswordKey.String()
			idle, err := RedisMaxIdleKey.Int()
			if err == nil || idle > 10 {
				RedisMaxIdle = idle
			} else {
				RedisMaxIdle = 10
			}

		}

	}

	//EMAIL配置

	mailKey, err := OriginFile.GetSection("email")

	if err == nil {
		EmailSmtpHostKey, err := mailKey.GetKey("smtp_host")
		EmailSmtpPortKey, err := mailKey.GetKey("smtp_port")
		EmailSmtpPassKey, err := mailKey.GetKey("smtp_pass")
		EmailPop3HostKey, err := mailKey.GetKey("pop3_host")
		EmailPop3PortKey, err := mailKey.GetKey("pop3_port")
		EmailPop3PassKey, err := mailKey.GetKey("pop3_pass")
		EmailAddressKey, err := mailKey.GetKey("email_address")
		EmailNameKey, err := mailKey.GetKey("email_name")
		if err == nil {
			EmailSmtpHost = EmailSmtpHostKey.String()
			EmailSmtpPort = EmailSmtpPortKey.String()
			EmailSmtpPass = EmailSmtpPassKey.String()
			EmailPop3Host = EmailPop3HostKey.String()
			EmailPop3Port = EmailPop3PortKey.String()
			EmailPop3Pass = EmailPop3PassKey.String()
			EmailAddress = EmailAddressKey.String()
			EmailName = EmailNameKey.String()
		}
	}

	thriftKey, err := OriginFile.GetSection("thrift")
	if err == nil {
		thriftListenPortKey, err := thriftKey.GetKey("listen_port")
		ThriftSslServerKeyKey, err := thriftKey.GetKey("ssl_server_key")
		ThriftSslServerCrtKey, err := thriftKey.GetKey("ssl_server_crt")
		ThriftSslCaCrtKey, err := thriftKey.GetKey("ssl_ca_crt")
		ThriftSslClientCrtKey, err := thriftKey.GetKey("ssl_client_crt")
		if err == nil {
			ThriftListenPort = thriftListenPortKey.String()
			ThriftSslServerKey = appPath + string(os.PathSeparator) + "conf" + string(os.PathSeparator) + ThriftSslServerKeyKey.String()
			ThriftSslServerCrt = appPath + string(os.PathSeparator) + "conf" + string(os.PathSeparator) + ThriftSslServerCrtKey.String()
			ThriftSslCaCrt = appPath + string(os.PathSeparator) + "conf" + string(os.PathSeparator) + ThriftSslCaCrtKey.String()
			ThriftSslClientCrt = appPath + string(os.PathSeparator) + "conf" + string(os.PathSeparator) + ThriftSslClientCrtKey.String()
		}
	}
}

func LoadByParam(section, key string) (result string, err error) {
	if OriginFile == nil {
		err = fmt.Errorf("No config file readed!")
		return
	}
	pKey, err := OriginFile.GetSection("thrift")
	if err != nil {
		return
	}
	rKey, err := pKey.GetKey(key)
	if err != nil {
		return
	}
	return rKey.String(), nil
}

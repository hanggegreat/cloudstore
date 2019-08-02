package config

// HOST主机信息
var (
	HOST = "127.0.0.1"
)

// MySQL
var (
	// MysqlLink : mysql 路由
	MysqlLink = "root:root@tcp(" + HOST + ":3306)/icloud?charset=utf8"
)

// Redis
var (
	// RedisHost : redis 路由
	RedisHost = HOST + ":" + "6379"
	// RedisPass : redis auth
	RedisPass = ""
)

// Ceph
var (
	// CephAccessKey : 访问Key
	CephAccessKey = ""
	// CephSecretKey : 访问密钥
	CephSecretKey = ""
	// CephGWEndpoint : gateway地址
	CephGWEndpoint = "http://" + HOST + ":9080"
)

// kodo
var (
	// KodoBucket : bucket桶
	KodoBucket = ""
	// KodoEndpoint : kodo endpoint
	KodoEndpoint = ""
	// KodoAccesskey : kodo访问key
	KodoAccessKey = ""
	// KodoSecretKey : kodo访问key secret
	KodoSecretKey = ""
	// KodoDomain : kodo文件域名，可以自己绑定，也可以qiniu临时域名
	KodoDomain = ""
)

// RabbitMQ
var (
	// AsyncTransferEnable ： 异步上传文件
	AsyncTransferEnable = true
	// TransExchangeName : 用于文件transfer的交换机
	TransExchangeName = ""
	// TransKodoQueueName : oss转移队列名
	TransKodoQueueName = ""
	// TransKodoErrQueueName : oss转移失败后写入另一个队列的队列名
	TransKodoErrQueueName = ""
	// TransKodoRoutingKey : routingkey
	TransKodoRoutingKey = ""
	// RabbitURL : rabbitmq服务的入口url
	RabbitURL = "amqp://guest:guest@" + HOST + ":5672/"
)

// pwdSalt
var (
	// pwdSalt : 秘钥Salt
	pwdSalt = ""
	// tokenSalt : TokenSalt
	tokenSalt = "_tokenSalt"
)

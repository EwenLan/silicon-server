package globaldefine

const (
	BaseGoVersion   = "go version go1.18.4 windows/amd64"
	SoftwareVersion = "0.5.0"
	ProjectHome     = "https://github.com/EwenLan/silicon-server"
	Author          = "Ewen Lan"
	Email           = "ewen_lan@outlook.com"
	BuildDate       = "2022-07-16"
)

// 只有前缀与dynamicPrefix匹配的请求才通过代码实现处理请求，否则在根路径下搜索静态文件
const DynamicPrefix = "/api"

// RequestBuffLen 请求缓冲区
const RequestBuffLen = 1000

// ResponseBuffLen 响应缓冲区
const ResponseBuffLen = 1000

package utils

// 只有前缀与dynamicPrefix匹配的请求才通过代码实现处理请求，否则在根路径下搜索静态文件
const DynamicPrefix = "/api"

// RequestBuffLen 请求缓冲区
const RequestBuffLen = 1000

// ResponseBuffLen 响应缓冲区
const ResponseBuffLen = 1000

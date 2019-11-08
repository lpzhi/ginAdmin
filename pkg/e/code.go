package e
/*
200 OK - 对成功的 GET、PUT、PATCH 或 DELETE 操作进行响应。也可以被用在不创建新资源的 POST 操作上
201 Created - 对创建新资源的 POST 操作进行响应。应该带着指向新资源地址的 Location 头
202 Accepted - 服务器接受了请求，但是还未处理，响应中应该包含相应的指示信息，告诉客户端该去哪里查询关于本次请求的信息
204 No Content - 对不会返回响应体的成功请求进行响应（比如 DELETE 请求）
304 Not Modified - HTTP 缓存 header 生效的时候用
400 Bad Request - 请求异常，比如请求中的 body 无法解析
401 Unauthorized - 没有进行认证或者认证非法
403 Forbidden - 服务器已经理解请求，但是拒绝执行它
404 Not Found - 请求一个不存在的资源
405 Method Not Allowed - 所请求的 HTTP 方法不允许当前认证用户访问
410 Gone - 表示当前请求的资源不再可用。当调用老版本 API 的时候很有用
415 Unsupported Media Type - 如果请求中的内容类型是错误的
422 Unprocessable Entity - 用来表示校验错误
429 Too Many Requests - 由于请求频次达到上限而被拒绝访问
 */
const (
	SUCCESS = 200
	SUCCESS_CREATED = 201
	SUCCESS_ACCEPTED = 202
	SUCCESS_NO_CENTENT = 204
	NOT_MODIFIED = 304
	BAD_REQUEST = 400
	UNAUTHORIZED = 401
	FORBIDDEN = 403
	NOT_FOUND = 404
	NOT_ALLOWED = 405
	GONE = 410
	UNSUPPORTED = 415
	UNPROCESSABLE = 422
	MANY_REQUEST = 429

	ERROR = 500
	INVALID_PARAMS  = 400
	ERROR_EXIST_TAG = 10001
	ERROR_NOT_EXIST_TAG = 10002
	ERROR_NOT_EXIST_ARTICLE = 10003

	ERROR_AUIH_CHECK_TOKE_FAIL = 20001
	ERROR_AUTH_CHECK_TOKE_TIMEOUT = 20002
	ERROR_AUIH_TOKEN = 20003
	ERROR_AUTH = 20004
)

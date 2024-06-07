package http

type Status int

const (
	// 1xx informational response
	Continue           Status = 100
	SwitchingProtocols Status = 101
	Processing         Status = 102
	EarlyHints         Status = 103

	// 2xx successful
	OK                   Status = 200
	Created              Status = 201
	Accepted             Status = 202
	NonAuthoritativeInfo Status = 203
	NoContent            Status = 204
	ResetContent         Status = 205
	PartialContent       Status = 206
	MultiStatus          Status = 207
	AlreadyReported      Status = 208
	IMUsed               Status = 226

	// 3xx redirection
	MultipleChoices   Status = 300
	MovedPermanently  Status = 301
	Found             Status = 302
	SeeOther          Status = 303
	NotModified       Status = 304
	UseProxy          Status = 305
	SwitchProxy       Status = 306
	TemporaryRedirect Status = 307
	PermanentRedirect Status = 308

	// 4xx client errors
	BadRequest                   Status = 400
	Unauthorized                 Status = 401
	PaymentRequired              Status = 402
	Forbidden                    Status = 403
	NotFound                     Status = 404
	MethodNotAllowed             Status = 405
	NotAcceptable                Status = 406
	ProxyAuthRequired            Status = 407
	RequestTimeout               Status = 408
	Conflict                     Status = 409
	Gone                         Status = 410
	LengthRequired               Status = 411
	PreconditionFailed           Status = 412
	RequestEntityTooLarge        Status = 413
	RequestURITooLong            Status = 414
	UnsupportedMediaType         Status = 415
	RequestedRangeNotSatisfiable Status = 416
	ExpectationFailed            Status = 417
	Teapot                       Status = 418
	MisdirectedRequest           Status = 421
	UnprocessableEntity          Status = 422
	Locked                       Status = 423
	FailedDependency             Status = 424
	TooEarly                     Status = 425
	UpgradeRequired              Status = 426
	PreconditionRequired         Status = 428
	TooManyRequests              Status = 429
	RequestHeaderFieldsTooLarge  Status = 431
	UnavailableForLegalReasons   Status = 451

	// 5xx server errors
	InternalServerError           Status = 500
	NotImplemented                Status = 501
	BadGateway                    Status = 502
	ServiceUnavailable            Status = 503
	GatewayTimeout                Status = 504
	HTTPVersionNotSupported       Status = 505
	VariantAlsoNegotiates         Status = 506
	InsufficientStorage           Status = 507
	LoopDetected                  Status = 508
	NotExtended                   Status = 510
	NetworkAuthenticationRequired Status = 511
)

func (s Status) String() string {
	switch s {
	case Continue:
		return "Continue"
	case SwitchingProtocols:
		return "Switching Protocols"
	case Processing:
		return "Processing"
	case EarlyHints:
		return "Early Hints"
	case OK:
		return "OK"
	case Created:
		return "Created"
	case Accepted:
		return "Accepted"
	case NonAuthoritativeInfo:
		return "Non-Authoritative Information"
	case NoContent:
		return "No Content"
	case ResetContent:
		return "Reset Content"
	case PartialContent:
		return "Partial Content"
	case MultiStatus:
		return "Multi-Status"
	case AlreadyReported:
		return "Already Reported"
	case IMUsed:
		return "IM Used"
	case MultipleChoices:
		return "Multiple Choices"
	case MovedPermanently:
		return "Moved Permanently"
	case Found:
		return "Found"
	case SeeOther:
		return "See Other"
	case NotModified:
		return "Not Modified"
	case UseProxy:
		return "Use Proxy"
	case SwitchProxy:
		return "Switch Proxy"
	case TemporaryRedirect:
		return "Temporary Redirect"
	case PermanentRedirect:
		return "Permanent Redirect"
	case BadRequest:
		return "Bad Request"
	case Unauthorized:
		return "Unauthorized"
	case PaymentRequired:
		return "Payment Required"
	case Forbidden:
		return "Forbidden"
	case NotFound:
		return "Not Found"
	case MethodNotAllowed:
		return "Method Not Allowed"
	case NotAcceptable:
		return "Not Acceptable"
	case ProxyAuthRequired:
		return "Proxy Authentication Required"
	case RequestTimeout:
		return "Request Timeout"
	case Conflict:
		return "Conflict"
	case Gone:
		return "Gone"
	case LengthRequired:
		return "Length Required"
	case PreconditionFailed:
		return "Precondition Failed"
	case RequestEntityTooLarge:
		return "Request Entity Too Large"
	case RequestURITooLong:
		return "Request-URI Too Long"
	case UnsupportedMediaType:
		return "Unsupported Media Type"
	case RequestedRangeNotSatisfiable:
		return "Requested Range Not Satisfiable"
	case ExpectationFailed:
		return "Expectation Failed"
	case Teapot:
		return "I'm a teapot"
	case MisdirectedRequest:
		return "Misdirected Request"
	case UnprocessableEntity:
		return "Unprocessable Entity"
	case Locked:
		return "Locked"
	case FailedDependency:
		return "Failed Dependency"
	case TooEarly:
		return "Too Early"
	case UpgradeRequired:
		return "Upgrade Required"
	case PreconditionRequired:
		return "Precondition Required"
	case TooManyRequests:
		return "Too Many Requests"
	case RequestHeaderFieldsTooLarge:
		return "Request Header Fields Too Large"
	case UnavailableForLegalReasons:
		return "Unavailable For Legal Reasons"
	case InternalServerError:
		return "Internal Server Error"
	case NotImplemented:
		return "Not Implemented"
	case BadGateway:
		return "Bad Gateway"
	case ServiceUnavailable:
		return "Service Unavailable"
	case GatewayTimeout:
		return "Gateway Timeout"
	case HTTPVersionNotSupported:
		return "HTTP Version Not Supported"
	case VariantAlsoNegotiates:
		return "Variant Also Negotiates"
	case InsufficientStorage:
		return "Insufficient Storage"
	case LoopDetected:
		return "Loop Detected"
	case NotExtended:
		return "Not Extended"
	case NetworkAuthenticationRequired:
		return "Network Authentication Required"
	default:
		return "Unknown"
	}
}

func (s Status) Code() int {
	return int(s)
}

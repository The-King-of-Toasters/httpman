package info

// Map type for holding status code info
type codemap map[int]doc

type doc struct {
	Category    string
	Description string
	Message     string
}

/*
The contents of this map are derived from the
HTTP response status codes article on the Mozilla Developer Network (MDN).
HTTP response status codes[1] by Mozilla Contributors[2] is
licensed under CC-BY-SA 2.5[3].

This map was originally assembled by Jonathan Shobrook in 2018 for use in his
statcode[4] program. The only changes made were to transform it into a
native map literal.

[1] https://developer.mozilla.org/en-US/docs/Web/HTTP/Status
[2] https://developer.mozilla.org/en-US/docs/Web/HTTP/Status$history
[3] http://creativecommons.org/licenses/by-sa/2.5/
[4] https://github.com/shobrook/statcode
*/
var cm = codemap{
	100: doc{
		Category: "Informational Responses",
		Description: "This interim response indicates that everything so far is " +
			"OK and that the client should continue with the request or ignore it " +
			"if it is already finished.",
		Message: "Continue"},
	101: doc{
		Category: "Informational Responses",
		Description: "This code is sent in response to an Upgrade request header " +
			"by the client, and indicates the protocol the server is switching to.",
		Message: "Switching Protocol"},

	102: doc{
		Category: "Informational Responses",
		Description: "This code indicates that the server has received and is " +
			"processing the request, but no response is available yet.",
		Message: "Processing (WebDAV)"},

	200: doc{
		Category: "Successful Responses",
		Description: "The request has succeeded. The meaning of a success varies " +
			"depending on the HTTP method: GET: The resource has been fetched " +
			"and is transmitted in the message body. HEAD: The entity headers are " +
			"in the message body. PUT or POST: The resource describing the result " +
			"of the action is transmitted in the message body. TRACE: The message " +
			"body contains the request message as received by the server",
		Message: "OK"},

	201: doc{
		Category: "Successful Responses",
		Description: "The request has succeeded and a new resource has been created " +
			"as a result of it. This is typically the response sent after a POST " +
			"request, or after some PUT requests.",
		Message: "Created"},

	202: doc{
		Category: "Successful Responses",
		Description: "The request has been received but not yet acted upon. It is " +
			"non-committal, meaning that there is no way in HTTP to later send an " +
			"asynchronous response indicating the outcome of processing the request. " +
			"It is intended for cases where another process or server handles the " +
			"request, or for batch processing.",
		Message: "Accepted"},

	203: doc{
		Category: "Successful Responses",
		Description: "This response code means returned meta-information set is not" +
			"exact set as available from the origin server, but collected from a " +
			"local or a third party copy. Except this condition, 200 OK response " +
			"should be preferred instead of this response.",
		Message: "Non-Authoritative Information"},

	204: doc{
		Category: "Successful Responses",
		Description: "There is no content to send for this request, but the headers " +
			"may be useful. The user-agent may update its cached headers for this " +
			"resource with the new ones.",
		Message: "No Content"},

	205: doc{
		Category: "Successful Responses",
		Description: "This response code is sent after accomplishing request to " +
			"tell user agent reset document view which sent this request.",
		Message: "Reset Content"},

	206: doc{
		Category: "Successful Responses",
		Description: "This response code is used because of range header sent " +
			"by the client to separate download into multiple streams.",
		Message: "Partial Content"},

	207: doc{
		Category: "Successful Responses",
		Description: "A Multi-Status response conveys information about multiple " +
			"resources in situations where multiple status codes might be appropriate.",
		Message: "Multi-Status (WebDav)"},

	208: doc{
		Category: "Successful Responses",
		Description: "Used inside a DAV: propstat response element to avoid enumerating " +
			"the internal members of multiple bindings to the same collection repeatedly.",
		Message: "Multi-Status (WebDav)"},

	226: doc{
		Category: "Successful Responses",
		Description: "The server has fulfilled a GET request for the resource, and the " +
			"response is a representation of the result of one or more " +
			"instance-manipulations applied to the current instance.",
		Message: "IM Used (HTTP Delta encoding)"},

	300: doc{
		Category: "Redirection Messages",
		Description: "The request has more than one possible response. The user-agent " +
			"or user should choose one of them. There is no standardized way of choosing" +
			"one of the responses.",
		Message: "Multiple Choice"},

	301: doc{
		Category:    "Redirection Messages",
		Description: "This response code means that the URI of the requested resource has been changed. Probably, the new URI would be given in the response.",
		Message:     "Moved Permanently"},

	302: doc{
		Category:    "Redirection Messages",
		Description: "This response code means that the URI of requested resource has been changed temporarily. New changes in the URI might be made in the future. Therefore, this same URI should be used by the client in future requests.",
		Message:     "Found"},

	303: doc{
		Category:    "Redirection Messages",
		Description: "The server sent this response to direct the client to get the requested resource at another URI with a GET request.",
		Message:     "See Other"},

	304: doc{
		Category:    "Redirection Messages",
		Description: "This is used for caching purposes. It tells the client that the response has not been modified, so the client can continue to use the same cached version of the response.",
		Message:     "Not Modified"},

	305: doc{
		Category:    "Redirection Messages",
		Description: "Was defined in a previous version of the HTTP specification to indicate that a requested response must be accessed by a proxy. It has been deprecated due to security concerns regarding in-band configuration of a proxy.",
		Message:     "Use Proxy"},

	306: doc{
		Category:    "Redirection Messages",
		Description: "This response code is no longer used, it is just reserved currently. It was used in a previous version of the HTTP 1.1 specification.",
		Message:     "Unused"},

	307: doc{
		Category:    "Redirection Messages",
		Description: "The server sends this response to direct the client to get the requested resource at another URI with same method that was used in the prior request. This has the same semantics as the 302 Found HTTP response code, with the exception that the user agent must not change the HTTP method used: If a POST was used in the first request, a POST must be used in the second request.",
		Message:     "Temporary Redirect"},

	308: doc{
		Category:    "Redirection Messages",
		Description: "This means that the resource is now permanently located at another URI, specified by the Location: HTTP Response header. This has the same semantics as the 301 Moved Permanently HTTP response code, with the exception that the user agent must not change the HTTP method used: If a POST was used in the first request, a POST must be used in the second request.",
		Message:     "Permanent Redirect"},

	400: doc{
		Category:    "Client Error Responses",
		Description: "This response means that server could not understand the request due to invalid syntax.",
		Message:     "Bad Request"},

	401: doc{
		Category:    "Client Error Responses",
		Description: "Although the HTTP standard specifies 'unauthorized', semantically this response means 'unauthenticated'. That is, the client must authenticate itself to get the requested response.",
		Message:     "Unauthorized"},

	402: doc{
		Category:    "Client Error Responses",
		Description: "This response code is reserved for future use. Initial aim for creating this code was using it for digital payment systems however this is not used currently.",
		Message:     "Payment Required"},

	403: doc{
		Category:    "Client Error Responses",
		Description: "The client does not have access rights to the content, i.e. they are unauthorized, so server is rejecting to give proper response. Unlike 401, the client's identity is known to the server.",
		Message:     "Forbidden"},

	404: doc{
		Category:    "Client Error Responses",
		Description: "The server can not find requested resource. In the browser, this means the URL is not recognized. In an API, this can also mean that the endpoint is valid but the resource itself does not exist. Servers may also send this response instead of 403 to hide the existence of a resource from an unauthorized client. This response code is probably the most famous one due to its frequent occurrence on the web.",
		Message:     "Not Found"},

	405: doc{
		Category:    "Client Error Responses",
		Description: "The request method is known by the server but has been disabled and cannot be used. For example, an API may forbid DELETE-ing a resource. The two mandatory methods, GET and HEAD, must never be disabled and should not return this error code.",
		Message:     "Method Not Allowed"},

	406: doc{
		Category:    "Client Error Responses",
		Description: "This response is sent when the web server, after performing server-driven content negotiation, doesn't find any content following the criteria given by the user agent.",
		Message:     "Not Acceptable"},

	407: doc{
		Category:    "Client Error Responses",
		Description: "This is similar to 401 but authentication is needed to be done by a proxy.",
		Message:     "Proxy Authentication Required"},

	408: doc{
		Category:    "Client Error Responses",
		Description: "This response is sent on an idle connection by some servers, even without any previous request by the client. It means that the server would like to shut down this unused connection. This response is used much more since some browsers, like Chrome, Firefox 27+, or IE9, use HTTP pre-connection mechanisms to speed up surfing. Also note that some servers merely shut down the connection without sending this message.",
		Message:     "Request Timeout"},

	409: doc{
		Category:    "Client Error Responses",
		Description: "This response is sent when a request conflicts with the current state of the server.",
		Message:     "Conflict"},

	410: doc{
		Category:    "Client Error Responses",
		Description: "This response would be sent when the requested content has been permanently deleted from server, with no forwarding address. Clients are expected to remove their caches and links to the resource. The HTTP specification intends this status code to be used for 'limited-time, promotional services'. APIs should not feel compelled to indicate resources that have been deleted with this status code.",
		Message:     "Gone"},

	411: doc{
		Category:    "Client Error Responses",
		Description: "Server rejected the request because the Content-Length header field is not defined and the server requires it.",
		Message:     "Length Required"},

	412: doc{
		Category:    "Client Error Responses",
		Description: "The client has indicated preconditions in its headers which the server does not meet.",
		Message:     "Precondition Failed"},

	413: doc{
		Category:    "Client Error Responses",
		Description: "Request entity is larger than limits defined by server; the server might close the connection or return an Retry-After header field.",
		Message:     "Payload Too Large"},

	414: doc{
		Category:    "Client Error Responses",
		Description: "Request entity is larger than limits defined by server; the server might close the connection or return an Retry-After header field.",
		Message:     "URI Too Long"},

	415: doc{
		Category:    "Client Error Responses",
		Description: "The media format of the requested data is not supported by the server, so the server is rejecting the request.",
		Message:     "Unsupported Media Type"},

	416: doc{
		Category:    "Client Error Responses",
		Description: "The range specified by the Range header field in the request can't be fulfilled; it's possible that the range is outside the size of the target URI's data.",
		Message:     "Requested Range Not Satisfiable"},

	417: doc{
		Category:    "Client Error Responses",
		Description: "This response code means the expectation indicated by the Expect request header field can't be met by the server.",
		Message:     "Expectation Failed"},

	418: doc{
		Category:    "Client Error Responses",
		Description: "The server refuses the attempt to brew coffee with a teapot.",
		Message:     "I'm a teapot"},

	421: doc{
		Category:    "Client Error Responses",
		Description: "The request was directed at a server that is not able to produce a response. This can be sent by a server that is not configured to produce responses for the combination of scheme and authority that are included in the request URI.",
		Message:     "Misdirected Request"},

	422: doc{
		Category:    "Client Error Responses",
		Description: "The request was well-formed but was unable to be followed due to semantic errors.",
		Message:     "Unprocessable Entity (WebDAV)"},

	423: doc{
		Category:    "Client Error Responses",
		Description: "The resource that is being accessed is locked.",
		Message:     "Locked (WebDAV)"},

	424: doc{
		Category:    "Client Error Responses",
		Description: "The request failed due to failure of a previous request.",
		Message:     "Failed Dependency (WebDAV)"},

	426: doc{
		Category:    "Client Error Responses",
		Description: "The server refuses to perform the request using the current protocol but might be willing to do so after the client upgrades to a different protocol. The server sends an Upgrade header in a 426 response to indicate the required protocol(s).",
		Message:     "Upgrade Required"},

	428: doc{
		Category:    "Client Error Responses",
		Description: "The origin server requires the request to be conditional. Intended to prevent the 'lost update' problem, where a client GETs a resource's state, modifies it, and PUTs it back to the server, when meanwhile a third party has modified the state on the server, leading to a conflict.",
		Message:     "Precondition Required"},

	429: doc{
		Category:    "Client Error Responses",
		Description: "The user has sent too many requests in a given amount of time ('rate limiting').",
		Message:     "Too Many Requests"},

	431: doc{
		Category:    "Client Error Responses",
		Description: "The server is unwilling to process the request because its header fields are too large. The request MAY be resubmitted after reducing the size of the request header fields.",
		Message:     "Request Header Fields Too Large"},

	451: doc{
		Category:    "Client Error Responses",
		Description: "The user requests an illegal resource, such as a web page censored by a government.",
		Message:     "Unavailable For Legal Reasons"},

	500: doc{
		Category:    "Server Error Responses",
		Description: "The server has encountered a situation it doesn't know how to handle.",
		Message:     "Internal Server Error"},

	501: doc{
		Category:    "Server Error Responses",
		Description: "The request method is not supported by the server and cannot be handled. The only methods that servers are required to support (and therefore that must not return this code) are GET and HEAD.",
		Message:     "Not Implemented"},

	502: doc{
		Category:    "Server Error Responses",
		Description: "This error response means that the server, while working as a gateway to get a response needed to handle the request, got an invalid response.",
		Message:     "Bad Gateway"},

	503: doc{
		Category:    "Server Error Responses",
		Description: "The server is not ready to handle the request. Common causes are a server that is down for maintenance or that is overloaded. Note that together with this response, a user-friendly page explaining the problem should be sent. This responses should be used for temporary conditions and the Retry-After: HTTP header should, if possible, contain the estimated time before the recovery of the service. The webmaster must also take care about the caching-related headers that are sent along with this response, as these temporary condition responses should usually not be cached.",
		Message:     "Service Unavailable"},

	504: doc{
		Category:    "Server Error Responses",
		Description: "This error response is given when the server is acting as a gateway and cannot get a response in time.",
		Message:     "Gateway Timeout"},

	505: doc{
		Category:    "Server Error Responses",
		Description: "The HTTP version used in the request is not supported by the server.",
		Message:     "HTTP Version Not Supported"},

	506: doc{
		Category:    "Server Error Responses",
		Description: "The server has an internal configuration error: transparent content negotiation for the request results in a circular reference.",
		Message:     "Variant Also Negotiates"},

	507: doc{
		Category:    "Server Error Responses",
		Description: "The server has an internal configuration error: the chosen variant resource is configured to engage in transparent content negotiation itself, and is therefore not a proper end point in the negotiation process.",
		Message:     "Insufficient Storage"},

	508: doc{
		Category:    "Server Error Responses",
		Description: "The server detected an infinite loop while processing the request.",
		Message:     "Loop Detected (WebDAV)"},

	510: doc{
		Category:    "Server Error Responses",
		Description: "Further extensions to the request are required for the server to fulfill it.",
		Message:     "Not Extended"},

	511: doc{
		Category:    "Server Error Responses",
		Description: "The 511 status code indicates that the client needs to authenticate to gain network access.",
		Message:     "Network Authentication Required"},
}

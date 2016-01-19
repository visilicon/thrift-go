struct Request {
1: required i64 UserId
2: required string Payload
}

struct Response {
1: required i32 Errno
2: required string Errmsg
}

service Puller {
	Response pull(1:Request request),
}

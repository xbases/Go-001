
syntax = "proto3";

package welcome;

// 用于生成指定语言go的包名称
option go_package = "welcome/api";

service Hello {
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}


message HelloRequest {
  int64 id = 1;
}


message HelloReply {
  string message = 1;
}
protoc -I proto/ proto/searchparam2.proto --go_out=plugins=grpc:proto
protoc -I proto/ proto/alimama.proto --go_out=plugins=grpc:proto
protoc -I proto/ proto/steepandcheap.proto --go_out=plugins=grpc:proto
protoc -I proto/ proto/mountainsteals.proto --go_out=plugins=grpc:proto
protoc -I proto/ proto/tmall.proto --go_out=plugins=grpc:proto
protoc -I proto/ proto/taobao.proto --go_out=plugins=grpc:proto
protoc -I proto/ proto/manhuadb.proto --go_out=plugins=grpc:proto
protoc -I proto/ proto/douban.proto --go_out=plugins=grpc:proto
protoc -I proto/ proto/jrj.proto --go_out=plugins=grpc:proto
protoc -I proto/ proto/telegraph.proto --go_out=plugins=grpc:proto
protoc -I proto/ proto/oabt.proto --go_out=plugins=grpc:proto
protoc -I proto/ proto/result.proto --go_out=plugins=grpc:proto
protoc -I dbpb/ dbpb/db.proto --go_out=plugins=grpc:dbpb
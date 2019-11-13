protoc -I proto/ proto/alimama.proto --go_out=plugins=grpc:proto
protoc -I proto/ proto/steepandcheap.proto --go_out=plugins=grpc:proto
protoc -I proto/ proto/mountainsteals.proto --go_out=plugins=grpc:proto
protoc -I proto/ proto/tmall.proto --go_out=plugins=grpc:proto
protoc -I proto/ proto/result.proto --go_out=plugins=grpc:proto
protoc -I dbpb/ dbpb/db.proto --go_out=plugins=grpc:dbpb
protoc -I proto/ proto/result.proto --go_out=plugins=grpc:proto
protoc -I dbpb/ dbpb/db.proto --go_out=plugins=grpc:dbpb
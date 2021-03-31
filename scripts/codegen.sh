find user_web/proto -type f -name "*.pb.go" -delete
# common
#protoc --proto_path=api/protobuf/hashtag/common --go_out=plugins=grpc:hashtag common.proto
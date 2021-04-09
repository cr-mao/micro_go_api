find user_web/proto -type f -name "*.pb.go" -delete
find order_web/proto -type f -name "*.pb.go" -delete
find goods_web/proto -type f -name "*.pb.go" -delete
# common
#protoc --proto_path=api/protobuf/hashtag/common --go_out=plugins=grpc:hashtag common.proto
protoc --proto_path=user_web/proto --go_out=plugins=grpc:user_web/proto user.proto
protoc --proto_path=goods_web/proto --go_out=plugins=grpc:goods_web/proto goods.proto
protoc --proto_path=order_web/proto --go_out=plugins=grpc:order_web/proto order.proto
protoc --proto_path=order_web/proto --go_out=plugins=grpc:order_web/proto inventory.proto
protoc --proto_path=order_web/proto --go_out=plugins=grpc:order_web/proto goods.proto

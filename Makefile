.PHONY: proto_products
proto_products: ## Регенерация протоколов pb
	if not exist grpc\pb\products mkdir grpc\pb\products
	protoc --go_out=grpc/pb/products --go-grpc_out=grpc/pb/products --proto_path=grpc/proto grpc/proto/products.proto

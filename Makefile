proto_products: ## Регенерация протоколов pb
	if not exist grpc\pb\products mkdir grpc\pb\products
	protoc --go_out=. --go-grpc_out=. --proto_path=grpc/proto grpc/proto/products.proto
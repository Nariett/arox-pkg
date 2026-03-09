proto_products: ## Регенерация протоколов pb и моков
	if not exist grpc\pb\products mkdir grpc\pb\products
	protoc --go_out=. --go-grpc_out=. --proto_path=grpc/proto grpc/proto/products.proto
	if not exist internal\mock\products mkdir internal\mock\products
	mockgen -source=grpc\pb\products\products_grpc.pb.go -destination=grpc\pb\mock\products\client.go -package=mockproducts ProductsServiceClient

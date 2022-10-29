.PHONY: protos

protos:
	# protoc -I protos/ protos/currency.proto --go_out=plugins=grpc:protos/currency
	protoc --go_out=protos/currency/ -I protos/ protos/currency.proto
	protoc --go-grpc_out=protos/currency/ -I protos/ protos/currency.proto

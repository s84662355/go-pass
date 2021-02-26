proto:
	protoc --go_out=. convention/protobuf/*.proto
	rm -rf lib/protocol/*
	mv convention/protobuf/*.go lib/protocol/

WORKING_DIR=$(shell echo $${PWD\#\#*/} )

build:
	@rm -rf *.pb.go messages
	@protoc --go_out=. protocol/messages/*.proto
	@mkdir -p messages
	@mv protocol/messages/Request.pb.go messages/Request.go
	@mv protocol/messages/Response.pb.go messages/Response.go

.PHONY: all test test-unit clean
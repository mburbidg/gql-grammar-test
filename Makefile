.DEFAULT_GOAL := generate

generate:
	go generate ./...
	mv parser/gen/grammar/* parser/gen
	rm -fr parser/gen/grammar
.PHONY:generate

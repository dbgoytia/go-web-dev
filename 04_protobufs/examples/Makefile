.PHONY: clean go_mod

export SRC_DIR = ${HOME}/go/src/github.com/go-web-dev/04_protobufs/examples
export DST_DIR = ${SRC_DIR}

clean:
	rm -rf todo.pb.go go.mod go.sum

go_mod:
	go mod init github.com/go-web-dev/04_protobufs/examples
	go mod tidy

go_protoc:
	protoc -I=${SRC_DIR} --go_out=${DST_DIR} ${SRC_DIR}/todo.proto;

list_task_gotest:
	go test list_task.go list_task_test.go
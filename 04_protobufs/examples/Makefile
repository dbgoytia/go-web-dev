.PHONY: clean go_mod

export SRC_DIR = ${HOME}/go/src/github.com/go-web-dev/04_protobufs/examples
export DST_DIR = ${SRC_DIR}

clean:
	rm -rf todo.pb.go go.mod go.sum

mod:
	go mod init github.com/go-web-dev/04_protobufs/examples
	go mod tidy

protoc:
	protoc -I=${SRC_DIR} --go_out=${DST_DIR} ${SRC_DIR}/todo.proto;

test:
	go test list_task.go list_task_test.go
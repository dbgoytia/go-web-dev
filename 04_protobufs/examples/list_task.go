package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	pb "github.com/protocolbuffers/protobuf/examples/tutorial"
)

func writeTask(w io.Writer, t *pb.Person){
	fmt.Fprintln(w, "Task ID:", t.Id)
	fmt.Fprintln(w, "  Name:", t.Name)
	fmt.Fprintln(w, "  Status:", t.Status)
	for _, td :=  range t.TaskDetails {
		switch td.Type {
		case pb.Task_WORK:
			fmt.Fprintln(w, "  Work task: ")
		}
		case pb.Task_OSS:
			fmt.Fprintln(w, "  OSS task: ")
		}
		case pb.Task_TRAINING:
			fmt.Fprintln(w, "  TRAINING task: ")
		}
		fmt.Fprintln(w, td.Description)
	}
}
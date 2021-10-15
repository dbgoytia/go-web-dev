package main

import (
	"fmt"
	"io"

	pb "github.com/go-web-dev/04_protobufs/examples/tutorial"
)

func writeTask(w io.Writer, t *pb.Task){
	fmt.Fprintln(w, "Task ID:", t.Id)
	fmt.Fprintln(w, "  Name:", t.Title)
	fmt.Fprintln(w, "  Status:", t.Status)
	for _, td :=  range t.Tasks{
		switch td.Type {
		case pb.Task_WORK:
			fmt.Fprint(w, "  Work task: ")
		case pb.Task_OSS:
			fmt.Fprint(w, "  OSS task: ")
		case pb.Task_TRAINING:
			fmt.Fprint(w, "  TRAINING task: ")
		}
		fmt.Fprintln(w, td.Description)
	}
}
package main

import (
	"bytes"
	"testing"

	pb "github.com/go-web-dev/04_protobufs/examples/tutorial"
)

func TestWriteTaskWritesTask(t *testing.T){

	buf := new(bytes.Buffer)

	// [START] populate proto
	ta := pb.Task{
		Id:		0001,
		Title: 	"Proto training",
		Status: "backlog",
		
		Tasks: []*pb.Task_TaskDetails {
			{
				Description: "Create a sample unit test for todo app.",
				Type: pb.Task_TRAINING,
			},
		},
	}
	// [END] populate proto

	writeTask(buf, &ta)
	got := buf.String()
	want := `Task ID: 001
Name: Proto training
Status: Backlog
Training task: Create a sample unit test for todo app.
`
	if got != want {
		t.Errorf("writeTask(%s) =>\n\t%q, want %q", ta.String(), got, want)
	}

}


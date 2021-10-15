package main

import (
	"bytes"
	"strings"
	"testing"

	pb "github.com/go-web-dev/04_protobufs/tutorialpb"
)

func TestWriteTaskWritesTask(t *testing.T){

	buf := new(bytes.Buffer)

	// [START] populate proto
	t = pb.Task{
		Id: 0001,
		Title: "Proto training",
		Status: "backlog",
		
		Tasks: []*pb.Task_TaskDetails {
			{
				Description: "Create a sample unit test for todo app."
				Type: pb.Task_TRAINING
			}
		}
	}
	// [END] populate proto

	writeTask(buf, &t)
	got := buf.String()
	want := `Task ID: 001
Name: Proto training
Status: Backlog
Training task: Create a sample unit test for todo app.
`
	if got != want {
		t.Errorf("writeTask(%s) =>\n\t%q, want %q", p.String(), got, want)
	}

}


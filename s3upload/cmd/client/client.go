package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"path/filepath"
	pb "yir/s3upload/api"

	"github.com/google/uuid"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50055", grpc.WithInsecure())
	if err != nil {
		panic(fmt.Errorf("grpc conn error: %w", err))
	}

	client := pb.NewS3UploadClient(conn)

	hdlr := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("open stream: %v", err)))
		return

		
		stream, err := client.Upload(context.Background())
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("open stream: %v", err)))
			return
		}

		fileID, _ := uuid.NewRandom()

		total := 0

		imgbuf := [1024 * 1024]byte{}
		for {
			n, err := r.Body.Read(imgbuf[:])
			total += n
			if n != 0 {
				err := stream.Send(&pb.UploadFile{
					Path: filepath.Join(fileID.String(), fileID.String()),
					File: imgbuf[:n],
				})
				if err != nil {
					w.WriteHeader(500)
					w.Write([]byte(fmt.Sprintf("send to stream: %v", err)))
					return
				}
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				w.WriteHeader(500)
				w.Write([]byte(fmt.Sprintf("read from body: %v", err)))
				return
			}
		}

		_, err = stream.CloseAndRecv()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("close stream: %v", err)))
			return
		}

		w.WriteHeader(200)
		w.Write([]byte(fmt.Sprintf("%s, %d", fileID.String(), total)))
	}

	http.HandleFunc("/load", hdlr)
	http.ListenAndServe("localhost:8090", nil)
}

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	pb "yir/s3upload/api"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50055", grpc.WithInsecure())
	if err != nil {
		panic(fmt.Errorf("grpc conn error: %w", err))
	}

	client := pb.NewS3UploadClient(conn)

	hdlr := func(w http.ResponseWriter, r *http.Request) {
		stream, err := client.UploadAndSplitUziFile(context.Background())
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("open stream: %v", err)))
			return
		}

		imgbuf := [1024 * 1024]byte{}
		for {
			n, err := r.Body.Read(imgbuf[:])
			if n != 0 {
				err := stream.Send(&pb.ImageStream{File: imgbuf[:n]})
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

		resp, err := stream.CloseAndRecv()
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("close stream: %v", err)))
			return
		}

		respJSON, err := json.Marshal(map[string]any{
			"uzi_id":    resp.UziId,
			"images_id": resp.ImagesIds,
		})
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("json marshal response: %v", err)))
			return
		}
		w.WriteHeader(200)
		w.Write(respJSON)
	}

	http.HandleFunc("/split", hdlr)
	http.ListenAndServe("localhost:8080", nil)
}

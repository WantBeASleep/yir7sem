package main

import (
	"bytes"
	"context"
	"encoding/json"
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

	loadHdl := func(w http.ResponseWriter, r *http.Request) {
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
			n, httpErr := r.Body.Read(imgbuf[:])
			total += n
			fmt.Println("total: ", total, "httpErr: ", httpErr)
			if n != 0 {
				msg := &pb.UploadFile{
					Path: filepath.Join(fileID.String(), fileID.String()),
					File: imgbuf[:n],
				}
				fmt.Printf("SEND --> msg.Path = %s, len(msg.File) = %d\n", msg.Path, len(msg.File))
				err := stream.Send(msg)
				if err != nil {
					w.WriteHeader(500)
					w.Write([]byte(fmt.Sprintf("send to stream: %v", err)))
					return
				}
			}
			if httpErr != nil {
				if httpErr == io.EOF {
					break
				}
				w.WriteHeader(500)
				w.Write([]byte(fmt.Sprintf("read from http body: %v", err)))
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

	pingHdl := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("2000 ok"))
	}

	getHdl := func(w http.ResponseWriter, r *http.Request) {
		path := &struct{ Path string }{}
		json.NewDecoder(r.Body).Decode(&path)

		stream, err := client.Get(context.Background(), &pb.GetRequest{Path: path.Path})
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("open stream: %v", err)))
			return
		}

		file := bytes.Buffer{}
		for {
			rec, err := stream.Recv()
			if rec != nil {
				file.Write(rec.FileContent)
			}
			if err != nil {
				if err == io.EOF {
					break
				}
				w.WriteHeader(500)
				w.Write([]byte(fmt.Sprintf("receive from stream: %v", err)))
				return
			}
		}

		w.Header().Add("Content-Type", "image/png")
		w.Write(file.Bytes())
	}

	http.HandleFunc("/load", loadHdl)
	http.HandleFunc("/get", getHdl)
	http.HandleFunc("/ping", pingHdl)
	http.ListenAndServe("localhost:8080", nil)
}

package utils

import (
	"fmt"
	"io"
	pb "yir/s3upload/api"
)

// NOT THREAD SAFE!
// не хочется просто плюнуть сюда Mutex'ом
// можно написать крутую кастом обертку, но работа на недельку где то
type UploadGRPCReader struct {
	stream pb.S3Upload_UploadServer

	path       string
	cntReadOps int
	cache      []byte
	eof        bool
}

func (r *UploadGRPCReader) recv() error {
	r.cntReadOps++
	msg, err := r.stream.Recv()
	if err != nil && err != io.EOF {
		return fmt.Errorf("read grpc stream: %w", err)
	}

	if msg != nil {
		if r.cntReadOps == 0 {
			r.path = msg.Path
		}
		r.cache = append(r.cache, msg.File...)
	}

	if err == io.EOF || msg == nil {
		r.eof = true
		return io.EOF
	}
	return nil
}

// прочитает поток если до этого не было чтения
func (r *UploadGRPCReader) GetPath() (string, error) {
	if r.cntReadOps == 0 {
		if err := r.recv(); err != nil && err != io.EOF {
			return "", fmt.Errorf("receive grpc msg: %w", err)
		}
	}

	if r.path == "" {
		return "", fmt.Errorf("empty path not supported")
	}
	return r.path, nil
}

func (r *UploadGRPCReader) Read(p []byte) (int, error) {
	for !r.eof && len(r.cache) < len(p) {
		if err := r.recv(); err != nil {
			if err == io.EOF {
				break
			}
			return 0, fmt.Errorf("receive grpc msg: %w", err)
		}
	}

	copyBytes := copy(p, r.cache)
	r.cache = r.cache[copyBytes:]

	if r.eof && copyBytes == 0 {
		return 0, io.EOF
	}
	return copyBytes, nil
}

// написать на дженериках универсальный - ебани на неделю, пока что так
func NewUploadGRPCReader(stream pb.S3Upload_UploadServer) *UploadGRPCReader {
	return &UploadGRPCReader{
		stream: stream,
	}
}

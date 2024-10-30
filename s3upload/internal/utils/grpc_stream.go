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

	meta       *pb.FileMeta
	cntReadOps int
	file       []byte
	eof        bool
}

func (r *UploadGRPCReader) recv() error {
	r.cntReadOps++
	msg, err := r.stream.Recv()
	if err != nil && err != io.EOF {
		return fmt.Errorf("read grpc stream: %w", err)
	}

	if msg != nil {
		// получение Path при первом чтении
		if r.cntReadOps == 1 {
			if msg.FileMeta == nil {
				return fmt.Errorf("required file meta")
			}
			r.meta = msg.FileMeta
		}
		r.file = append(r.file, msg.File...)
	}

	if err == io.EOF {
		r.eof = true
		return io.EOF
	}
	return nil
}

// прочитает поток если до этого не было чтения
func (r *UploadGRPCReader) GetMeta() (*pb.FileMeta, error) {
	if r.cntReadOps == 0 {
		if err := r.recv(); err != nil && err != io.EOF {
			return nil, fmt.Errorf("receive grpc msg: %w", err)
		}
	}

	if r.meta == nil {
		return nil, fmt.Errorf("empty path not supported")
	}
	return r.meta, nil
}

func (r *UploadGRPCReader) Read(p []byte) (int, error) {
	for !r.eof && len(r.file) < len(p) {
		if err := r.recv(); err != nil {
			if err == io.EOF {
				break
			}
			return 0, fmt.Errorf("receive grpc msg: %w", err)
		}
	}

	copyBytes := copy(p, r.file)
	r.file = r.file[copyBytes:]

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

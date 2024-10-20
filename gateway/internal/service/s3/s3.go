package s3

import (
	"yir/gateway/internal/pb/s3pb"

	"google.golang.org/grpc"
)


type S3 struct {
	conn *grpc.ClientConn
	client s3pb.S3UploadClient
}

func (s *S3) Connect(addr, port string, opts []grpc.DialOption) (err error) {
	s.conn, err = grpc.NewClient(addr+port, opts...)
	if err != nil{
		return err
	}
	s.client = s3pb.NewS3UploadClient(s.conn)
	return nil
}

func (s *S3) Close(){
	s.conn.Close()
}

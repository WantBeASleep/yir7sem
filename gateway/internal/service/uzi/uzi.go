package uzi

import (
	"yir/gateway/internal/pb/uzipb"

	"google.golang.org/grpc"
)

type Uzi struct {
	conn   *grpc.ClientConn
	client uzipb.UziAPIClient
}

func (u *Uzi) Connect(addr, port string, opts []grpc.DialOption) (err error) {
	u.conn, err = grpc.NewClient(addr+port, opts...)
	if err != nil {
		return err
	}
	u.client = uzipb.NewUziAPIClient(u.conn)
	return nil
}

func (u *Uzi) Close() {
	u.conn.Close()
}

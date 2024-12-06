package main

import (
	"yir/pkg/brokerlib"
	pb "yir/uzi/internal/generated/broker/consume/uziupload"

	"google.golang.org/protobuf/proto"
)

func main() {
	b, _ := brokerlib.NewProducer([]string{"localhost:19092"})
	msg := &pb.UziUpload{UziId: "b0b27250-a958-468b-bd44-3355ab2b3f4d"}
	p, err := proto.Marshal(msg)
	if err != nil {
		panic("marshal")
	}
	b.Send("uziupload", "", p)
}

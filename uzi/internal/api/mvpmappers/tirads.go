package mvpmappers

import (
	kafka "yir/uzi/api/broker"
	pb "yir/uzi/api/grpcapi"
	"yir/uzi/internal/entity"
)

func PBTiradsToTirads(tirads *pb.Tirads) *entity.Tirads {
	return &entity.Tirads{
		Tirads23: tirads.Tirads_23,
		Tirads4:  tirads.Tirads_4,
		Tirads5:  tirads.Tirads_5,
	}
}

func KafkaTiradsToTirads(tirads *kafka.Tirads) *entity.Tirads {
	return &entity.Tirads{
		Tirads23: tirads.Tirads_23,
		Tirads4:  tirads.Tirads_4,
		Tirads5:  tirads.Tirads_5,
	}
}

func TiradsToPBTirads(tirads *entity.Tirads) *pb.Tirads {
	return &pb.Tirads{
		Tirads_23: tirads.Tirads23,
		Tirads_4:  tirads.Tirads4,
		Tirads_5:  tirads.Tirads5,
	}
}

package mvpmappers

import (
	pb "yir/uzi/api/grpcapi"
	"yir/uzi/internal/usecases/dto"
)

func PBFormationReqToDTOFormation(formation *pb.FormationRequest) *dto.Formation {
	return &dto.Formation{
		Ai:     formation.Ai,
		Tirads: PBTiradsToTirads(formation.Tirads),
	}
}

func PBFormationsReqToDTOFormations(formations []*pb.FormationRequest) []dto.Formation {
	dtoFormations := make([]dto.Formation, 0, len(formations))
	for _, formation := range formations {
		dtoFormations = append(dtoFormations, *PBFormationReqToDTOFormation(formation))
	}

	return dtoFormations
}

func DTOFormationToPBFormationResp(formation *dto.Formation) *pb.FormationResponse {
	return &pb.FormationResponse{
		Id:     formation.Id.String(),
		Ai:     formation.Ai,
		Tirads: TiradsToPBTirads(formation.Tirads),
	}
}

func DTOFormationsToPBFormationsResp(formations []dto.Formation) []*pb.FormationResponse {
	PBFormations := make([]*pb.FormationResponse, 0, len(formations))
	for _, v := range formations {
		PBFormations = append(PBFormations, DTOFormationToPBFormationResp(&v))
	}

	return PBFormations
}

func PBCreateFormationWithSegmentsReqToDTOFormationWithSegments(formationWithSegments *pb.CreateFormationWithSegmentsRequest) *dto.FormationWithSegments {
	return &dto.FormationWithSegments{
		Formation: PBFormationReqToDTOFormation(formationWithSegments.Formation),
		Segments:  PBSegmentsReqToDTOSegments(formationWithSegments.Segments),
	}
}

func DTOFormationWithSegmentsToPBFormationWithSegments(formationWithSegments *dto.FormationWithSegments) *pb.FormationWithSegments {
	return &pb.FormationWithSegments{
		Formation: DTOFormationToPBFormationResp(formationWithSegments.Formation),
		Segments: DTOSegmentsToPBSegmentsResp(formationWithSegments.Segments),
	}
}
package mvpmappers

import (
	pb "yir/uzi/api/grpcapi"
	"yir/uzi/internal/usecases/dto"
)

func DTOReportToPBReport(report *dto.Report) *pb.Report {
	return &pb.Report{
		Uzi:        UziToPBUziResp(report.Uzi),
		Images:     ImagesToPBImagesResp(report.Images),
		Formations: DTOFormationsToPBFormationsResp(report.Formations),
		Segments:   DTOSegmentsToPBSegmentsResp(report.Segments),
	}
}

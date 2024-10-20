package s3

import "yir/gateway/internal/pb/s3gen"

type S3Service interface {
	UploadAndSplitUziFile(s3gen.S3Upload_UploadAndSplitUziFileServer) s3gen.UploadUziFileResponse
	GetByPathImage(s3gen.GetImageRequest) s3gen.ImageStream
}

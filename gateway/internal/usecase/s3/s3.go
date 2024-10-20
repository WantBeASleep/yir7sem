package s3

type S3UseCase struct {
	Service *S3Service
}

func New(service *S3Service) *S3UseCase {
	return &S3UseCase{
		Service: service,
	}
}


// Жестко курить с захаром
func (muc *MedUseCase) Name1(ctx context.Context){
	
	// prepare request
	// call remote procedure
	return , nil
}

func (muc *MedUseCase) Name2(ctx context.Context) {
	
	// prepare request
	// call remote procedure
	return , nil
}

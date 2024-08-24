package validation

import (
	"fmt"
	pb "yir/auth/api/v0/auth"
	"yir/auth/internal/enity"
)

func ValidateTokenRefreshRequest(req *pb.TokenRefreshRequest) error {
	if req == nil {
		return fmt.Errorf("token refresh request is nil")
	}

	if validate.Var(req.RefreshToken, "required,token") != nil {
		return enity.ErrInvalidToken
	}

	return nil
}

package validation

import (
	"fmt"
	pb "yir/auth/api/auth"
	"yir/auth/internal/entity"
)

func ValidateTokenRefreshRequest(req *pb.TokenRefreshRequest) error {
	if req == nil {
		return fmt.Errorf("token refresh request is nil")
	}

	if validate.Var(req.RefreshToken, "required,token") != nil {
		return entity.ErrInvalidToken
	}

	return nil
}

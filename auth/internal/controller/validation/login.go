// плохое решение создавать в ините
package validation

import (
	"fmt"
	pb "yir/auth/api/v0/auth"
)

func ValidateLoginRequest(req *pb.LoginRequest) error {
	if req == nil {
		return fmt.Errorf("login request is nil")
	}

	if validate.Var(req.Email, "required,email") != nil {
		return fmt.Errorf("email must be smt@smt.smt")
	}

	if validate.Var(req.Password, "required,password") != nil {
		return fmt.Errorf("password must not nul, contains upper, lower case, digit(s) and min 8 length")
	}

	return nil
}

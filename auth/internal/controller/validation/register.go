package validation

import (
	"fmt"
	pb "yir/auth/api/v0/auth"
)

func ValidateRegisterRequest(req *pb.RegisterRequest) error {
	if req == nil {
		return fmt.Errorf("register request request is nil")
	}

	if validate.Var(req.Email, "required,email") != nil {
		return fmt.Errorf("email must be smt@smt.smt")
	}

	if validate.Var(req.LastName, "required") != nil {
		return fmt.Errorf("last name must be set")
	}

	if validate.Var(req.FirstName, "required") != nil {
		return fmt.Errorf("first name must be set")
	}

	if validate.Var(req.FathersName, "required") != nil {
		return fmt.Errorf("fathers name must be set")
	}

	if validate.Var(req.MedOrganization, "required") != nil {
		return fmt.Errorf("med organization must be set")
	}

	if validate.Var(req.Password, "required,password") != nil {
		return fmt.Errorf("password must not nul, contains upper, lower case, digit(s) and min 8 length")
	}

	return nil
}

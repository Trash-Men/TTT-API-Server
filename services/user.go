package services

type UserService struct{}

func (_ UserService) Login(id, password string) (string, error) {
	hashedPassword, error := PasswordService{}.Hash(password)

	if error != nil {
		return "", error
	}

	user, error := _repositories.UserRepository.GetUser(id, hashedPassword)

	if error != nil {
		return "", error
	}

	error = PasswordService{}.Validate(password, user.Password)

	if error != nil {
		return "", error
	}

	token, error := JwtService{}.CreateToken(user.Id, user.Role)

	if error != nil {
		return "", error
	}

	return token, nil
}

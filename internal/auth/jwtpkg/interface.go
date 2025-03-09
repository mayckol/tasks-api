package jwtpkg

type TokenServiceInterface interface {
	GenerateToken(userClaims UserClaims) (string, *UserClaims, error)
	VerifyToken(tokenStr string) (*UserClaims, error)
}

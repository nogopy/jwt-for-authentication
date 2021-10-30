package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/nogopy/jwt-for-authentication/internal/app/jwt-for-authentication/repositories"
	"github.com/nogopy/jwt-for-authentication/internal/app/jwt-for-authentication/utils"
	"github.com/nogopy/jwt-for-authentication/internal/app/jwt-for-authentication/utils/exception"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"time"
)

type LoginService struct {
	userRepo repositories.UserRepositoryInterface
}

func NewLoginService(
	userRepo repositories.UserRepositoryInterface,
) *LoginService {
	return &LoginService{
		userRepo: userRepo,
	}
}

type CustomClaimsExample struct {
	*jwt.StandardClaims
	UserId int `json:"user_id"`
}

func (service *LoginService) Login(ctx *gin.Context) (LoginResultSchema, error) {
	var body RegisterBodySchema

	if err := ctx.ShouldBindJSON(&body); err != nil {
		log.Errorf("cannot bind request body, exception: %s", err.Error())
		return LoginResultSchema{}, err
	}

	verifyBytes, err := ioutil.ReadFile("/Users/lap01651/nogopy/jwt-for-authentication/jwtRS256.key")
	// TODO: Please set private key in your env, not in your code!
	verifyKey, err := jwt.ParseRSAPrivateKeyFromPEM(verifyBytes)

	if err != nil {
		log.Errorf("cannot parse key, exception: %s", err.Error())
		return LoginResultSchema{}, err
	}

	user, err := service.userRepo.FindByUsername(body.UserName)
	if err != nil {
		return LoginResultSchema{}, err
	}

	if user.ID <= 0 {
		return LoginResultSchema{}, exception.WrongPassword
	}

	if user.Password != utils.Hash(body.Password) {
		return LoginResultSchema{}, exception.WrongPassword
	}

	t := jwt.New(jwt.GetSigningMethod("RS256"))

	t.Claims = &CustomClaimsExample{
		&jwt.StandardClaims{
			Issuer:    "nocopy",
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
		},
		user.ID,
	}

	signedToken, _ := t.SignedString(verifyKey)

	return LoginResultSchema{
		AccessToken:  signedToken,
		RefreshToken: "not ready yet",
		TokenType:    "bearer",
		ExpiresIn:    600,
		ExpiredAt:    int(time.Now().Add(time.Minute * 10).Unix()),
		UserId:       user.ID,
	}, nil
}

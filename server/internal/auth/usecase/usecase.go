package usecase

import (
	"context"
	"errors"
	"github.com/antonpodkur/Blog/config"
	"github.com/antonpodkur/Blog/internal/auth"
	"github.com/antonpodkur/Blog/internal/models"
	"github.com/antonpodkur/Blog/pkg/db"
	"github.com/antonpodkur/Blog/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
	"time"
)

type authUsecase struct {
	cfg         *config.Config
	mongoClient *mongo.Client
	ctx         context.Context
}

func NewAuthUsecase(cfg *config.Config, mongoClient *mongo.Client, ctx context.Context) auth.Usecase {
	return &authUsecase{
		cfg:         cfg,
		mongoClient: mongoClient,
		ctx:         ctx,
	}
}

func (u *authUsecase) SignUp(user *models.SignUpInput) (*models.UserDBResponse, error) {
	usersCollection := db.OpenCollection(u.mongoClient, u.cfg, "users")

	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	user.Email = strings.ToLower(user.Email)
	user.PasswordConfirm = ""
	user.Verified = true
	user.Role = "user"

	hashedPassword, _ := utils.HashPassword(user.Password)
	user.Password = hashedPassword

	res, err := usersCollection.InsertOne(u.ctx, &user)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("user with that email already exist")
		}
		return nil, err
	}

	opt := options.Index()
	opt.SetUnique(true)
	index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}

	if _, err := usersCollection.Indexes().CreateOne(u.ctx, index); err != nil {
		return nil, errors.New("could not create index for email")
	}

	var newUser *models.UserDBResponse
	query := bson.M{"_id": res.InsertedID}

	err = usersCollection.FindOne(u.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}

	return newUser, nil
}

func (u *authUsecase) SignIn(input *models.SignInInput) (*models.UserDBResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (u *authUsecase) GetUserById(id string) (*models.UserDBResponse, error) {
	usersCollection := db.OpenCollection(u.mongoClient, u.cfg, "users")
	oid, _ := primitive.ObjectIDFromHex(id)

	var user *models.UserDBResponse

	query := bson.M{"_id": oid}
	err := usersCollection.FindOne(u.ctx, query).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.UserDBResponse{}, err
		}
		return nil, err
	}

	return user, nil
}

func (u *authUsecase) GetUserByEmail(email string) (*models.UserDBResponse, error) {
	usersCollection := db.OpenCollection(u.mongoClient, u.cfg, "users")
	var user *models.UserDBResponse

	query := bson.M{"email": strings.ToLower(email)}
	err := usersCollection.FindOne(u.ctx, query).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.UserDBResponse{}, err
		}
		return nil, err
	}

	return user, nil

}

package model

import (
	"fmt"
	"time"

	"github.com/TonyChinwe/libraryapp/config/constant"
	"github.com/TonyChinwe/libraryapp/utils/helpers"
	"github.com/google/uuid"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Username  string             `bson:"username,omitempty" json:"username,omitempty"`
	Role      string             `bson:"role,omitempty" json:"role,omitempty"`
	Password  string             `bson:"password,omitempty" json:"password"`
	Salt      string             `bson:"salt,omitempty" json:"salt,omitempty"`
	IsActive  bool               `bson:"isActive,omitempty" json:"isActive,omitempty"`
	CreatedAT time.Time          `bson:"createdAt,omitempty" json:"createdAt,omitempty"`
	UpdatedAT time.Time          `bson:"updatedAt,omitempty" json:"updatedAt,omitempty"`
}

type UserUpdated struct {
	ModifiedCount int64 `json:"modifiedcount"`
	Result        User
}

type UserDeleted struct {
	DeletedCount int64 `json:"deletedcount"`
}

func ComparePassword2(password1 string, password2 string) error {
	incoming := []byte(password1)
	existing := []byte(password2)
	fmt.Println("inc: " + string(incoming))
	fmt.Println("exist: " + string(existing))
	err := bcrypt.CompareHashAndPassword(incoming, existing)
	return err
}

func (u *User) ComparePassword(password string) error {
	incoming := []byte(password + u.Salt)
	existing := []byte(u.Password)
	fmt.Println("inc: " + string(incoming))
	fmt.Println(": " + string(existing))
	fmt.Println("salt: " + u.Salt)

	err := bcrypt.CompareHashAndPassword(existing, incoming)
	return err
}

func (u *User) Initialize() error {
	salt := uuid.New().String()
	passwordBytes := []byte(u.Password + salt)
	hash, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash[:])
	u.Salt = salt
	u.CreatedAT = helpers.CurrentDate()
	u.UpdatedAT = helpers.CurrentDate()
	u.IsActive = true
	u.Role = constant.UserRole
	return nil
}

type LoginReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshReq struct {
	Token string `json:"token"`
}

type SignupReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

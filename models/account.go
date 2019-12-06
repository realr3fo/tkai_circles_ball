package models

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"os"
	u "github.com/realr3fo/tkai_circles_ball/utils"
)

/*
JWT claims struct
*/
type Token struct {
	UserId uint
	jwt.StandardClaims
}

//a struct to rep user account
type Account struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
	Token string `json:"token";sql:"-"`
}

//Validate incoming user details...
func (account *Account) Validate() (map[string] interface{}, bool, error) {

	if len(account.Password) < 6 {
		err := errors.New("password length should be > 6")
		return u.Message(false, "Password is required"), false, err
	}

	//Username must be unique
	temp := &Account{}

	//check for errors and duplicate username
	err := GetDB().Table("accounts").Where("username = ?", account.Username).First(temp).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return u.Message(false, "Connection db error. Please retry"), false, err
	}
	if temp.Username != "" {
		err := errors.New("username is already in use by another user")
		return u.Message(false, "Username already in use by another user."), false, err
	}

	return u.Message(false, "Requirement passed"), true, nil
}

func (account *Account) Create() (map[string] interface{}, error) {

	if resp, ok, err := account.Validate(); !ok {
		return resp, err
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(account.Password), bcrypt.DefaultCost)
	account.Password = string(hashedPassword)

	GetDB().Create(account)

	if account.ID <= 0 {
		var err = errors.New("failed to create account, connection error")
		return u.Message(false, "Failed to create account, connection error."), err
	}

	//Create new JWT token for the newly registered account
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString

	account.Password = "" //delete password

	response := u.Message(true, "Account has been created")
	response["account"] = account
	return response, nil
}

func Login(username, password string) (map[string] interface{}, error) {

	account := &Account{}
	err := GetDB().Table("accounts").Where("username = ?", username).First(account).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return u.Message(false, "Username not found"), err
		}
		return u.Message(false, "Connection error. Please retry"), err
	}

	err = bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword { //Password does not match!
		return u.Message(false, "Invalid login credentials. Please try again"), err
	}
	//Worked! Logged In
	account.Password = ""

	//Create JWT token
	tk := &Token{UserId: account.ID}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tk)
	tokenString, _ := token.SignedString([]byte(os.Getenv("token_password")))
	account.Token = tokenString //Store the token in the response

	resp := u.Message(true, "Logged In")
	resp["account"] = account
	return resp, nil
}

func GetUser(u uint) *Account {

	acc := &Account{}
	GetDB().Table("accounts").Where("id = ?", u).First(acc)
	if acc.Username == "" { //User not found!
		return nil
	}

	acc.Password = ""
	return acc
}
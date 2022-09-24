package utils

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"main/model"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

var (
	lowerCharSet   = "abcdedfghijklmnopqrst"
	upperCharSet   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharSet = "!@#$%&*"
	numberSet      = "0123456789"
	allCharSet     = lowerCharSet + upperCharSet + specialCharSet + numberSet
	mySigningKey   = []byte("DFGDFGhcsadkjhfwe+ě+23123asldxjhsdljfh1234234")
)

type ClientData struct {
	Db *gorm.DB
}

type userPermission struct {
	ParentId uuid.UUID
	Id       uuid.UUID
	Name     string
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func IsValidUUID(uuid uuid.UUID) bool {
	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid.String())
}

func IsAuthorized(w http.ResponseWriter, r *http.Request, site uuid.UUID, c ClientData) (bool, string) {
	if r.Header["Authorization"] != nil {
		var sendToken = strings.Replace(r.Header["Authorization"][0], "Bearer ", "", 1)
		token, err := jwt.Parse(sendToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("JWT token not pass")
			}
			return mySigningKey, nil
		})

		claims := token.Claims.(jwt.MapClaims)
		// check if store belongs to account
		if IsValidUUID(site) {
			var user userPermission
			idFromClaim, _ := uuid.Parse(fmt.Sprintf("%v", claims["id"]))
			c.Db.Model(&model.User{Id: idFromClaim}).First(&user)
			if IsValidUUID(user.ParentId) {
				if isPermitted(user.ParentId, site, c) == false {
					http.Error(w, "Not Permitted", http.StatusForbidden)
					return false, ""
				}
			} else {
				if isPermitted(user.Id, site, c) == false {
					http.Error(w, "Not Permitted", http.StatusForbidden)
					return false, ""
				}
			}
		}

		if err != nil && token != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Fatalf(err.Error())
			return false, ""
		}
		return true, fmt.Sprintf("%v", claims["id"])
	} else {
		http.Error(w, "Not Authorized", http.StatusForbidden)
		return false, ""
	}
}

func SetupCORS(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

func GenerateJWT(name string, id string) (map[string]string, error) {

	// access token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["client"] = name
	claims["id"] = id
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		log.Fatalf(err.Error())
	}

	// refresh token
	refreshToken := jwt.New(jwt.SigningMethodHS256)
	rtClaims := refreshToken.Claims.(jwt.MapClaims)
	rtClaims["id"] = id
	rtClaims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	refreshTokenString, err := refreshToken.SignedString(mySigningKey)

	if err != nil {
		log.Fatalf(err.Error())
	}

	return map[string]string{
		"access_token":  tokenString,
		"refresh_token": refreshTokenString,
	}, nil
}

func isPermitted(userId uuid.UUID, siteId uuid.UUID, c ClientData) bool {
	var site model.Site
	// special check for exit uuid
	c.Db.First(&model.Site{}, "id = ?", siteId).Scan(&site)
	if IsValidUUID(site.Id) == false {
		return true
	} else {
		c.Db.First(&model.Site{}, "user_id = ? AND id = ?", userId, siteId).Scan(&site)
		return IsValidUUID(site.Id)
	}
}

func GeneratePassword(passwordLength, minSpecialChar, minNum, minUpperCase int) string {
	var password strings.Builder

	//Set special character
	for i := 0; i < minSpecialChar; i++ {
		random := rand.Intn(len(specialCharSet))
		password.WriteString(string(specialCharSet[random]))
	}

	//Set numeric
	for i := 0; i < minNum; i++ {
		random := rand.Intn(len(numberSet))
		password.WriteString(string(numberSet[random]))
	}

	//Set uppercase
	for i := 0; i < minUpperCase; i++ {
		random := rand.Intn(len(upperCharSet))
		password.WriteString(string(upperCharSet[random]))
	}

	remainingLength := passwordLength - minSpecialChar - minNum - minUpperCase
	for i := 0; i < remainingLength; i++ {
		random := rand.Intn(len(allCharSet))
		password.WriteString(string(allCharSet[random]))
	}
	inRune := []rune(password.String())
	rand.Shuffle(len(inRune), func(i, j int) {
		inRune[i], inRune[j] = inRune[j], inRune[i]
	})
	return string(inRune)
}

func GetClaim(sendToken string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(sendToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("JWT token not pass")
		}
		return mySigningKey, nil
	})
	claims := token.Claims.(jwt.MapClaims)
	return claims, err
}

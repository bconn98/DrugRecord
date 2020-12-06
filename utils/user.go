/**
File: user
Description: Implements features for an unique user including passwords
@author Bryan Conn
@date 10/7/2018
*/
package utils

import (
	"log"
	"strings"

	"github.com/jimlawless/whereami"
	"golang.org/x/crypto/bcrypt"
)

/**
Declares constants that determine the validity of an entered password
*/
const (
	UE   = 0 // Empty Username
	US   = 1 // Username contains spaces
	PE   = 2 // Empty Password
	PS   = 3 // Password contains spaces
	PIV  = 4 // Password too long or short
	TN   = 5 // Taken Name
	GOOD = 6 // No issues
)

/**
Function: MakeUser
Description: Creates a new user and adds them to the database
@param acUsername The username of the new user
@param acPassword The password used for the user
@return The validity of the new user
*/
func MakeUser(acUsername string, acPassword string) int {
	lcValidCheck := validateInfo(acUsername, acPassword)
	if lcValidCheck != GOOD {
		return lcValidCheck
	}
	lcPassVal := computePassVal(acPassword)
	// Check if the username exists
	lsTestUser := User{}
	if FindUser(acUsername) != lsTestUser {
		return TN
	}
	AddUser(acUsername, lcPassVal)
	return GOOD
}

/**
Function: validateInfo
Description: Validates if the username and password are empty or contain spaces
@param acUsername The username of the new user
@param acPassword The password of the new user
@return If the new user's information is valid
*/
func validateInfo(acUsername string, acPassword string) int {
	lnMaxPasswordLen := 72
	lnMinPasswordLen := 10
	if acUsername == "" {
		return UE
	} else if strings.Contains(acUsername, " ") {
		return US
	} else if acPassword == "" {
		return PE
	} else if strings.Contains(acPassword, " ") {
		return PS
	} else if len(acPassword) < lnMinPasswordLen || len(acPassword) > lnMaxPasswordLen {
		return PIV
	} else {
		return GOOD
	}
}

/**
Function: FindUser
Description: Looks for user in the database
@param acName The name of the user to be found
@return The user if its found, if not a blank user
*/
func FindUser(acName string) User {
	lsUsers := GetUsers()
	for _, lsUser := range lsUsers {
		if lsUser.GetUserName() == acName {
			return lsUser
		}
	}
	return User{}
}

/**
Function: GetUserName
Description: Returns the username of the user
@return The name of the user
*/
func (asUser User) GetUserName() string {
	return asUser.UserName
}

/**
Function: CheckPassword
Description: Determines if the password matches the users password
@param asUser The user that is being checked
@param acPassword The password that was entered
@return If the password matches
*/
func CheckPassword(asUser User, acPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(asUser.PassVal), []byte(acPassword))
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

/**
Function: computePassVal
Description: Computes the value of given password using a unique formula (will be changed when released)
@param acPassword The password being computed
@return The value of the password as a hash and salt
*/
func computePassVal(acPassword string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(acPassword), bcrypt.MinCost)
	issue(err, whereami.WhereAmI())
	return string(hash)
}

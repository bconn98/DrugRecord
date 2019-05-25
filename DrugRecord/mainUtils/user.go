/**
File: user
Description: Implements features for an unique user including passwords
@author Bryan Conn
@date 10/7/2018
*/
package mainUtils

import "strings"

/**
Declares constants that determine the validity of an entered password
*/
const (
	UE   = 0 // Empty Username
	US   = 1 // Username contains spaces
	PE   = 2 // Empty Password
	PS   = 3 // Password contains spaces
	TN   = 4 // Taken Name
	GOOD = 5 // No issues
)

/**
Function: MakeUser
Description: Creates a new user and adds them to the database
@param username The username of the new user
@param password The password used for the user
@return The validity of the new user
*/
func MakeUser(username string, password string) int {
	validCheck := validateInfo(username, password)
	if validCheck != GOOD {
		return validCheck
	}
	passVal := computePassVal(password)
	//Check if the username exists
	test := User{}
	if FindUser(username) != test {
		return TN
	}
	AddUser(username, passVal)
	return GOOD
}

/**
Function: validateInfo
Description: Validates if the username and password are empty or contain spaces
@param username The username of the new user
@param password The password of the new user
@return If the new user's information is valid
*/
func validateInfo(username string, password string) int {
	if username == "" {
		return UE
	} else if strings.Contains(username, " ") {
		return US
	} else if password == "" {
		return PE
	} else if strings.Contains(password, " ") {
		return PS
	} else {
		return GOOD
	}
}

/**
Function: FindUser
Description: Looks for user in the database
@param name The name of the user to be found
@return The user if its found, if not a blank user
*/
func FindUser(name string) User {
	users := GetUsers()
	for _, user := range users {
		if user.GetUserName() == name {
			return user
		}
	}
	return User{}
}

/**
Function: GetUserName
Description: Returns the username of the user
@return The name of the user
*/
func (user User) GetUserName() string {
	return user.UserName
}

/**
Function: CheckPassword
Description: Determines if the password matches the users password
@param user The user that is being checked
@param password The password that was entered
@return If the password matches
*/
func CheckPassword(user User, password string) bool {
	return user.PassVal == computePassVal(password)
}

/**
Function: computePassVal
Description: Computes the value of given password using a unique formula (will be changed when released)
@param password The password being computed
@return The value of the password as an int
*/
func computePassVal(password string) int {
	var i int
	val := 0
	passwordLength := len(password)
	for i = 0; i < passwordLength; i++ {
		val += ((int(password[i])*31 + i) / 7) - 5
	}
	return val
}

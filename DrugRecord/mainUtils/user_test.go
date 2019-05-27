package mainUtils

/**
File: user_test
Description: Tests all the functions in the user file
@author Bryan Conn
@date 6/23/18
*/

import "testing"

/**
Function: TestMakeUser
Description: Tests if a user is properly made and can be found in the map of users
*/
func TestMakeUser(t *testing.T) {
	MakeUser("Bryan", "zoo123")
	testUser1 := FindUser("Bryan")
	testUser2 := User{"Bryan", 2157}
	if testUser1 != testUser2 {
		t.Error("Either make user or find user messed up")
	}
}

/**
Function: TestCheckPassword
Description: Tests if the password matches by testing the capital letters vs
lower case letters
*/
func TestCheckPassword(t *testing.T) {
	testUser := FindUser("Bryan")
	if !CheckPassword(testUser, "zoo123") {
		t.Error("Password doesn't match!")
	} else if CheckPassword(testUser, "Zoo123") {
		t.Error("The passwords should not be matching!")
	}
}

/**
Function: TestUser_GetUserName
Description: Tests if the username of a user can be properly accessed.
*/
func TestUser_GetUserName(t *testing.T) {
	testUser1 := User{"Bryan", 2157}
	if testUser1.GetUserName() != "Bryan" {
		t.Error("The username doesn't match!")
	}
}

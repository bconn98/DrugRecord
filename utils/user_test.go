package utils

/**
File: user_test
Description: Tests all the functions in the user file
@author Bryan Conn
@date 6/23/18
*/

import (
	"testing"
)

/**
Function: TestCheckPassword
Description: Tests if the password matches by testing the capital letters vs
lower case letters
*/
func TestCheckPassword(t *testing.T) {
	testUser := User{UserName: "test", PassVal: computePassVal("zoo123")}
	t.Log("Expect to see \"crypto/bcrypt: hashedPassword is not the hash of the given password\"")

	// This should fall through without a log
	if !CheckPassword(testUser, "zoo123") {
		t.Error("Password doesn't match!")
	} else if CheckPassword(testUser, "Zoo123") { // This should return false WITH a log
		t.Error("The passwords should not be matching!")
	}
}

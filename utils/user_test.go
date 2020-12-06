package utils

/**
File: user_test
Description: Tests all the functions in the user file
@author Bryan Conn
@date 6/23/18
*/

import "testing"

/**
Function: TestCheckPassword
Description: Tests if the password matches by testing the capital letters vs
lower case letters
*/
func TestCheckPassword(t *testing.T) {
	testUser := User{UserName: "test", PassVal: computePassVal("zoo123")}
	if !CheckPassword(testUser, "zoo123") {
		t.Error("Password doesn't match!")
	} else if CheckPassword(testUser, "Zoo123") {
		t.Error("The passwords should not be matching!")
	}
}
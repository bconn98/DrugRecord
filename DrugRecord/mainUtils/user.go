package mainUtils

import "strings"


type User struct {
	UserName string
	PassVal int
}

const (
	UE = 0// Empty Username
	US = 1// Username contains spaces
	PE = 2// Empty Password
	PS = 3 // Password contains spaces
	TN = 4 // Taken Name
	GOOD = 5 // No issues
)

func MakeUser(username string, password string) (int) {
	validNum := validateInfo(username, password)
	if validNum != GOOD {
		return validNum
	}
	passVal := computePassVal(password)
	test := User{}
	if FindUser(username) != test {
		return TN
	}
	AddUser(username, passVal)
	return GOOD
}

func validateInfo(username string, password string) (int) {
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

func FindUser(name string) (User){
	users := GetUsers()
	for _, user := range users {
		if user.GetUserName() == name {
			return user
		}
	}
	return User{}
}

func (user User) GetUserName() (string) {
	return user.UserName
}

func CheckPassword(user User, password string) (bool){
	 return user.PassVal == computePassVal(password)
}

func computePassVal(password string) (int) {
	var i int
	val := 0
	len := len(password)
	for i = 0; i < len; i++ {
		val += ( ( int(password[i]) * 31 + i ) / 7 ) - 5
	}
	return val
}
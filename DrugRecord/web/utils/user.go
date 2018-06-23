package utils

type User struct {
	userName string
	passVal int
}

var users = make(map[string]User)

func MakeUser(name string, password string) {
	passVal := computePassVal(password)
	user := User{name, passVal}
	users[name] = user
}

func FindUser(name string) (User){
	if user, ok := users[name]; ok {
		return user
	}
	return User{}
}

func (user User) GetUserName() (string) {
	return user.userName
}

func CheckPassword(user User, password string) (bool){
	 return user.passVal == computePassVal(password)
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
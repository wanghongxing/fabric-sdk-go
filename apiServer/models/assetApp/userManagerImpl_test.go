package assetApp

import (
	"fmt"
	"testing"
)

/*func TestUserManager(t *testing.T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	name := "bill" + strconv.Itoa(r.Intn(1000)) //note: the len(name) in database is 18. len(passwd) is 20 bytes
	register(name, t)
	login(name, name, t) //passwd=name
	updateUser(name, t)
}

func register(name string, t *testing.T) {
	impl := UserManagerImpl{}

	u := &user.User{
		Name:   name,
		Passwd: name,
		Email:  name + "@linux.com",
		Phone:  "123456",
	}

	id, ok := impl.Register(u)
	if !ok {
		t.Error("Register failed")
	} else {
		t.Log(id)
	}
}

func login(name, passwd string, t *testing.T) {
	impl := UserManagerImpl{}
	signedToken, err := impl.Login(name, passwd)
	if err != nil {
		t.Fatal("Login failed", err)
	} else {
		t.Log("signedToken", signedToken)
	}
}

func updateUser(name string, t *testing.T) {
	impl := UserManagerImpl{}
	u := &user.User{
		Name:  name,
		Email: "billhan@linux.com",
		Phone: "110",
	}
	err := impl.UpdateInfo(u)
	if err != nil {
		t.Error("Update failed")
	} else {
		fmt.Println("update success")
	}
}*/

func TestGetUser(t *testing.T) {
	fmt.Println("hh")
	impl := UserManagerImpl{}
	userInfo, err := impl.GetUserInfo("hxy")
	if err != nil {
		t.Errorf("getUser failed")
	} else {
		fmt.Println("get user success")
		t.Log("user: ", userInfo)
	}
}

/*
Copyright Beijing Sansec Technology Development Co., Ltd. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at


      http://www.apache.org/licenses/LICENSE-2.0


Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package user

import (
	"errors"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id     int    `json:"id"`   //primary key
	Type   int    `json:"type"` //0: admin, 1: user
	Name   string `json:"name"` //unique
	Passwd string `json:"passwd"`
	Email  string `json:"mail"`
	Phone  string `json:"phone"`
}
type Secret struct {
	Name   string `json:"name"`
	Passwd string `json:"passwd"`
}
type UpdateUserArgs struct {
	Type  string `json:"type"` //0: admin, 1: user
	Name  string `json:"name"` //unique
	Email string `json:"mail"`
	Phone string `json:"phone"`
}

//temporary data structure, will use hardware device in the fulture.
type UserIdentity struct {
	UserId       int    `json:"userId"`
	UserName     string `json:"userName"`
	EnrollName   string `json:"enrollName"`
	EnrollSecret string `json:"enrollSecret"`
}

/*
type Identity struct {
	Id          int    `json:"id"`
	Key         string `json:"key"`
	Certificate string `json:"certificate"`
	UserId      int    `json:"userId"`
}
*/

func AddUser(user *User) (id int64, err error) {
	o := orm.NewOrm()

	id64, err := o.Insert(user)
	fmt.Println(id64, err)
	if err != nil {
		return -1, err
	}
	return id64, nil
}

func GetUserByName(username string) (*User, error) {
	o := orm.NewOrm()
	u := User{}

	err := o.Raw("SELECT id, type, name, passwd, phone, email FROM user WHERE name = ?", username).QueryRow(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func GetUserById(userid int) (*User, error) {
	o := orm.NewOrm()
	u := User{}
	err := o.Raw("SELECT id, type, name, passwd, phone, email FROM user WHERE id = ?", userid).QueryRow(&u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}
func UpdateUser(newU *UpdateUserArgs) error {
	o := orm.NewOrm()
	oldU := User{}

	err := o.Raw("SELECT * from user WHERE name = ?", newU.Name).QueryRow(&oldU)
	if err != nil {
		return err
	}
	if newU.Email != "" {
		oldU.Email = newU.Email
	}
	if newU.Phone != "" {
		oldU.Phone = newU.Phone
	}
	if newU.Type != "" {
		typ, err := strconv.Atoi(newU.Type)
		if err != nil {
			return err
		}
		oldU.Type = typ
	}
	_, err = o.Update(&oldU)
	if err != nil {
		return err
	}

	return nil
}
func UpdatePasswd(name string, oldPwd string, newPwd string) error {
	o := orm.NewOrm()
	oldU := User{}
	err := o.Raw("SELECT * from user WHERE name = ?", name).QueryRow(&oldU)
	if err != nil {
		return err
	}
	if oldPwd != oldU.Passwd {
		return errors.New("passwd incorrect")
	}
	oldU.Passwd = newPwd
	_, err = o.Update(&oldU)
	if err != nil {
		return err
	}

	return nil
}

func Login(ss *Secret) (*User, error) {
	o := orm.NewOrm()
	u := User{}

	err := o.Raw("SELECT * FROM user WHERE name = ?", ss.Name).QueryRow(&u)
	if err != nil {
		return nil, err
	}

	//	fmt.Println(reflect.TypeOf(passwd), reflect.TypeOf(u.Passwd))

	if u.Passwd == ss.Passwd {
		return &u, nil
	}
	return nil, errors.New("Invalid password")
}

func init() {
	//register model
	orm.RegisterModel(new(User))
	//register driver
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//set default database
	user := beego.AppConfig.String("mysqluser")
	pwd := beego.AppConfig.String("mysqlpass")
	url := beego.AppConfig.String("mysqlurls")
	db := beego.AppConfig.String("mysqldb")
	connection := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8", user, pwd, url, db)
	orm.RegisterDataBase("default", "mysql", connection, 30)

	/*
		//max idle connections
		orm.SetMaxIdleConns("default", 30)
		//max opened connections
		orm.SetMaxOpenConns("default", 30)
	*/
}

package models

import (
	"agile/pkg/dbManager"
	"database/sql"
	"errors"
	"fmt"
)

type User struct {
	Id          int      `json:"id"`
	Telephone   string   `json:"phone"`
	Password    string   `json:"password"`
	AccessToken string   `json:"accessToken"`
	Roles       []string `json:"roles"`
	Blocked     bool     `json:"blocked"`
}

type Role struct {
	Id   int    `json:"id"`
	Role string `json:"role"`
}

func (u *User) SignIn() error {
	var (
		err   error
		exist int
	)

	err = dbManager.Get().QueryRow(`select count(*),telnumber,id,blocked from public.users where telnumber=$1 and pass=$2 group by telnumber,id`, u.Telephone, u.Password).Scan(&exist, &u.Telephone, &u.Id, &u.Blocked)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return fmt.Errorf("Неверный логин или пароль")
		}
		return fmt.Errorf("SignIn err: %v", err)
	}

	_, banned, _ := u.CheckBan(u.Telephone)
	if banned {
		return errors.New("Пользователь заблокирован")
	}

	//select user roles
	rows, err := dbManager.Get().Query(`select distinct r.rname from public.users u inner join roles r on u.fk_role = r.id  where telnumber=$1`, u.Telephone)
	if err != nil {
		return fmt.Errorf("SignIn err: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var role string
		rows.Scan(&role)
		u.Roles = append(u.Roles, role)
	}

	fmt.Println("roles:", u.Roles)
	return nil
}

func (u *User) SignUp() error {
	var err error

	exists, err := u.checkExists()
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("Такой пользователь уже существует")
	}

	//TODO change set role
	_, err = dbManager.Get().Exec(`insert into public.users(telnumber,pass,fk_role) values ($1,$2,$3)`, u.Telephone, u.Password, 1)
	if err != nil {
		return fmt.Errorf("SignUp err: %v", err)
	}

	return nil
}

func (u *User) Select() (User, error) {
	var user User
	err := dbManager.Get().QueryRow(`select id,telnumber,pass,blocked from public.users where id=$1`, u.Id).Scan(&user.Id, &user.Telephone, &user.Password, &user.Blocked)
	if err != nil {
		fmt.Errorf("err select user:%v", err)
	}
	return user, err
}

func (u *User) Update() error {
	selectedUser, err := u.Select()
	if err != nil {
		return err
	}

	_, err = dbManager.Get().Exec(`update public.users set telnumber =$1,pass =$2 where id=$3 `, u.Telephone, u.Password, selectedUser.Id)
	if err != nil {
		return fmt.Errorf("update selectedUser err: %v", err)
	}
	return err
}

func (u *User) checkExists() (bool, error) {
	var (
		exist int
		err   error
	)

	err = dbManager.Get().QueryRow(`select count(*) from public.users where telnumber=$1`, u.Telephone).Scan(&exist)
	if err != nil {
		return false, fmt.Errorf("checkExist err: %v", err)
	}

	return exist > 0, err
}

func (u *User) GetAll() ([]User, error) {
	var users = make([]User, 0)

	rows, err := dbManager.Get().Query(`select id,telnumber,pass from public.users`)
	if err != nil {
		fmt.Errorf("err select user:%v", err)
	}
	defer rows.Close()

	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Id, &user.Telephone, &user.Password)
		if err != nil {
			fmt.Println("err scan user.geall")
			return nil, err
		}
		users = append(users, user)
	}

	return users, err
}

func (u *User) SetRole(phone string, role int) error {
	_, banned, err := u.CheckBan(phone)
	if banned {
		return fmt.Errorf("Такого пользователя не существует!")
	}

	_, err = dbManager.Get().Exec(`update public.users set fk_role=$1 where telnumber=$2 `, role, phone)
	if err != nil {
		fmt.Println("SetRole() ", err)
		return fmt.Errorf("Такого пользователя не существует!")
	}
	return err
}

func (u *User) Ban(phone string) error {
	_, userExists, err := u.GetByPhone(phone)
	if err != nil {
		fmt.Println("Ban() ", err)
		return fmt.Errorf("Такого пользователя не существует")
	}
	if !userExists {
		return fmt.Errorf("Такого пользователя не существует")

	}

	_, err = dbManager.Get().Exec(`insert into public.banned(telnumber) values($1)`, phone)
	if err != nil {
		return fmt.Errorf("setrole err: %v", err)
	}
	return err
}

func (u *User) CheckBan(phone string) (int, bool, error) {
	var id int
	err := dbManager.Get().QueryRow(`select id from public.banned where telnumber=$1`, phone).Scan(&id)
	if err != nil {
		return id, false, fmt.Errorf("CheckBan err: %v", err)
	}
	if id != 0 {
		return id, true, err
	}
	return id, false, err
}

func (u *User) GetByPhone(phone string) (int, bool, error) {
	var id int
	fmt.Println("PHONE PIDARAS ", phone)
	err := dbManager.Get().QueryRow(`select id from public.users where telnumber=$1`, phone).Scan(&id)
	if err != nil {
		return id, false, fmt.Errorf("CheckBan err: %v", err)
	}
	fmt.Println(id)
	if id != 0 {
		return id, true, err
	}
	return id, false, err
}

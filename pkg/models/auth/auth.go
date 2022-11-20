package auth

import (
	"agile/pkg/dbManager"
	"database/sql"
	"fmt"
)

type User struct {
	Id          int      `json:"id"`
	Telephone   string   `json:"phone"`
	Password    string   `json:"password"`
	AccessToken string   `json:"accessToken"`
	Roles       []string `json:"roles"`
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

	err = dbManager.Get().QueryRow(`select count(*),telnumber from public.users where telnumber=$1 and pass=$2 group by telnumber`, u.Telephone, u.Password).Scan(&exist, &u.Telephone)
	if err != nil {
		if err == sql.ErrNoRows {
			err = nil
			return fmt.Errorf("bad login or password")
		}
		return fmt.Errorf("SignIn err: %v", err)
	}

	if exist == 0 {
		return fmt.Errorf("user not found")
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
		return fmt.Errorf("user already exist")
	}

	//TODO change set role
	_, err = dbManager.Get().Exec(`insert into public.users(telnumber,pass,fk_role) values ($1,$2,$3)`, u.Telephone, u.Password, 2)
	if err != nil {
		return fmt.Errorf("SignUp err: %v", err)
	}

	return nil
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

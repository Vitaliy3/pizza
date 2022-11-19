package auth

import (
	"agile/pkg/dbManager"
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
		count int
	)

	err = dbManager.Get().QueryRow(`select count(*),telnumber from public.users where telnumber=$1 and pass=$2 group by telnumber`, u.Telephone, u.Password).Scan(&count, &u.Telephone)
	if err != nil {
		return fmt.Errorf("SignIn err: %v", err)
	}

	// user not found
	if count == 0 {
		return fmt.Errorf("user not found")
	}

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

	_, err = dbManager.Get().Exec(`insert into public.users(telnumber,pass) values ($1,$2)`, u.Telephone, u.Password)
	if err != nil {
		return fmt.Errorf("SignUp err: %v", err)
	}

	return nil
}

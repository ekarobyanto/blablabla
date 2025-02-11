package users

import "github.com/jmoiron/sqlx"

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) FindAll() ([]User, error) {
	users := []User{}
	err := u.db.Select(&users, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepo) FindByID(id int) (*User, error) {
	user := User{}
	err := u.db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *UserRepo) Create(user *User) error {
	_, err := u.db.NamedExec("INSERT INTO users (name, email, password) VALUES (:name, :email, :password)", user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepo) FindByEmail(email string) (*User, error) {
	user := User{}
	err := u.db.Get(&user, "SELECT * FROM users WHERE email = $1", email)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

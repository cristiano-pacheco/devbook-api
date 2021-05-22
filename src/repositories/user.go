package repositories

import (
	"api/src/models"
	"database/sql"
)

type user struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *user {
	return &user{db}
}

func (repository user) Create(user models.User) (uint64, error) {
	stmt, err := repository.db.Prepare(
		"insert into users (name, nick, email, password) values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(user.Name, user.Nick, user.Email, user.Password)
	if err != nil {
		return 0, err
	}

	lastInsertedID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(lastInsertedID), nil
}

func (repository user) Search(nameOrNick string) ([]models.User, error) {
	nameOrNick = "%" + nameOrNick + "%"

	rows, err := repository.db.Query(
		"select id, name, nick, email, created_at from users where name like ? or nick like ?",
		nameOrNick,
		nameOrNick,
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User

		err = rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (repository user) Get(ID uint64) (models.User, error) {
	rows, err := repository.db.Query(
		"select id, name, nick, email, created_at from users where id = ?",
		ID,
	)

	if err != nil {
		return models.User{}, err
	}

	defer rows.Close()

	var user models.User

	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Name, &user.Nick, &user.Email, &user.CreatedAt)
		if err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (repository user) Update(ID uint64, user models.User) error {
	stmt, err := repository.db.Prepare(
		"update users set name = ?, nick = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Nick, user.Email, ID)
	if err != nil {
		return err
	}

	return nil
}

func (repository user) Delete(ID uint64) error {
	stmt, err := repository.db.Prepare(
		"delete from users where id = ?",
	)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(ID)
	if err != nil {
		return err
	}

	return nil
}

func (repository user) GetByEmail(email string) (models.User, error) {
	user := models.User{}

	stmt, err := repository.db.Prepare("select id, password from users where email = ?")
	if err != nil {
		return user, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(email).Scan(&user.ID, &user.Password)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (repository user) Follow(userID, followerID uint64) error {
	stmt, err := repository.db.Prepare("insert ignore into followers values (user_id, follower_id) values (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, followerID)
	if err != nil {
		return err
	}

	return nil
}

func (repository user) UnFollow(userID, followerID uint64) error {
	stmt, err := repository.db.Prepare("delete from followers where user_id = ? and follwer_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userID, followerID)
	if err != nil {
		return err
	}

	return nil
}

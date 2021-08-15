package models

import "errors"

type Person struct {
	Id        int    `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func AllPerson() ([]Person, error) {
	rows, err := Database.Query("SELECT * FROM people")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var persons []Person

	for rows.Next() {
		var person Person

		err = rows.Scan(&person.Id, &person.Firstname, &person.Lastname)
		if err != nil {
			return nil, err
		}

		persons = append(persons, person)
	}

	return persons, nil
}

func AddPerson(firstname string, lastname string) error {
	statement, _ := Database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
	defer statement.Close()

	_, err := statement.Exec(firstname, lastname)
	if err != nil {
		return err
	}

	return nil
}

func PersonById(id int) (*Person, error) {
	row, err := Database.Query("SELECT * FROM people WHERE id = ?", id)
	if err != nil {
		return nil, err
	}
	defer row.Close()

	var person Person

	if row.Next() {
		err = row.Scan(&person.Id, &person.Firstname, &person.Lastname)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("id does not exist")
	}

	return &person, nil
}

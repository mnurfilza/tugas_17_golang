package tugas17

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Mahasiswa struct {
	Npm   string
	Nama  string
	Kelas string
}

func Connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/db_belajar_golang")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func Tugas17() ([]Mahasiswa, error) {
	db, err := Connect()
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer db.Close()
	rows, err := db.Query("select npm, nama, kelas from tb_student")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer rows.Close()

	var res []Mahasiswa

	for rows.Next() {
		each := Mahasiswa{}
		err := rows.Scan(&each.Npm, &each.Nama, &each.Kelas)
		if err != nil {
			return nil, err
		}

		res = append(res, each)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return res, nil
}

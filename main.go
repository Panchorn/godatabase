package main

import (
	// "database/sql"
	"errors"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func main() {
	var err error
	db, err = sqlx.Open("mysql", "local_user:p@ssw0rd@tcp(127.0.0.1:3306)/local_db")
	if err != nil {
		panic(err)
	}

	// ====================================
	// GET COVERS
	// ====================================
	// covers, err := GetCovers()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }

	// for i, cover := range covers {
	// 	fmt.Printf("%d %#v\n", i, cover)
	// }
	// ====================================

	// ====================================
	// GET COVER
	// ====================================
	// cover, err := GetCover(99)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Printf("%#v\n", cover)
	// ====================================

	// ====================================
	// ADD COVER
	// ====================================
	// cover := Cover {
	// 	Id: 99,
	// 	Name: "cover-non",
	// }
	// err = AddCover(cover)
	// if err!= nil {
    //     fmt.Println(err)
    //     return
    // }
	// ====================================

	// ====================================
	// UPDATE COVER
	// ====================================
	// cover := Cover {
	// 	Id: 99,
	// 	Name: "cover-non-2",
	// }
	// err = UpdateCover(cover)
	// if err!= nil {
    //     fmt.Println(err)
    //     return
    // }
	// ====================================

	// ====================================
	// DELETE COVER
	// ====================================
	// err = DeleteCover(99)
	// if err!= nil {
    //     fmt.Println(err)
    //     return
    // }
	// ====================================

	// ====================================
	// GET COVERS X
	// ====================================
	covers, err := GetCoversX()
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, cover := range *covers {
		fmt.Printf("%d %#v\n", i, cover)
	}
	// ====================================

	// ====================================
	// GET COVER X
	// ====================================
	cover, err := GetCoverX(3)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v\n", cover)
	fmt.Println(cover)
	fmt.Println(*cover)
	// ====================================

}

type Cover struct {
	Id   int
	Name string
}

func GetCovers() ([]Cover, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	query := "select id, name from cover"
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	covers := []Cover{}
	for rows.Next() {
		cover := Cover{}
		err = rows.Scan(&cover.Id, &cover.Name)
		if err != nil {
			return nil, err
		}

		covers = append(covers, cover)
	}
	return covers, nil
}

func GetCover(id int) (*Cover, error) {
	err := db.Ping()
	if err != nil {
		return nil, err
	}

	query := "select id, name from cover where id=?"
	row := db.QueryRow(query, id)

	cover := Cover{}
	err = row.Scan(&cover.Id, &cover.Name)
	if err != nil {
		return nil, err
	}

	return &cover, nil
}

func AddCover(cover Cover) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	query := "insert into cover (id, name) values (?, ?)"

	result, err := db.Exec(query, cover.Id, cover.Name)
	if err!= nil {
        return err
    }
	affected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
    }
	if affected <= 0 {
		return errors.New("cover not inserted")
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func UpdateCover(cover Cover) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	query := "update cover set name=? where id=?"

	result, err := db.Exec(query, cover.Name, cover.Id)
	if err!= nil {
        return err
    }
	affected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
    }
	if affected <= 0 {
		return errors.New("cover not updated")
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func DeleteCover(id int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	query := "delete from cover where id=?"

	result, err := db.Exec(query, id)
	if err!= nil {
        return err
    }
	affected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
    }
	if affected <= 0 {
		return errors.New("cover not deleted")
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func GetCoversX() (*[]Cover, error) {
	query := "select id, name from cover"
	covers := []Cover{}
	err := db.Select(&covers, query)
	if err!= nil {
        return nil, err
    }
	return &covers, nil
}

func GetCoverX(id int) (*Cover, error) {
	query := "select id, name from cover where id=?"
	cover := Cover{}
	err := db.Get(&cover, query, id)
	if err != nil {
		return nil, err
	}
	return &cover, nil
}

package data

import (
	"context"
	"database/sql"
	"log"
	"time"
)

const dbTimeout = time.Second * 3

var db *sql.DB

// New is the function used to create an instance of the data package. It returns the type
// Model, which embeds all the types we want to be available to our application.
func New(dbPool *sql.DB) Models {
	db = dbPool

	return Models{
		Language:     Language{},
		Webframework: Webframework{},
	}
}

// Models is the type for this package. Note that any model that is included as a member
// in this type is available to us throughout the application, anywhere that the
// app variable is used, provided that the model is also added in the New function.
type Models struct {
	Language     Language
	Webframework Webframework
}

type Language struct {
	Name              string `json:"name"`
	HaveWorkedWithCnt int    `json:"have_worked_with_cnt"`
	WantToWorkWithCnt int    `json:"want_to_work_with_cnt"`
}

type Webframework struct {
	Name              string `json:"name"`
	HaveWorkedWithCnt int    `json:"have_worked_with_cnt"`
	WantToWorkWithCnt int    `json:"want_to_work_with_cnt"`
}

func (l *Language) GetLanguages() ([]*Language, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select name, have_worked_with_cnt, want_to_work_with_cnt from languages`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var languages []*Language
	for rows.Next() {
		var language Language
		err := rows.Scan(
			&language.Name,
			&language.HaveWorkedWithCnt,
			&language.WantToWorkWithCnt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		languages = append(languages, &language)
	}

	return languages, nil
}

func (w *Webframework) GetWebframeworks() ([]*Webframework, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select name, have_worked_with_cnt, want_to_work_with_cnt from frameworks`

	rows, err := db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var webframeworks []*Webframework
	for rows.Next() {
		var webframework Webframework
		err := rows.Scan(
			&webframework.Name,
			&webframework.HaveWorkedWithCnt,
			&webframework.WantToWorkWithCnt,
		)
		if err != nil {
			log.Println("Error scanning", err)
			return nil, err
		}

		webframeworks = append(webframeworks, &webframework)
	}

	return webframeworks, nil
}

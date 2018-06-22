package models

import (
	"database/sql"
)

//Database _
type Database struct {
	*sql.DB
}

//GetSnippet _
func (db *Database) GetSnippet(id int) (*Snippet, error) {

	stmnt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() AND id=?`

	row := db.QueryRow(stmnt, id)
	s := &Snippet{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)

	if err == sql.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return s, nil
}

// LatestSnippets _
func (db *Database) LatestSnippets() (Snippets, error) {
	stmnt := `SELECT id, title, content, created, expires WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT = 10`
	rows, err := db.Query(stmnt)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	snippets := Snippets{}

	for rows.Next() {
		s := &Snippet{}
		err := rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		snippets = append(snippets, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return snippets, nil
}

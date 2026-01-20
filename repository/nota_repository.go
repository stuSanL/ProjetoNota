package repository

import (
	"ProjetoNota/model"
	"database/sql"
	"errors"
	"fmt"
)

type NotaRepository struct {
	conn *sql.DB
}

func NewNotaRepository(conn *sql.DB) NotaRepository {
	return NotaRepository{
		conn: conn,
	}
}

func (nr *NotaRepository) GetNotas() ([]model.Nota, error) {
	query := "SELECT * FROM nota"
	rows, err := nr.conn.Query(query)
	if err != nil {
		fmt.Println(err)
		return []model.Nota{}, err
	}

	var notaList []model.Nota
	var notaObj model.Nota

	for rows.Next() {
		err = rows.Scan(
			&notaObj.Id,
			&notaObj.Data,
			&notaObj.Pontuacao,
			&notaObj.Texto)

		if err != nil {
			fmt.Println(err)
			return []model.Nota{}, err
		}

		notaList = append(notaList, notaObj)
	}

	return notaList, nil
}

func (nr *NotaRepository) AddNota(nota model.Nota) (int, error) {

	var id int

	query, err := nr.conn.Prepare("INSERT INTO nota ( data, pontuacao, texto) VALUES ($1, $2, $3) RETURNING id_nota")

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	err = query.QueryRow(nota.Data, nota.Pontuacao, nota.Texto).Scan(&id)

	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	_ = query.Close()
	return nota.Id, nil
}

func (nr *NotaRepository) GetNota(id int) (*model.Nota, error) {
	query, err := nr.conn.Prepare("SELECT * FROM nota WHERE id_nota = $1")

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	var nota model.Nota

	err = query.QueryRow(id).Scan(
		&nota.Id,
		&nota.Data,
		&nota.Pontuacao,
		&nota.Texto)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		fmt.Println(err)
		return nil, err
	}
	_ = query.Close()
	return &nota, nil
}

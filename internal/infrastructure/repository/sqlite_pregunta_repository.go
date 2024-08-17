package repository

import (
	"database/sql"
	"time"

	"github.com/matienzar/tarot-api/internal/core/domain"

	_ "github.com/mattn/go-sqlite3"
)

type SQLitePreguntaRepository struct {
	db *sql.DB
}

func NewSQLitePreguntaRepository(dataSourceName string) (*SQLitePreguntaRepository, error) {
	db, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}

	return &SQLitePreguntaRepository{db: db}, nil
}

func (r *SQLitePreguntaRepository) Save(pregunta *domain.Pregunta) (int, error) {
	stmt, err := r.db.Prepare("INSERT INTO preguntas (pregunta, respuesta, fecha_creacion) VALUES (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(pregunta.Texto, pregunta.Respuesta.Texto, time.Now())
	if err != nil {
		return 0, err
	}

	// Obtener el ID autogenerado
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *SQLitePreguntaRepository) FindAll() ([]*domain.Pregunta, error) {
	rows, err := r.db.Query("SELECT id, pregunta, respuesta, fecha_creacion FROM preguntas")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var preguntas []*domain.Pregunta
	for rows.Next() {
		var id int
		var pregunta, respuesta string
		var fechaCreacion time.Time

		err = rows.Scan(&id, &pregunta, &respuesta, &fechaCreacion)
		if err != nil {
			return nil, err
		}

		preguntas = append(preguntas, &domain.Pregunta{
			ID:        id,
			Texto:     pregunta,
			Respuesta: domain.Respuesta{Texto: respuesta},
			FechaHora: fechaCreacion,
		})
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return preguntas, nil
}

// Implement the FindByID method to fulfill the PreguntaRepository interface
func (r *SQLitePreguntaRepository) FindByID(id int) (*domain.Pregunta, error) {
	row := r.db.QueryRow("SELECT id, pregunta, respuesta, fecha_creacion FROM preguntas WHERE id = ?", id)

	var pregunta, respuesta string
	var fechaCreacion time.Time

	err := row.Scan(&id, &pregunta, &respuesta, &fechaCreacion)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No record found
		}
		return nil, err
	}

	return &domain.Pregunta{
		ID:        id,
		Texto:     pregunta,
		Respuesta: domain.Respuesta{Texto: respuesta},
		FechaHora: fechaCreacion,
	}, nil
}

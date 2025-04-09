package infraestructure

import (
	"actividad/src/config"
	"actividad/src/perfumes/domain"
	"fmt"
	"log"
)

type MySQL struct {
	conn *config.Conn_MySQL
}

var _ domain.IPerfume = (*MySQL)(nil) // Verifica que implementa IPerfume

// Constructor para MySQL
func NewMySQL() domain.IPerfume {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

// Implementación de SavePerfume
func (mysql *MySQL) SavePerfume(marca string, modelo string, precio float32) (int32, error) {
	query := "INSERT INTO perfume (marca, modelo, precio) VALUES (?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, marca, modelo, precio)
	if err != nil {
		return 0, fmt.Errorf("Error al ejecutar la consulta: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Error al obtener el ID del último registro insertado: %v", err)
	}
	return int32(id), nil
}


// Implementación de GetAll
func (mysql *MySQL) GetAll() ([]domain.Perfume, error) {
	query := "SELECT id, marca, modelo, precio FROM perfume"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar la consulta SELECT: %v", err)
	}
	defer rows.Close()

	var perfumes []domain.Perfume
	for rows.Next() {
		var perfume domain.Perfume
		if err := rows.Scan(&perfume.ID, &perfume.Marca, &perfume.Modelo, &perfume.Precio); err != nil {
			log.Printf("Error al escanear la fila: %v", err)
			continue
		}
		perfumes = append(perfumes, perfume)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error iterando sobre las filas: %v", err)
	}
	return perfumes, nil
}

// Implementación de UpdatePerfume
func (mysql *MySQL) UpdatePerfume(id int32, marca string, modelo string, precio float32) error {
	query := "UPDATE perfume SET marca = ?, modelo = ?, precio = ? WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, marca, modelo, precio, id)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta de actualización: %v", err)
	}
	return nil
}

// Implementación de DeletePerfume
func (mysql *MySQL) DeletePerfume(id int32) error {
	query := "DELETE FROM perfume WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta de eliminación: %v", err)
	}
	return nil
}

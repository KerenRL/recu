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

var _ domain.IPerfume = (*MySQL)(nil)

func NewMySQL() domain.IPerfume {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) SavePerfume(marca string, modelo string, precio float32) error {
	query := "INSERT INTO perfume (marca, modelo, precio) VALUES (?, ?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, marca, modelo, precio)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Perfume guardado correctamente: Marca: %s Modelo: %s - Precio: %.2f", marca, modelo, precio)
	} else {
		log.Println("[MySQL] - No se insertó ninguna fila")
	}
	return nil
}

func (mysql *MySQL) GetAll() ([]domain.Perfume, error) {
	query := "SELECT * FROM perfume"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar la consulta SELECT: %v", err)
	}
	defer rows.Close()

	var perfumes []domain.Perfume

	for rows.Next() {
		var perfume domain.Perfume
		if err := rows.Scan(&perfume.ID, &perfume.Marca, &perfume.Modelo, &perfume.Precio); err != nil {
			fmt.Printf("Error al escanear la fila: %v\n", err)
		}
		perfumes = append(perfumes, perfume)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("Error iterando sobre las filas: %v\n", err)
	}
	return perfumes, nil
}

func (mysql *MySQL) UpdatePerfume(id int32, marca string, modelo string, precio float32) error {
	query := "UPDATE perfume SET marca = ?, modelo = ?, precio = ? WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, marca, modelo, precio, id)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta de actualización: %v", err)
	}
	return nil
}

func (mysql *MySQL) DeletePerfume(id int32) error {
	query := "DELETE FROM perfume WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta de eliminación: %v", err)
	}
	return nil
}

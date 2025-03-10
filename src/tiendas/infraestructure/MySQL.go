package infraestructure

import (
	"actividad/src/config"
	"actividad/src/tiendas/domain"
	"fmt"
	"log"
)

type MySQL struct {
	conn *config.Conn_MySQL
}

var _ domain.ITienda = (*MySQL)(nil)

func NewMySQL() domain.ITienda {
	conn := config.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) SaveTienda(nombre string, direccion string) error {
	query := "INSERT INTO tienda (nombre, direccion) VALUES (?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, nombre, direccion)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Tienda guardada correctamente: Nombre: %s Dirección: %s", nombre, direccion)
	} else {
		log.Println("[MySQL] - No se insertó ninguna fila")
	}
	return nil
}

func (mysql *MySQL) GetAll() ([]domain.Tienda, error) {
	query := "SELECT * FROM tienda"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar la consulta SELECT: %v", err)
	}
	defer rows.Close()

	var tiendas []domain.Tienda

	for rows.Next() {
		var tienda domain.Tienda
		if err := rows.Scan(&tienda.ID, &tienda.Nombre, &tienda.Direccion); err != nil {
			fmt.Printf("Error al escanear la fila: %v\n", err)
		}
		tiendas = append(tiendas, tienda)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("Error iterando sobre las filas: %v\n", err)
	}
	return tiendas, nil
}

func (mysql *MySQL) UpdateTienda(id int32, nombre string, direccion string) error {
	query := "UPDATE tienda SET nombre = ?, direccion = ? WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, nombre, direccion, id)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta de actualización: %v", err)
	}
	return nil
}

func (mysql *MySQL) DeleteTienda(id int32) error {
	query := "DELETE FROM tienda WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("Error al ejecutar la consulta de eliminación: %v", err)
	}
	return nil
}

package infraestructure

import (
	"actividad/src/config"
	"actividad/src/tiendas/domain"
	"fmt"
)

type MySQL struct {
	conn *config.Conn_MySQL
}

var _ domain.ITienda = (*MySQL)(nil)

func NewMySQL() domain.ITienda {
	conn := config.GetDBPool()
	return &MySQL{conn: conn}
}

func (mysql *MySQL) SaveTienda(nombre string, ubicacion string) (int32, error) {
	query := "INSERT INTO tienda (nombre, ubicacion) VALUES (?, ?)"
	result, err := mysql.conn.ExecutePreparedQuery(query, nombre, ubicacion)
	if err != nil {
		return 0, fmt.Errorf("Error al ejecutar la consulta: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("Error al obtener el ID del último registro insertado: %v", err)
	}
	return int32(id), nil
}

func (mysql *MySQL) GetAll() ([]domain.Tienda, error) {
	query := "SELECT id, nombre, ubicacion FROM tienda"
	rows, err := mysql.conn.FetchRows(query)
	if err != nil {
		return nil, fmt.Errorf("Error al ejecutar la consulta SELECT: %v", err)
	}
	defer rows.Close()

	var tiendas []domain.Tienda
	for rows.Next() {
		var tienda domain.Tienda
		if err := rows.Scan(&tienda.ID, &tienda.Nombre, &tienda.Ubicacion); err != nil {
			fmt.Printf("Error al escanear la fila: %v\n", err)
		}
		tiendas = append(tiendas, tienda)
	}
	return tiendas, nil
}

func (mysql *MySQL) DeleteTienda(id int32) error {
	query := "DELETE FROM tienda WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		return fmt.Errorf("Error al eliminar la tienda: %v", err)
	}
	return nil
}

func (mysql *MySQL) UpdateTienda(id int32, nombre string, ubicacion string) error {
	if id <= 0 {
		return fmt.Errorf("ID de tienda inválido")
	}

	query := "UPDATE tienda SET nombre = ?, ubicacion = ? WHERE id = ?"
	_, err := mysql.conn.ExecutePreparedQuery(query, nombre, ubicacion, id)
	if err != nil {
		return fmt.Errorf("Error al actualizar la tienda: %v", err)
	}
	return nil
}

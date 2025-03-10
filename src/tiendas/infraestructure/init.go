package infraestructure

import (
    "actividad/src/config"
    "log"
)

func InitTienda() error {
    log.Println("Inicializando tiendas...")

    db, err := config.GetDBConnection()
    if err != nil {
        return err
    }
    defer db.Close()

    log.Println("Conexión a la base de datos para tiendas establecida correctamente")
    return nil
}
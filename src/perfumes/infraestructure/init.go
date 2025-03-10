package infraestructure

import (
    "actividad/src/config"
    "log"
)

func InitPerfume() error {
    log.Println("Inicializando perfumes...")

    db, err := config.GetDBConnection()
    if err != nil {
        return err
    }
    defer db.Close()

    log.Println("Conexión a la base de datos para perfumes establecida correctamente")
    return nil
}
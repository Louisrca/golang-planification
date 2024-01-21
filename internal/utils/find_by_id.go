package utils

import (
	"database/sql"
	"fmt"
	"log"
	"reflect"
)

// func FindByID(db *sql.DB, model interface{}, table string, id string) error {

// 	err:= db.QueryRow("SELECT id FROM ? WHERE id = ?",table, id ).Scan(&model)
// 	log.Println(id)

// 	if err!=nil{
// 		log.Println("l'ID est incorrecte")
// 		return err
// 	}
// 	log.Println("Tout est ok, la fonction fonctionne")
// 	return nil
// }



func FindByID(db *sql.DB, model interface{}, table string, id string) error {
    // Vérifier si le modèle est un pointeur vers une structure
    val := reflect.ValueOf(model)
    if val.Kind() != reflect.Ptr || val.Elem().Kind() != reflect.Struct {
        return fmt.Errorf("le modèle doit être un pointeur vers une structure")
    }

    // Construire la requête SQL
    query := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", table)

    // Préparer les champs pour la méthode Scan
    elem := val.Elem()
    fields := make([]interface{}, elem.NumField())
    for i := 0; i < elem.NumField(); i++ {
        fields[i] = elem.Field(i).Addr().Interface()
    }

    // Exécuter la requête
    err := db.QueryRow(query, id).Scan(fields...)
    if err != nil {
        if err == sql.ErrNoRows {
            log.Println("Aucun enregistrement trouvé avec l'ID fourni")
        } else {
            log.Printf("Erreur lors de la recherche par ID: %v", err)
        }
        return err
    }

    log.Println("Enregistrement trouvé avec succès")
    return nil
}



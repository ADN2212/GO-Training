package main

import (
	"database/sql"
	"fmt"
	//El _ siver para que el compilador no lanze un "unused import"
	_ "github.com/lib/pq" //Existen otros drivers, esto es usado como un implicit por el aql package
)

//postgres://admin:admin123@localhost:5433/orion_dev

const (
	host     = "localhost"
	port     = 5433
	user     = "admin"
	password = "admin123"
	dbname   = "orion_dev"
)

func main() {

	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		fmt.Print(err.Error())
		//panic(err)
		return
	}

	defer db.Close()

	err = db.Ping() //Veridica si existe una conexion a la BD y crea una de ser necesaria.
	if err != nil {
		fmt.Println("Failed to Ping")
		//panic(err)
		return
	}

	fmt.Println("Successfully connected!")

	//Ejemplo de una single row query:
	azuaId := "0198e6e8-6846-72e4-bf27-1d8564f0d64e"
	sqlStatemet := `SELECT name FROM cities WHERE id = $1;`
	row := db.QueryRow(sqlStatemet, azuaId)

	var name string
	queryError := row.Scan(&name) //Usa el puntero para volcar el valor hallado en la variable

	if queryError != nil {
		fmt.Printf("ERROR: %s", queryError.Error())
		return
	}

	fmt.Printf("The province name is: %s \n", name)

	//Ejemplo de query para mas de un dato:
	sqlStatemet2 := `SELECT id, name, province FROM cities WHERE id = $1;`

	type City struct {
		ID       string
		Name     string
		Province int8
	}

	var city City

	cityRow := db.QueryRow(sqlStatemet2, azuaId)//Este se usa cuando se espera un solo registro.
	queryError2 := cityRow.Scan(&city.ID, &city.Name, &city.Province) //Ha de obedecri el orden de los campos en la query ?? Si y los tipos deben coincidir.

	if queryError2 != nil {
		fmt.Printf("ERROR: %s", queryError2.Error())
		return
	}

	fmt.Println(city)

	//Query para multiples registros:
	var cities []City

	citiesRows, err3 := db.Query(`SELECT id, name, province FROM cities LIMIT 10`)

	if err3 != nil {
		println(err3.Error())
		return
	}

	var currCity City

	defer citiesRows.Close() //Cerra la conexion con la BD

	for citiesRows.Next() { //Esto es como un iterator, retorna un boleano
		singleScanErr := citiesRows.Scan(&currCity.ID, &currCity.Name, &currCity.Province)
		if singleScanErr != nil {
			println(singleScanErr.Error())
			return //Lo usual es usar panic, pero esto lo hago para deterner el hilo de ejecucion
		}
		cities = append(cities, currCity)
	}

	println("------------------------------------------------------")
	for i := range cities {
		fmt.Println(cities[i])
	}
	println("------------------------------------------------------")

	//Ejemplo de UPDATE query:
	updateQuery := `UPDATE cities SET name = $2 WHERE id = $1;`

	//Este metodo se usa cuando no se quiere que se retornen rows.
	res, updateError := db.Exec(updateQuery, "28feb055-c731-4af8-9a02-fceca44e751a", "Konoha")

	if updateError != nil {
		println(updateError.Error())
		return
	}

	//Se podria retornar un 404 en caso de que 0 rows was afected
	rowAfected, err10 := res.RowsAffected()

	if err10 != nil {
		print(err10.Error())
		return
	}

	fmt.Printf("%v rows were affected \n", rowAfected)

	//Ejemplo de delete query:
	deleteQuery := `DELETE from cities WHERE id = $1;`

	deleteRes, deleteError := db.Exec(deleteQuery, "28feb055-c731-4af8-9a02-fceca44e751a")

	if deleteError != nil {
		fmt.Println(deleteError.Error())
		return
	}

	//Se podria retornar not found en caso de que no se afecte (elimine) ninguna row
	deletedRows, err11 := deleteRes.RowsAffected()

	if err11 != nil {
		print(err11.Error())
		return
	}

	fmt.Printf("%v rows were deleted \n", deletedRows)

	//Ejemplo de INSERT query:
	insertQuery := `INSERT INTO cities (id, name, province) VALUES ($1, $2, $3)` //Aunque se usen estas templetes el sql package no permite injecciones SQL

	_, insertError := db.Exec(insertQuery, "28feb055-c731-4af8-9a02-fceca44e751a", "Gothan", 1)

	if insertError != nil {
		fmt.Println(insertError.Error())
		return
	}

	print("Succesful insertion")

}

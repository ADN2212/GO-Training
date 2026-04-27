package main

import (
	"context"
	"database/sql"
	"log"
	"reflect"

	//"github.com/jackc/pgx/v5"
	//"github.com/jackc/pgx/v5/pgtype"
	_ "github.com/lib/pq"
	"tutorial.sqlc.dev/app/tutorial" //Ahora nuestro tutorail es como un package al que podemos acceder desde cualquier lugar.
)

func run() error {
	
	ctx := context.Background()

	//En el tuturial se uso pgx pero con este diver tambien se puede
	conn, err := sql.Open("postgres", "host=localhost user=adn password=123456 dbname=magratea_dev port=5432 sslmode=disable")
	if err != nil {
		return err
	}
	defer conn.Close()

	queries := tutorial.New(conn)//el argumento que resive el metodo new ha de ser una conexion a la base de datos

	//Ver name GetAllusers en querys.sql
	users, err := queries.GetAllUsers(ctx)
	if err != nil {
		return err
	}

	for _, user := range users {
		log.Println(user.Username.String)
	} 

	//Ver name Adduser en querys.sql 
	insertedUser, err := queries.AddUser(ctx, tutorial.AddUserParams{
		Username: sql.NullString{"Maneka", true},//Que son estos tipo ???
		Password: sql.NullString{"123456", true},
		Language: sql.NullInt16{1, true},

	})
	if err != nil {
		return err
	}
	log.Println(insertedUser)

	//Porque estas funciones necesitan resivir un contexto como argumento ???
	fetchedUser, err := queries.GetUserByUsername(ctx, insertedUser.Username)
	if err != nil {
		return err
	}

	// prints true
	log.Println(reflect.DeepEqual(insertedUser, fetchedUser))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

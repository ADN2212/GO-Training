package main


import (
	"database/sql"
	"fmt"
	//El _ siver para que el compilador no lanze un "unused import"
	_ "github.com/lib/pq" //Existen otros drivers, esto es usado como un implicit por el aql package
)


type TipoAve int

const (
	Gallo TipoAve = iota 
	Gallina           
)

type Ave struct {
	Id int
	Name string
	Tipo TipoAve
}

//Tabla que sirve definir las relaciones entre aves
type AveConnection struct {
	Id int
	PadreId int
	MadreId int
	PadroteId int
}


//juanadonisnunezcollado:manetha1126@localhost:5432

const (
	host     = "localhost"
	port     = 5432
	user     = "juanadonisnunezcollado"
	password = "manetha1126"
	dbname   = "GallosDB"
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
	

}

//SQL para almacenar el grafo en forma de lista de abjacencia.:

// CREATE TABLE aves (
//   id INTEGER PRIMARY KEY,
//   type SMALLINT NOT NULL,--un enum 1 gallo dos gallina
//   name TEXT NOT NULL
// );

// CREATE TABLE aves_connections (
// 	id INTEGER PRIMARY KEY,
// 	aveId INTEGER NOT NULL REFERENCES aves(id), 
// 	padreId INTEGER REFERENCES aves(id),
//   	madreId INTEGER REFERENCES aves(id)
// );

func deleteAvedb(db *sql.DB, id int) {
	//El delete deberi ser soft 
}

func createAve(db *sql.DB, ave Ave, padreId, madreId *int) {
	//Crea el ave
	//Crea una conexion
}

func updateAveParents(db *sql.DB, id, newPadreId, newMadreId int) {
	//Vericar que el padre exista O(n)
	//Verificar que la madre exista O(n)
	//Verificar que ninguno sea descendiente del ave actual O(???)
	//Update O(1)
}

//Deberia ser capaz de mostrar todas las aves en el subtree 
func ViewDescendants(db *sql.DB, aveId int, padroteId int) {

	//Buscar el minimo de todos los ids decendientes del ave
	//fil


}


//func (ave *Ave)showTree() {}


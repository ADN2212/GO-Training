package main

import (
	"context"
	"encoding/csv"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"golang.org/x/text/encoding/charmap"
	"golang.org/x/text/transform"
	"io"
	"os"
	"time"
)

//Todo esto se hizo con esta tabla
// CREATE TABLE taxpayers (
//     rnc                   TEXT,
//     business_name         TEXT,
//     economic_activity     TEXT,
//     date_operations_began TEXT,
//     status                TEXT,
//     payment_regime        TEXT
// );

// Este tipo debe implementar los metodos, Next, Values, y Err para que la
type RowSrc struct {
	cr     *csv.Reader
	values []any
	err    error
}

// Esto funciona como un iterador
func (r *RowSrc) Next() bool {
	record, err := r.cr.Read() //Esto es como un next para el csv, lee una row cada vez que se llama
	if errors.Is(err, io.EOF) {
		return false
	}

	if err != nil {
		r.err = err
		return false
	}

	r.values = make([]any, 6)
	r.values[0] = record[0]
	r.values[1] = record[1]
	r.values[2] = record[2]
	r.values[3] = record[3]
	r.values[4] = record[4]
	r.values[5] = record[5]

	return true
}

func (r *RowSrc) Values() ([]any, error) {
	return r.values, r.err
}

func (r *RowSrc) Err() error {
	return r.err
}

const dns = "host=localhost port=5432 user=juanadonisnunezcollado password=123456 dbname=rnc_test sslmode=disable"

var colunmNames = []string{"rnc", "business_name", "economic_activity", "date_operations_began", "status", "payment_regime"}

func main() {
	start := time.Now()
	fmt.Println("Abriendo al chivo")
	r, err := openCSV("RNC_Contribuyentes_Actualizado_01_Ago_2026.csv")
	if err != nil {
		panic(err.Error())
	}

	dbCtx := context.Background()

	conn, err := pgx.Connect(dbCtx, dns)
	if err != nil {
		panic(err.Error())
	}

	defer func() {
		fmt.Println("Cerrando al chivo : /")
		r.Close()
		conn.Close(dbCtx)
	}()

	count, err := insert(conn, r)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%v filas insertadas en %v segundos \n", count, time.Since(start))

}

func insert(conn *pgx.Conn, f io.ReadCloser) (int64, error) {
	rowSrc := RowSrc{
		cr: csv.NewReader(f),
	}

	count, err := conn.CopyFrom(context.Background(), pgx.Identifier{"taxpayers"}, colunmNames, &rowSrc)
	if err != nil {
		return 0, fmt.Errorf("conn.CopyFrom %w", err)
	}
	return count, nil
}

func openCSV(path string) (io.ReadCloser, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	// The source file is Windows-1252; decode to UTF-8 on the fly.
	decoded := transform.NewReader(f, charmap.Windows1252.NewDecoder())

	return struct {
		io.Reader
		io.Closer
	}{decoded, f}, nil
}

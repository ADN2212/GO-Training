package main

import "fmt"

func main() {
    // Initialize a map for the integer values
    ints := map[string]int64{
        "first":  34,
        "second": 12,
    }

    // Initialize a map for the float values
    floats := map[string]float64{
        "first":  35.98,
        "second": 26.99,
    }

    fmt.Printf("Non-Generic Sums: %v and %v\n",
        SumInts(ints),
        SumFloats(floats))

	fmt.Printf("Generic Sums: %v and %v\n",
	//Vease como los tipos genericos se "resiven" como parametros de la funcion dentro de los corchetes.
	//Ojo, el compilador puede inferir los typos de los argumentos.
    SumIntsOrFloats[string, int64](ints),
    SumIntsOrFloats[string, float64](floats))

	fmt.Printf("Generic Sums with Constraint: %v and %v\n",
    SumNumbers(ints),//Vease como no es requerido pasar los parameter types.
    SumNumbers(floats))	

}


//Estas dos funciones ejecutan la misma logica 
//pero los mapas que resiven como inputs son distintos en el tipo que tienen como valor
//Esto se resuelve con un generic type.
// SumInts adds together the values of m.
func SumInts(m map[string]int64) int64 {
    var s int64
    for _, v := range m {
        s += v
    }
    return s
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
    var s float64
    for _, v := range m {
        s += v
    }
    return s
}
//-----------------------------------------------------

// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
//Que es comparable ?
func SumIntsOrFloats[K comparable, V int64 | float64](m map[K]V) V {
	var s V
    for _, v := range m {
		s += v
    }
    return s
}

type Number interface {//Esta interface se puede ver como un alias para los
	int64 | float64
}

//
func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
    for _, v := range m {
		s += v
    }
    return s
}

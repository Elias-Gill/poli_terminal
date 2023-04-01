package excelParser

import (
	"strconv"

	"github.com/xuri/excelize/v2"
)

type materia struct {
	Nombre   string
	Semestre int
	Seccion  string
	Profesor string
	Parcial1 string
	Parcial2 string
	Final1   string
	Final2   string
}

type rowLimit struct {
	inicio int
	fin    int
}

// determinar donde termina la lista de materias
func getValidRows(mat [][]string) rowLimit {
	res := rowLimit{inicio: 1, fin: 1}
	// detemrinar el inicio
	for i := range mat {
		if mat[0][i] == "1" {
			res.inicio = i
			break
		}
	}
	// detemrinar el inicio
	for i := res.inicio; i < len(mat[0]); i++ {
		if mat[0][i] == "" {
			res.fin = i
			break
		}
	}
	return res
}

// retorna la lista de materias de la carrera con fechas de finales, semestre,
// parciales, profesor y seccion
func GetListaMaterias(fname string) map[string]materia {
	file, err := excelize.OpenFile(fname)
	if err != nil {
		panic("no se puede abrir el archivo")
	}

	// abrir el archivo excel
	cols, err := file.GetCols(file.GetSheetName(6))
	if err != nil {
		panic("no se pueden traer las columnas")
	}

	// determinar donde empieza la lista de materias
	validRows := getValidRows(cols)
	asignaturas := make(map[string]materia)

	// comenzar a cargar la lista de asignaturas
	for row := validRows.inicio; row < validRows.fin+1; row++ {
		s, _ := strconv.Atoi(cols[3][row])
		asignaturas[string(cols[2][row])] = materia{
			Nombre:   string(cols[2][row]),
			Semestre: s,
			Seccion:  string(cols[9][row]),
			Profesor: string(cols[13][row]) + " " + string(cols[12][row]),
			Parcial1: string(cols[15][row]) + " " + string(cols[16][row]) + " " + string(cols[17][row]),
			Parcial2: string(cols[18][row]) + " " + string(cols[19][row]) + " " + string(cols[20][row]),
			Final1:   string(cols[21][row]) + " " + string(cols[22][row]) + " " + string(cols[23][row]),
			Final2:   string(cols[24][row]) + " " + string(cols[25][row]) + " " + string(cols[26][row]),
		}
	}
	return asignaturas
}

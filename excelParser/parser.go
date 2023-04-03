package excelParser

import (
	"strconv"

	"github.com/xuri/excelize/v2"
)

type Materia struct {
	Nombre   string `json:"nombre"`
	Semestre int    `json:"semestre"`
	Seccion  string `json:"seccion"`
	Profesor string `json:"profesor"`
	Parcial1 string `json:"parcial_1"`
	Parcial2 string `json:"parcial_2"`
	Final1   string `json:"final_1"`
	Final2   string `json:"final_2"`
}

type rowLimit struct {
	inicio int
	fin    int
}

// Determinar donde comeinza y termina la lista de materias
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
func GetListaMaterias(fname string) ([]Materia, error) {
	// abrir el archivo excel
	file, err := excelize.OpenFile(fname)
	if err != nil {
		return nil, err
	}
	// parsear las columnas
	cols, err := file.GetCols(file.GetSheetName(6))
	if err != nil {
		return nil, err
	}

	// determinar donde empieza la lista de materias
	validRows := getValidRows(cols)
	asignaturas := []Materia{}

	// comenzar a cargar la lista de asignaturas
	for row := validRows.inicio; row < validRows.fin+1; row++ {
		s, _ := strconv.Atoi(cols[3][row])
		asignaturas = append(asignaturas, Materia{
			Nombre:   string(cols[2][row]),
			Semestre: s,
			Seccion:  string(cols[9][row]),
			Profesor: string(cols[13][row]) + " " + string(cols[12][row]),
			Parcial1: string(cols[15][row]) + " " + string(cols[16][row]) + " " + string(cols[17][row]),
			Parcial2: string(cols[18][row]) + " " + string(cols[19][row]) + " " + string(cols[20][row]),
			Final1:   string(cols[21][row]) + " " + string(cols[22][row]) + " " + string(cols[23][row]),
			Final2:   string(cols[24][row]) + " " + string(cols[25][row]) + " " + string(cols[26][row]),
		})
	}
	return asignaturas, nil
}

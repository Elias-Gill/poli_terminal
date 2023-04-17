package configManager

import (
	"encoding/json"
	"os"
	"runtime"

	ep "github.com/elias-gill/poli_terminal/excelParser"
)

// Configuracion general del usuario. Estos datos son almacenados en el
// archivo de configuracion
type Configurations struct {
	ExcelFile       string       `json:"file_horario"`
	MateriasUsuario []*ep.Materia `json:"lista_materias"`
	MateriasExcel   []*ep.Materia `json:"lista_excel"`
	Sheet           int          `json:"sheet_number"`
}

var usersConfig = LoadUserConfig()
var configPaths = searchConfigFiles()

type paths struct {
	path string
	file string
}

// Retorna le configuracion actual del usuario
func GetUserConfig() *Configurations {
	return usersConfig
}

// Cambia las materias de la configuracion del usuario
func (c *Configurations) ChangeMateriasUsuario(m []*ep.Materia) {
	c.MateriasUsuario = m
}

// Cambia el archivo excel y lo parsea de manera asincrona
func (c *Configurations) ChangeExcelFile(f string) error {
	aux, err := ep.Parse(f, c.Sheet)
	if err != nil {
		return err
	}
	c.ExcelFile = f
	c.MateriasExcel = aux
	return nil
}

// Parsear la configuracion del usuario y la guarda en memoria.
// Pensada para ser llamada una sola vez durante el inicio de la app, dentro de una
// GoRoutine TODO: asincronia
func LoadUserConfig() *Configurations {
	// asegurarse que el archivo exista
	ensureConfigExistence()
	file, _ := os.Open(configPaths.file)
	defer file.Close()
	// parsear
	var config Configurations
	json.NewDecoder(file).Decode(&config)
	// cargar las materias del excel TODO: cambio de carrera (sheet)
	config.MateriasExcel, _ = ep.Parse(config.ExcelFile, config.Sheet)
	return &config
}

// Escribe la nueva configuracion del usuario en un archivo json
func (c Configurations) WriteUserConfig() {
	data, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		panic("No se pudo generar la configuracion. Lo sentimos, algo salio mal")
	}
	err = os.WriteFile(configPaths.file, data, 0644)
	if err != nil {
		panic("No se pudo escribir en el archivo de configuracion.\nAsegurese de tener los permisos necesarios")
	}
}

// Se asegura de que los archivos de configuracion esten creados,
// los crea de ser necesario. Panic cuando no se puede crear el archivo de configuracion
func ensureConfigExistence() {
	// Crear la carpeta
	if _, err := os.Stat(configPaths.path); os.IsExist(err) {
		err := os.Mkdir(configPaths.path, os.ModeDir)
		if err != nil {
			panic("No se pudo crear la carpeta para la configuracion. \nAsegurate de tener los permisos adecuados")
		}
	}

	// Crear el archivo
	if _, err := os.Stat(configPaths.file); os.IsExist(err) {
		_, err = os.Create(configPaths.file)
		if err != nil {
			panic("No se pudo crear el archivo de configuracion del horario. \nAsegurate de tener los permisos adecuados")
		}
	}
}

// Determinar la ubicacion del archivo de configuracion en runtime dependiendo del OS.
func searchConfigFiles() paths {
	osys := runtime.GOOS
	userPaht, err := os.UserHomeDir()
	if err != nil {
		panic("Cannot determine your home directory, somehting goes wrong")
	}

	if osys == "windows" {
		p := userPaht + "AppData/Local/politerm/"
		f := p + "config.json"
		return paths{file: f, path: p}
	}

	p := userPaht + "/.config/politerm/"
	f := p + "config.json"
	return paths{file: f, path: p}
}

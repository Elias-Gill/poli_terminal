package configManager

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"

	ep "github.com/elias-gill/poli_terminal/excelParser"
)

// Configuracion general del usuario. Estos datos son almacenados en el
// archivo de configuracion
type Configurations struct {
	ExcelFile       string        `json:"excel_file"`
	MateriasUsuario []*ep.Materia `json:"lista_materias"`
	MateriasExcel   []*ep.Materia `json:"lista_excel"`
	Sheet           int           `json:"sheet_number"`
}

// TODO: no estaria proporcionar metodos y que la config sea inmutable desde afuera (mandar una copia)
var configPaths = searchConfigFiles()
var usersConfig = loadUserConfig()

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

// Cambia el archivo excel y lo parsea
func (c *Configurations) ChangeExcelFile(f string) error {
	aux, err := ep.Parse(f, c.Sheet)
	if err != nil {
		panic("no se puede abrir el excel")
	}
	c.ExcelFile = f
	c.MateriasExcel = aux
	return nil
}

// Parsear la configuracion del usuario y la guarda en memoria. Solo se llama una vez al inicio
// del programa
func loadUserConfig() *Configurations {
	// asegurarse que el archivo exista
	ensureConfigExistence()
    file, _ := os.Open(configPaths.file) // INFO: no hace falta revisar el error
	defer file.Close()
	// parsear
	var config Configurations
	json.NewDecoder(file).Decode(&config)

	// seleccionar una hoja por defecto (IIN)
	if config.Sheet == 0 {
		config.Sheet = 6
	}

	// revisar si el excel ya no esta "pre parseado"
	if len(config.MateriasExcel) == 0 {
		// cargar las materias del excel
		aux, err := ep.Parse(config.ExcelFile, config.Sheet)
		if err == nil {
			config.MateriasExcel = aux
			config.WriteUserConfig()
		}
	}
	return &config
}

// Escribe la nueva configuracion del usuario en un archivo json
func (c Configurations) WriteUserConfig() {
	data, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		fmt.Print("No se pudo generar la configuracion. Lo sentimos, algo salio mal")
		os.Exit(1)
	}
	err = os.WriteFile(configPaths.file, data, 0644)
	if err != nil {
		fmt.Print("No se pudo escribir en el archivo de configuracion.\nAsegurese de tener los permisos necesarios")
		os.Exit(1)
	}
}

// Se asegura de que los archivos de configuracion esten creados,
// los crea de ser necesario. Panic cuando no se puede crear el archivo de configuracion
func ensureConfigExistence() {
	// Crear la carpeta
	if _, err := os.Stat(configPaths.path); os.IsNotExist(err) {
		err := os.Mkdir(configPaths.path, 0777)
		if err != nil {
			fmt.Print("No se pudo crear la carpeta para la configuracion. \nAsegurate de tener los permisos adecuados")
			os.Exit(1)
		}
	}

	// Crear el archivo
	if _, err := os.Stat(configPaths.file); os.IsNotExist(err) {
		_, err = os.Create(configPaths.file)
		if err != nil {
			fmt.Print("No se pudo crear el archivo de configuracion. \nAsegurate de tener los permisos adecuados")
			os.Exit(1)
		}
	}
}

// Determina la ubicacion del archivo de configuracion.
func searchConfigFiles() paths {
	osys := runtime.GOOS
	userPaht, err := os.UserHomeDir()
	if err != nil {
		fmt.Print("Cannot determine your home directory, somehting goes wrong")
		os.Exit(1)
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

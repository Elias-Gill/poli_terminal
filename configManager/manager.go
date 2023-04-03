package configManager

import (
	"encoding/json"
	"os"
	"runtime"

	"github.com/elias-gill/poli_terminal/excelParser"
)

// Configuracion general del usuario. Estos datos son almacenados en el
// archivo de configuracion
type Configurations struct {
	FHorario string                `json:"file_horario"`
	Materias []excelParser.Materia `json:"lista_materias"`
}

var rutas = detConfigFile()

type conf struct {
	path string
	file string
}

// Determinar la ubicacion del archivo de configuracion en runtime dependiendo del OS.
func detConfigFile() conf {
	osys := runtime.GOOS
	userPaht, err := os.UserHomeDir()
	if err != nil {
		panic("Cannot determine your home directory, somehting goes wrong")
	}

	if osys == "windows" {
		p := userPaht + "AppData/Local/politerm/"
		f := p + "config.json"
		return conf{file: f, path: p}
	}

	p := userPaht + "/.config/politerm/"
	f := p + "config.json"
	return conf{file: f, path: p}
}

// Parsear la configuracion del usuario
func GetUserConfig() Configurations {
	// asegurarse que el archivo exista
	ensureExistence()
	file, _ := os.Open(rutas.file)
	defer file.Close()
	// parsear
	var config Configurations
	json.NewDecoder(file).Decode(&config)
	return config
}

// Escribe la nueva configuracion del usuario en un archivo json
func WriteUserConfig(c Configurations) {
	data, err := json.MarshalIndent(c, "", " ")
	if err != nil {
		panic("No se pudo generar la configuracion. Lo sentimos, algo salio mal")
	}
	err = os.WriteFile(rutas.file, data, 0644)
	if err != nil {
		panic("No se pudo escribir en el archivo de configuracion.\nAsegurese de tener los permisos necesarios")
	}
}

// Se asegura de que los archivos de configuracion esten creados, si es que no lo estan
// los crea. Panic cuando no se puede crear el archivo de configuracion
func ensureExistence() {
	// Crear la carpeta
	if _, err := os.Stat(rutas.path); os.IsExist(err) {
		err := os.Mkdir(rutas.path, os.ModeDir)
		if err != nil {
			panic("No se pudo crear la carpeta para la configuracion. \nAsegurate de tener los permisos adecuados")
		}
	}

	// Crear el archivo
	if _, err := os.Stat(rutas.file); os.IsExist(err) {
		_, err = os.Create(rutas.file)
		if err != nil {
			panic("No se pudo crear el archivo de configuracion del horario. \nAsegurate de tener los permisos adecuados")
		}
	}
}

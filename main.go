package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	SearchFiles(os.Args[1])
	fmt.Println("Caracteres reemplazados y el resultado se ha guardado en el archivo de salida.")
}

func SearchFiles(d string) {
	dir, err := os.ReadDir(d)
	if err != nil {
		log.Fatal(err)
	}
	for _, de := range dir {
		position := d + "/" + de.Name()
		if de.IsDir() {
			if de.Name() != ".git" {
				SearchFiles(position)
			}
		} else {
			// Carácter a reemplazar y carácter de reemplazo
			oldP := os.Args[2]
			newP := os.Args[3]

			// Leer el contenido del archivo de entrada
			content, err := ioutil.ReadFile(position)
			if err != nil {
				fmt.Println("Error al leer el archivo de entrada:", err)
				return
			}

			// Realizar la sustitución de caracteres en el contenido
			modifiedContent := strings.ReplaceAll(string(content), oldP, newP)

			// Crear o abrir el archivo de salida en modo escritura
			output, err := os.Create(position)
			if err != nil {
				fmt.Println("Error al crear o abrir el archivo de salida:", err)
				return
			}
			defer output.Close()

			// Escribir el contenido modificado en el archivo de salida
			_, err = output.WriteString(modifiedContent)
			if err != nil {
				fmt.Println("Error al escribir en el archivo de salida:", err)
				return
			}
		}
	}
}

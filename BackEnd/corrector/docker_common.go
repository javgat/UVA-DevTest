// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package corrector executes Pruebas for the Code QuestionAnswers
package corrector

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
	"uva-devtest/persistence/dao"
)

const docker_scripts = "./docker/"

func createError(out []byte, err error) error {
	if err == nil {
		return nil
	}
	return errors.New(err.Error() + ": " + string(out))
}

func prepareFolder(sharedDir string) error {
	out, err := exec.Command(docker_scripts+"prepare_folder.sh", sharedDir).Output()
	if err != nil {
		log.Println("Error en prepareFolder()")
	}
	return createError(out, err)
}

func cleanFolder(sharedDir string) error {
	out, err := exec.Command(docker_scripts+"clean_folder.sh", sharedDir).Output()
	if err != nil {
		log.Println("Error en cleanFolder()")
	}
	return createError(out, err)
}

func execDocker(sharedDir string) error {
	out, err := exec.Command(docker_scripts+"exec_container.sh", sharedDir).Output()
	if err != nil {
		log.Println("Error en execDocker()")
	}
	return createError(out, err)
}

func writePruebaInput(p *dao.Prueba, sharedDir string) error {
	err := ioutil.WriteFile(sharedDir+"inputs/"+strconv.FormatInt(p.ID, 10), []byte(*p.Entrada), os.ModePerm)
	return err
}

func writeCodeProgram(resp string, sharedDir string) error {
	err := ioutil.WriteFile(sharedDir+"code.cpp", []byte(resp), os.ModePerm)
	return err
}

func readPruebaInput(p *dao.Prueba, aid int64, qid int64, sharedDir string) (*dao.Ejecucion, error) {
	data, err := ioutil.ReadFile(sharedDir + "outputs/" + strconv.FormatInt(p.ID, 10))
	data2, err2 := ioutil.ReadFile(sharedDir + "errors/" + strconv.FormatInt(p.ID, 10))
	if err != nil && err2 != nil {
		return nil, err
	}
	dataST := string(data)
	data2ST := string(data2)
	var ej = &dao.Ejecucion{
		Pruebaid:          &p.ID,
		RespuestaExamenid: &aid,
		Preguntaid:        &qid,
		Estado:            &data2ST,
		SalidaReal:        &dataST,
	}

	return ej, err
}

func readCompilationErrors(sharedDir string) (string, error) {
	data, err := ioutil.ReadFile(sharedDir + "erroresCompilacion.txt")
	if err == nil {
		return string(data), nil
	}
	return "", nil
}

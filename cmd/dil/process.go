/*
Sniperkit-Bot
- Status: analyzed
*/

package main

import (
	"io/ioutil"
	"log"

	"github.com/senseyeio/diligent"
)

func runDep(deper diligent.Deper, reper diligent.Reporter, filePath string) {
	fileBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	deps, err := deper.Dependencies(fileBytes)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err = reper.Report(deps); err != nil {
		log.Fatal(err.Error())
	}
}

package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/bitrise-io/go-utils/log"
)

const InitFileName = "_tmp_init.gradle"

func main() {
	createInitFile()
	runGradle()
	deleteInitFile()

	os.Exit(0)
}

func createInitFile() {
	f, err := os.Create(InitFileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	f.WriteString(`gradle.buildFinished() {
		def isParallelExecutionEnabled = rootProject.properties.containsKey('org.gradle.parallel') && rootProject.properties.get('org.gradle.parallel').toLowerCase() == 'true'
		println isParallelExecutionEnabled
	}`)
}

func deleteInitFile() {
	os.Remove(InitFileName)
}

func runGradle() {
	cmd := "./gradlew --init-script _tmp_init.gradle -q | tail -n 1"
	out, _ := exec.Command("bash", "-c", cmd).Output()
	if string(out) == "true\n" {
		log.Donef("✓ Parallel Execution")
	} else {
		log.Warnf("! Parallel Execution")
	}
}

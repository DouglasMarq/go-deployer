package worker

import (
	"fmt"
	"log"
	"os/exec"

	"github.com/DouglasMarq/go-deployer/constants"
)

func Execute(scriptType constants.ScriptType, fullPath string) error {
	cmd := exec.Command(string(scriptType), fullPath)

	output, err := cmd.CombinedOutput()
	fmt.Println(output)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

package main

import (
	"goci/common_library/logger"
	"os"
	"os/exec"

	"go.uber.org/zap"
	"gopkg.in/yaml.v3"
)
 
type Pipeline struct{
	Steps []Step `yaml:"steps"`
}
type Step struct{
	Name string `yaml:"name"`
	Run string `yaml:"run"`
}


func main() {
	// if len(os.Args) != 3 {
	// 	fmt.Println("[Error]\ncorrect command: goci run <filename.yaml>")
	// 	return
	// }

	// filename := os.Args[2]
	// file, err := os.ReadFile(filename)
	// if err != nil{
	// 	panic(err)
	// }


	logger := logger.NewLogger("pipeline.log")
	var pipe Pipeline
	file, _ := os.ReadFile("pipeline.yaml")
	err := yaml.Unmarshal(file, &pipe)
	
	if err != nil {
		logger.Fatal("", zap.Error(err))
	}
	for _, command := range pipe.Steps{
		logger.Info("Running %s:\n", zap.String("step", command.Name))
		cmd := exec.Command("sh", "-c", command.Run)
		output, err := cmd.CombinedOutput()
		if err != nil{
			logger.Err("Command "+command.Name+" failed", zap.Error(err))
		}
		logger.Info(string(output))
	}
	
	


	
}

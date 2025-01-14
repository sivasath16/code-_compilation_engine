package backend

import (
	"fmt"
	"log"
	"os/exec"
)

// ExecuteTask runs the given code in a Docker container based on the language
func ExecuteTask(code, language string) (string, error) {
	var image string

	// Map the language to its Docker image
	switch language {
	case "javascript":
		image = "code-executor-javascript"
	case "python":
		image = "code-executor-python"
	case "java":
		image = "code-executor-java"
	case "cpp":
		image = "code-executor-cpp"
	default:
		return "", fmt.Errorf("unsupported language: %s", language)
	}

	// Log the code and language
	log.Printf("Executing code: %s, Language: %s", code, language)

	// Run the Docker container
	cmd := exec.Command("docker", "run", "--rm", "-e", fmt.Sprintf("CODE=%s", code), image)

	// Log the exact command being run
	log.Printf("Running command: %s", cmd.String())

	// Execute the command
	output, err := cmd.CombinedOutput()
	if err != nil {
		// Log runtime errors or combined output
		log.Printf("Error executing Docker container: %v, Output: %s", err, output)
		return string(output), err
	}

	// Log the Docker output
	log.Printf("Docker output: %s", output)
	return string(output), nil
}

package env

import (
	"strings"

	log "github.com/newrelic/newrelic-diagnostics-cli/logger"
	"github.com/newrelic/newrelic-diagnostics-cli/tasks"
)

// NodeEnvYarnVersion - This struct defines the Ruby version
type NodeEnvYarnVersion struct {
	// cmdExecutor      tasks.CmdExecFunc
	// npmVersionGetter getNpmVersionFunc
}

// Identifier - This returns the Category, Subcategory and Name of each task
func (p NodeEnvYarnVersion) Identifier() tasks.Identifier {
	return tasks.IdentifierFromString("Node/Env/YarnVersion")
}

// Explain - Returns the help text for each individual task
func (p NodeEnvYarnVersion) Explain() string {
	return "Determine NPM version"
}

// Dependencies - Returns the dependencies for each task.
func (p NodeEnvYarnVersion) Dependencies() []string {
	return []string{
		"Node/Env/Version",
	}
}

// Execute - The core work within each task
func (p NodeEnvYarnVersion) Execute(options tasks.Options, upstream map[string]tasks.Result) tasks.Result {

	if upstream["Node/Env/Version"].Status != tasks.Info {
		return tasks.Result{
			Status:  tasks.None,
			Summary: "Node.js is not installed. This task did not run",
		}
	}

	// yarnVersion, err := yarnVersionGetter()
	return tasks.Result{
		Status:  tasks.Info,
		Summary: "We found yarn version: ",
	}
}

func getYarnVersion() (string, error) {
	version, err := tasks.CmdExecutor("npm", "-v")
	if err != nil {
		log.Debug("Error running npm -v ", version)
		return "", err
	}
	return strings.TrimSpace(string(version)), nil
}
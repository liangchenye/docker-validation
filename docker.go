package docker_validation

import (
	"errors"
	"os/exec"
)

type ConformanceContainer struct {
	Name string
}

func NewConformanceContainer(name string) (cc ConformanceContainer, err error) {
	if name != "docker" {
		err = errors.New("Only 'docker' supported")
	} else {
		cc.Name = name
	}

	return
}

func (cc *ConformanceContainer) IsAlive() bool {
	cmd := exec.Command("docker", "ps")
	if _, err := cmd.CombinedOutput(); err != nil {
		return false
	}

	return true
}

func (cc *ConformanceContainer) Run(sync bool, image string, args []string) error {
	argsNew := make([]string, len(args)+2)
	argsNew[0] = "run"
	argsNew[1] = image
	for i, arg := range args {
		argsNew[i+2] = arg
	}

	cmd := exec.Command("docker", argsNew...)
	err := cmd.Start()
	if sync {
		err = cmd.Wait()
	}

	return err
}

func (cc *ConformanceContainer) Start() error {
	cmd := exec.Command("service", cc.Name, "start")
	_, err := cmd.CombinedOutput()

	return err
}

func (cc *ConformanceContainer) Restart() error {
	cmd := exec.Command("service", cc.Name, "restart")
	_, err := cmd.CombinedOutput()

	return err
}

func (cc *ConformanceContainer) Stop() error {
	cmd := exec.Command("service", cc.Name, "stop")
	_, err := cmd.CombinedOutput()

	return err
}

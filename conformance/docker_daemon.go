package docker_validation

import (
	"errors"
	"os/exec"
	"path"
)

var (
	ErrUnsupportedRuntime        = errors.New("Only 'docker' supported")
	ErrUnsupportedServiceManager = errors.New("Only 'init' or 'systemd' supported")
)

type ConformanceContainerD struct {
	Name string

	serviceType string
}

func NewConformanceContainerD(name string) (ccd ConformanceContainerD, err error) {
	if name != "docker" {
		err = ErrUnsupportedRuntime
		return
	}

	var out []byte
	cmd := exec.Command("readlink", "/proc/1/exe")
	out, err = cmd.CombinedOutput()
	if err != nil {
		return
	}

	ccd.Name = name
	ccd.serviceType = path.Base(string(out))
	if ccd.serviceType != "init" && ccd.serviceType != "systemd" {
		err = ErrUnsupportedServiceManager
	}
	return
}

func (cc *ConformanceContainerD) IsAlive() bool {
	cmd := exec.Command("docker", "ps")
	if _, err := cmd.CombinedOutput(); err != nil {
		return false
	}

	return true
}

func (cc *ConformanceContainerD) Run(sync bool, image string, args []string) error {
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

func (cc *ConformanceContainerD) Start() (err error) {
	return execService(cc.serviceType, cc.Name, "start")
}

func (cc *ConformanceContainerD) Restart() error {
	return execService(cc.serviceType, cc.Name, "restart")
}

func (cc *ConformanceContainerD) Stop() error {
	return execService(cc.serviceType, cc.Name, "stop")
}

func execService(serviceType string, serviceName string, serviceOper string) (err error) {
	switch serviceType {
	case "init":
		_, err = exec.Command("service", serviceName, serviceOper).CombinedOutput()
	case "systemd":
		_, err = exec.Command("systemctl", serviceOper, serviceName).CombinedOutput()
	default:
		err = ErrUnsupportedServiceManager
	}
	return
}

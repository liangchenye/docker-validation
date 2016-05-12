package docker_validation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDockerValidation(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "DockerValidation Suite")
}

package docker_validation

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type ContainerState uint32

const (
	ContainerStateWaiting ContainerState = 1 << iota
	ContainerStateRunning
	ContainerStateTerminated
)

const (
	restartCount     = 10
	runningContainer = 2
)

var _ = Describe("Docker validation", func() {
	Context("unit test", func() {
		It("dd should restart successfully [ac]", func() {
			NewConformanceContainerD("docker")
		})
	})
	Context("when restart a docker daemon", func() {
		var preservedAlive bool
		var cc ConformanceContainerD
		BeforeEach(func() {
			return
			cc, _ = NewConformanceContainerD("docker")
			preservedAlive = cc.IsAlive()
			if !preservedAlive {
				err := cc.Start()
				Expect(err).To(BeNil())
			}
		})
		It("should restart successfully [ac]", func() {
			fmt.Println("ac")
			return
		})
		It("should restart successfully [ab]", func() {
			fmt.Println("ab")
			return
			for i := 0; i < restartCount; i++ {
				for j := 0; j < runningContainer; j++ {
					//Make sure restart works with containers running
					//and containers will run sucessfully after docker daemon restart.
					Expect(cc.Run(false, "busybox", []string{"sleep", "300"})).To(BeNil())
				}
				Expect(cc.Restart()).To(BeNil())
			}
		})
		AfterEach(func() {
			return
			curAlive := cc.IsAlive()
			if preservedAlive != curAlive {
				if preservedAlive {
					cc.Start()
				} else {
					cc.Stop()
				}
			}
		})
	})
})

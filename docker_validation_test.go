package docker_validation

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const (
	restartCount     = 10
	runningContainer = 2
)

var _ = Describe("Docker validation", func() {

	Context("when restart a docker daemon", func() {
		var preservedAlive bool
		var cc ConformanceContainer
		BeforeEach(func() {
			cc, _ = NewConformanceContainer("docker")
			preservedAlive = cc.IsAlive()
			if !preservedAlive {
				err := cc.Start()
				Expect(err).To(BeNil())
			}
		})
		It("should restart successfully", func() {
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

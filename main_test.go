package main_test

import (
	"os"
	"os/exec"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"
)

var _ = Describe("main", func() {
	Context("validates env vars are set", func() {
		var (
			pathToTrackertix string
			err              error
		)

		BeforeEach(func() {
			os.Setenv("PIVOTAL_TRACKER_API_TOKEN", "some-token")
			os.Setenv("PROJECT_ID", "1234567")
			pathToTrackertix, err = Build("github.com/aaronshurley/trackertix")
			Expect(err).NotTo(HaveOccurred())
		})
		AfterEach(func() {
			os.Unsetenv("PIVOTAL_TRACKER_API_TOKEN")
			os.Unsetenv("PROJECT_ID")
			CleanupBuildArtifacts()
		})

		Context("when PIVOTAL_TRACKER_API_TOKEN is not set", func() {
			BeforeEach(func() {
				os.Unsetenv("PIVOTAL_TRACKER_API_TOKEN")
			})

			It("fails", func() {
				session, err := Start(&exec.Cmd{Path: pathToTrackertix}, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				Eventually(session).Should(Exit(1))
			})
		})

		Context("when PROJECT_ID is not set", func() {
			BeforeEach(func() {
				os.Unsetenv("PROJECT_ID")
			})

			It("fails", func() {
				session, err := Start(&exec.Cmd{Path: pathToTrackertix}, GinkgoWriter, GinkgoWriter)
				Expect(err).NotTo(HaveOccurred())
				Eventually(session).Should(Exit(1))
			})
		})
	})
})

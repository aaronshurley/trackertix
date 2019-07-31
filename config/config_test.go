package config_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"os"

	. "github.com/aaronshurley/trackertix/config"
)

var _ = Describe("Config", func() {
	Context("NewConfig", func() {
		Context("when required environment variables are provided", func() {
			BeforeEach(func() {
				os.Setenv("PIVOTAL_TRACKER_API_TOKEN", "some-token")
			})

			It("returns a config with provided values", func() {
				c, err := NewConfig()
				Expect(c.TrackerToken).To(Equal("some-token"))
				Expect(err).NotTo(HaveOccurred())
			})

			AfterEach(func() {
				os.Unsetenv("PIVOTAL_TRACKER_API_TOKEN")
			})
		})

		Context("when required environment variables are not provided", func() {
			It("returns an error", func() {
				_, err := NewConfig()
				Expect(err).To(HaveOccurred())
			})
		})
	})
})

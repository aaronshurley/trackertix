package tracker_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/aaronshurley/trackertix/tracker"
)

var _ = Describe("Tracker", func() {
	var (
		opts ConfigOpts
		t    Tracker
	)

	BeforeEach(func() {
		opts = ConfigOpts{
			Token:     "some-token",
			ProjectID: 123456789,
		}

		t = NewTracker(opts)
	})

	Context("SetupProject", func() {
		Context("when brand new project", func() {
			It("adds release markers", func() {
				Expect(0).To(Equal(0))
			})
		})

		Context("when project is already setup", func() {
			It("doesn't add release markers", func() {
				Expect(0).To(Equal(0))
			})
		})
	})
})

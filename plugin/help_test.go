package plugin_test

import (
	"errors"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"code.cloudfoundry.org/cflocal/mocks"
	. "code.cloudfoundry.org/cflocal/plugin"
)

var _ = Describe("Help", func() {
	var (
		mockCtrl *gomock.Controller
		mockCLI  *mocks.MockCliConnection
		mockUI   *mocks.MockUI
		help     *Help
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockCLI = mocks.NewMockCliConnection(mockCtrl)
		mockUI = mocks.NewMockUI()
		help = &Help{
			CLI: mockCLI,
			UI:  mockUI,
		}
	})

	AfterEach(func() {
		mockCtrl.Finish()
	})

	Describe("#Short", func() {
		It("should output the short usage message", func() {
			help.Short()
			Expect(string(mockUI.Out.Contents())).To(Equal("Usage:" + ShortUsage + "\n\n"))
		})
	})

	Describe("#Long", func() {
		It("should run `cf help local`", func() {
			mockCLI.EXPECT().CliCommand("help", "local")
			help.Long()
			Expect(mockUI.Err).NotTo(HaveOccurred())
		})

		Context("when `cf help local` fails", func() {
			It("should output the error", func() {
				mockCLI.EXPECT().CliCommand("help", "local").Return(nil, errors.New("some error"))
				help.Long()
				Expect(mockUI.Err).To(MatchError("some error"))
			})
		})
	})
})

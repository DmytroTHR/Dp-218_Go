package services_test

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestServicesProblem(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Services Suite")
}

type coloredText struct {
	message string
	color   string
}

func (ct coloredText) ColorableString() string {
	return fmt.Sprintf("{{%s}}%s", ct.color, ct.message)
}

func (ct coloredText) String() string {
	return fmt.Sprintf("%s", ct.message)
}

func newYellowText(message string) *coloredText {
	return &coloredText{message: message, color: "yellow"}
}

var stringProcess string

// var _ = BeforeSuite(func() {
// 	stringProcess = fmt.Sprintf("process #%d", GinkgoParallelProcess())
// 	AddReportEntry("Before suite method", ReportEntryVisibilityAlways, newYellowText("BeforeSuite -"+stringProcess))
// 	DeferCleanup(func() {
// 		AddReportEntry("Defer cleanup Suite level", ReportEntryVisibilityAlways, newYellowText("DeferCleanup -"+stringProcess))
// 	})
// })

// var _ = AfterSuite(func() {
// 	AddReportEntry("After suite method", ReportEntryVisibilityAlways, newYellowText("AfterSuite -"+stringProcess))
// })

var _ = SynchronizedBeforeSuite(
	func() []byte {
		AddReportEntry("Sync Before suite method", ReportEntryVisibilityFailureOrVerbose, newYellowText("SynchronizedBeforeSuite process1"))
		DeferCleanup(func() {
			AddReportEntry("Defer cleanup Sync process1", ReportEntryVisibilityFailureOrVerbose, newYellowText("DeferCleanup process1"))
		})
		return []byte("passed from process #1")
	},
	func(b []byte) {
		curProc := fmt.Sprintf("#%d", GinkgoParallelProcess())
		AddReportEntry("Sync Before suite method", ReportEntryVisibilityFailureOrVerbose,
			newYellowText("SynchronizedBeforeSuite allprocesses "+curProc+"\nReceived from process1: "+string(b)))
		DeferCleanup(func() {
			AddReportEntry("Defer cleanup Sync allprocesses", ReportEntryVisibilityFailureOrVerbose, newYellowText("DeferCleanup allprocesses "+curProc))
		})
	},
)

var _ = SynchronizedAfterSuite(
	func() {
		curProc := fmt.Sprintf("#%d", GinkgoParallelProcess())
		AddReportEntry("Sync After suite method", ReportEntryVisibilityFailureOrVerbose, newYellowText("SynchronizedAfterSuite allprocesses "+curProc))
	},
	func() {
		AddReportEntry("Sync After suite method", ReportEntryVisibilityFailureOrVerbose, newYellowText("SynchronizedAfterSuite process1"))
	},
)
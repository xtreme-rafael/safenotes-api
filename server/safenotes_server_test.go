package server_test

import (
	. "github.com/xtreme-rafael/safenotes-api/server"

	. "github.com/cfmobile/gospy"
	"github.com/gorilla/mux"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Server/SafeNotesServer", func() {
	var subject ServerRunner

	Context("when using constructor NewSafeNotesServer()", func() {
		BeforeEach(func() {
			subject = NewSafeNotesServer(*mux.NewRouter())
		})

		It("should create a valid ServerRunner", func() {
			Expect(subject).NotTo(BeNil())
		})

		It("should have created a valid SafeNotesServer", func() {
			Expect(subject).To(BeAssignableToTypeOf(&SafeNotesServer{}))
		})
	})

	Context("when calling RunServer()", func() {
		var spy *GoSpy

		BeforeEach(func() {
			spy = SpyAndFake(&HttpListenAndServe)
			subject = NewSafeNotesServer(*mux.NewRouter())

			subject.RunServer()
		})

		AfterEach(func() {
			spy.Restore()
		})

		It("should have called HttpListenAndServe with the default port", func() {
			Expect(spy.CallCount()).To(Equal(1))
			Expect(spy.ArgsForCall(0)[0]).To(ContainSubstring(DefaultPort))
		})
	})
})

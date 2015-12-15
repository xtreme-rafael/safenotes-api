package server_test

import (
	. "github.com/xtreme-rafael/safenotes-api/server"

	. "github.com/cfmobile/gospy/ginkgo_ext/matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
)

var _ = Describe("Server/Extern", func() {
	It("HttpListenAndServe should point to http.ListenAndServe", func() {
		Expect(HttpListenAndServe).To(BeFunction(http.ListenAndServe))
	})
})

package crypto_test

import (
	. "github.com/xtreme-rafael/safenotes-api/crypto"

	"crypto/rand"
	"encoding/base64"
	. "github.com/cfmobile/gospy/ginkgo_ext/matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Crypto Extern", func() {
	Context("when calling RandRead", func() {
		It("should be calling rand.Read", func() {
			Expect(RandRead).To(BeFunction(rand.Read))
		})
	})

	Context("when calling Base64EncodeToString", func() {
		It("should be calling base64.URLEncoding.EncodeToString", func() {
			Expect(Base64EncodeToString).To(BeFunction(base64.URLEncoding.EncodeToString))
		})
	})
})

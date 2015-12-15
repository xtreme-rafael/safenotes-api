package crypto_test

import (
	. "github.com/xtreme-rafael/safenotes-api/crypto"

	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xtreme-rafael/safenotes-api/test"
)

var _ = Describe("Crypto", func() {
	var inputNumberOfBytes int
	var resultError error

	kDefaultNumberOfBytes := 8

	kByte := byte(0x01)
	fakeRandReadSuccess := func(b []byte) (n int, err error) {
		for i, _ := range b {
			b[i] = kByte // Assigns the same byte to all positions in the slice
		}
		return len(b), nil
	}

	BeforeEach(func() { // reset input and result error variables before each test
		inputNumberOfBytes = kDefaultNumberOfBytes
		resultError = nil
	})

	Describe("GenerateRandomBytes", func() {
		var resultByteArray []byte

		BeforeEach(func() { // reset result bytearray before each test
			resultByteArray = nil
		})

		JustBeforeEach(func() {
			resultByteArray, resultError = GenerateRandomBytes(inputNumberOfBytes)
		})

		Context("when the user requests an invalid number of bytes", func() {
			BeforeEach(func() {
				inputNumberOfBytes = -1
			})

			It("should return an InvalidNumberOfBytesError", func() {
				Expect(resultError).To(Equal(InvalidNumberOfBytesError))
			})
		})

		Context("when the OS is able to generate cryptographically secure random bytes", func() {
			test.MockWithFunc(&RandRead, fakeRandReadSuccess)

			It("should return a byte array with the correct number of bytes", func() {
				Expect(resultByteArray).To(HaveLen(kDefaultNumberOfBytes))
			})

			It("should return a bytearray consistent with what rand.Read returned", func() {
				kExpectedBytes := []byte{kByte, kByte, kByte, kByte, kByte, kByte, kByte, kByte}
				Expect(resultByteArray).To(Equal(kExpectedBytes))
			})

			It("should not return an error", func() {
				Expect(resultError).To(BeNil())
			})
		})

		Context("when the random byte generation fails", func() {
			kRandReadError := errors.New("some error")
			test.MockWithReturn(&RandRead, 0, kRandReadError)

			It("should return UnableToGenerateBytesError", func() {
				Expect(resultError).To(Equal(UnableToGenerateBytesError))
			})

			It("should not return a byte array", func() {
				Expect(resultByteArray).To(BeNil())
			})
		})
	})

	Describe("GenerateRandomBase64String", func() {
		var resultBase64String string

		JustBeforeEach(func() {
			resultBase64String, resultError = GenerateRandomBase64String(inputNumberOfBytes)
		})

		Context("when generating random bytes fail", func() {
			kRandReadError := errors.New("some error")
			test.MockWithReturn(&RandRead, 0, kRandReadError)

			It("should fail with an error", func() {
				Expect(resultError).NotTo(BeNil())
			})

			It("should not return a result string", func() {
				Expect(resultBase64String).To(BeEmpty())
			})
		})

		Context("when ", func() {

		})

		Context("when generating random bytes succeed and base64 encoding works", func() {
			kBase64String := "AQEBAQEBAQE="
			test.MockWithReturn(&Base64EncodeToString, kBase64String)
			test.MockWithFunc(&RandRead, fakeRandReadSuccess)

			It("should return the base64 encoded string", func() {
				Expect(resultBase64String).To(Equal(kBase64String))
			})

			It("should not return an error", func() {
				Expect(resultError).To(BeNil())
			})
		})
	})
})

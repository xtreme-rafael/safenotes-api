package db_test

import (
	. "github.com/xtreme-rafael/safenotes-api/db"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var _ = Describe("DbexecutorImpl", func() {
	var subject DBExecutor

	var panicked bool

	panicRecover := func() {
		panicked = recover() != nil
	}

	BeforeEach(func() {
	    panicked = false
	})

	Describe("NewDBExecutor", func() {
		Context("when calling constructor with a nil DB", func() {
		    BeforeEach(func() {
				defer panicRecover()
		        subject = NewDBExecutor(nil)
		    })

			It("should have panicked", func() {
			    Expect(panicked).To(BeTrue())
			})

			It("should not have created the DBExecutor", func() {
			    Expect(subject).To(BeNil())
			})
		})

		Context("when calling the constructor with a valid DB", func() {
			BeforeEach(func() {
				defer panicRecover()
			    db, _, _ := sqlmock.New()
				subject = NewDBExecutor(db)
			})

			It("should not have panicked", func() {
			    Expect(panicked).To(BeFalse())
			})

			It("should have returned a valid DBExecutor", func() {
			    Expect(subject).NotTo(BeNil())
			})

			It("should have returned an object of type DBExecutorImpl", func() {
				_, typeOK := subject.(*DBExecutorImpl)
			    Expect(typeOK).To(BeTrue())
			})
		})
	})
})

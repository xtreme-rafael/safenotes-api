package db_test

import (
	. "github.com/xtreme-rafael/safenotes-api/db"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xtreme-rafael/go-sqlmock"
	"database/sql"
	"errors"
)

var _ = Describe("DbexecutorImpl", func() {
	var subject DBExecutor
	var panicked bool

	panicRecover := func() {
		panicked = recover() != nil
	}

	BeforeEach(func() {
		subject = nil
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

	Describe("Methods", func() {
	    var db *sql.DB
		var dbmock sqlmock.Sqlmock

		BeforeEach(func() {
			db, dbmock = nil, nil
		    db, dbmock, _ = sqlmock.New()
			subject = NewDBExecutor(db)
		})

		Describe("CloseDB()", func() {
			var errorResult error

			JustBeforeEach(func() {
			    errorResult = subject.CloseDB()
			})

			Context("when closing the DB succeeds", func() {
				BeforeEach(func() {
					dbmock.ExpectClose()
				})

				It("should have closed the DB", func() {
					err := dbmock.ExpectationsWereMet()
					Expect(err).To(BeNil())
				})

				It("should not have returned an error", func() {
				    Expect(errorResult).To(BeNil())
				})
			})

			Context("when closing the DB fails", func() {
				kError := errors.New("some error")

				BeforeEach(func() {
			        dbmock.ExpectClose().WillReturnError(kError)
			    })

				It("should have closed the DB", func() {
				    err := dbmock.ExpectationsWereMet()
					Expect(err).To(BeNil())
				})

				It("should have returned the same error as the DB.Close() call returned", func() {
				    Expect(errorResult).To(Equal(kError))
				})
			})
		})
	})
})

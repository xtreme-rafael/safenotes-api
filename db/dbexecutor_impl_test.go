package db_test

import (
	. "github.com/xtreme-rafael/safenotes-api/db"

	"database/sql"
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/xtreme-rafael/go-sqlmock"
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

		Describe("ExecWithDB()", func() {
			var errorResult error
			var worker func(*sql.DB) error

			BeforeEach(func() {
				errorResult = nil
				worker = nil
			})

			JustBeforeEach(func() {
				defer panicRecover()
				errorResult = subject.ExecWithDB(worker)
			})

			Context("when ExecWithDB() is called with a proper worker", func() {
				query := "create table test;"

				BeforeEach(func() {
					worker = func(db *sql.DB) error {
						_, err := db.Exec(query)
						return err
					}
				})

				itShouldHaveExecutedTheWorkerWithTheDB := func() {
					It("should have executed the worker with the DB in the executor", func() {
						err := dbmock.ExpectationsWereMet()
						Expect(err).To(BeNil())
					})
				}

				Context("and the worker succeeds", func() {
					BeforeEach(func() {
						dbmock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(0, 0))
					})

					itShouldHaveExecutedTheWorkerWithTheDB()
				})

				Context("and the worker fails", func() {
					kError := errors.New("some error")

					BeforeEach(func() {
						dbmock.ExpectExec(query).WillReturnError(kError)
					})

					itShouldHaveExecutedTheWorkerWithTheDB()

					It("should have returned the error from the worker", func() {
						Expect(errorResult).To(Equal(kError))
					})
				})
			})

			Context("when ExecWithDB() is called with a nil worker", func() {
				BeforeEach(func() {
					worker = nil
				})

				It("should have panicked", func() {
					Expect(panicked).To(BeTrue())
				})
			})
		})

		Describe("ExecWithTx()", func() {
			var errorResult error
			var worker func(*sql.Tx) error

			BeforeEach(func() {
				errorResult = nil
				worker = nil
			})

			JustBeforeEach(func() {
				errorResult = subject.ExecWithTx(worker)
			})

			Context("when ExecWithTx() is called with a valid worker", func() {
				query := "create table test;"

				BeforeEach(func() {
					worker = func(tx *sql.Tx) error {
						_, err := tx.Exec(query)
						return err
					}
				})

				Context("and the transaction creation succeeds", func() {
					BeforeEach(func() {
						dbmock.ExpectBegin()
					})

					itShouldHaveExecutedAllTheExpectedDBOperations := func() {
						It("should have executed all the DB commands that were expected", func() {
							err := dbmock.ExpectationsWereMet()
							Expect(err).To(BeNil())
						})
					}

					Context("when the worker succeeds", func() {
						BeforeEach(func() {
							dbmock.ExpectExec(query).WillReturnResult(sqlmock.NewResult(0, 0))
							dbmock.ExpectCommit()
						})

						itShouldHaveExecutedAllTheExpectedDBOperations()

						It("should not have returned an error", func() {
							Expect(errorResult).To(BeNil())
						})
					})

					Context("when the worker fails", func() {
						kError := errors.New("some error")

						BeforeEach(func() {
							dbmock.ExpectExec(query).WillReturnError(kError)
							dbmock.ExpectRollback()
						})

						itShouldHaveExecutedAllTheExpectedDBOperations()

						It("should have returned the error from the worker", func() {
							Expect(errorResult).To(Equal(kError))
						})
					})
				})

				Context("and the transaction creation fails", func() {
					kError := errors.New("some error")

					BeforeEach(func() {
						dbmock.ExpectBegin().WillReturnError(kError)
					})

					It("should have returned the error from the transaction creation", func() {
						Expect(errorResult).To(Equal(kError))
					})
				})
			})
		})
	})
})

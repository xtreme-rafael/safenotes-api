package test

import (
	. "github.com/cfmobile/gospy"
	. "github.com/onsi/ginkgo"
)

func Mock(target interface{}) *GoSpy {
	var spy *GoSpy

	BeforeEach(func() {
		spy = SpyAndFake(target)
	})

	AfterEach(func() {
		spy.Restore()
	})

	return spy
}

func MockWithReturn(target interface{}, fakeReturnValues ...interface{}) *GoSpy {
	var spy *GoSpy

	BeforeEach(func() {
		spy = SpyAndFakeWithReturn(target, fakeReturnValues...)
	})

	AfterEach(func() {
		spy.Restore()
	})

	return spy
}

func MockWithFunc(target interface{}, mockFunc interface{}) *GoSpy {
	var spy *GoSpy

	BeforeEach(func() {
		spy = SpyAndFakeWithFunc(target, mockFunc)
	})

	AfterEach(func() {
		spy.Restore()
	})

	return spy
}

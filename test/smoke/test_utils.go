package main

import (
	"fmt"
	"log"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/stretchr/testify/assert"
)

type TestFixture func()

type AssertErr struct {
	Message      string
	Caller       string
	EntityParent string
}

func initializeSentry() {
	err := sentry.Init(sentry.ClientOptions{
		Dsn: "https://24fa8186c38d4f8d85e79773d82b5453@o565734.ingest.sentry.io/6242521",
	})
	if err != nil {
		log.Fatal(err)
	}
}

func flushSentry() {
	sentry.Flush(time.Second)
}

func sendErrorsToSentry(errors []string) {
	for _, err := range errors {
		sentry.CaptureMessage(err)
	}
}

func mockAssertFail(t assert.TestingT, message string, msgAndArgs ...interface{}) bool {
	t.Errorf("\n%s", message)
	return false
}

func getCaller() string {
	callersArray := assert.CallerInfo()
	firstCaller := strings.Join(callersArray, " -> ")
	return firstCaller
}

func (a *AssertErr) toString() string {
	return fmt.Sprintf("Assert Error on %s Smoke Test:\n\t%s\n\tCaller: %s", a.EntityParent, a.Message, a.Caller)
}

func convertAssertErrListToStrList(errors []AssertErr) []string {
	strList := []string{}
	for _, err := range errors {
		strList = append(strList, err.toString())
	}
	return strList
}

func ExecuteSmokeTests(t *testing.T, testGroupType reflect.Type, beforeEach TestFixture, afterEach TestFixture) {
	testGroup := reflect.New(testGroupType).Elem().Interface()
	for i := 0; i < testGroupType.NumMethod(); i++ {
		m := testGroupType.Method(i)
		t.Run(m.Name, func(t *testing.T) {
			if beforeEach != nil {
				beforeEach()
			}

			in := []reflect.Value{reflect.ValueOf(testGroup), reflect.ValueOf(t)}
			m.Func.Call(in)

			if afterEach != nil {
				afterEach()
			}
		})
	}
}

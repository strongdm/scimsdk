package main

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/getsentry/sentry-go"
	"github.com/stretchr/testify/assert"
)

const userEntityName = "users"
const groupEntityName = "groups"

var assertErrs []AssertErr = []AssertErr{}

type TestFixture func()

type AssertErr struct {
	Message string
	Caller  string
}

func initializeSentry() {
	sentryDSN := os.Getenv("SENTRY_DSN")
	if sentryDSN == "" {
		log.Fatal("You must set the SENTRY_DSN Secret Key!")
	}
	err := sentry.Init(sentry.ClientOptions{
		Dsn: sentryDSN,
	})
	if err != nil {
		log.Fatal(err)
	}
}

func flushSentry() {
	sentry.Flush(time.Second)
}

func sendErrorsToSentry() {
	errs := convertAssertErrListToStrList(assertErrs)
	for _, err := range errs {
		sentry.CaptureMessage(err)
	}
	assertErrs = []AssertErr{}
}

func getCaller() string {
	callersArray := assert.CallerInfo()
	firstCaller := strings.Join(callersArray, " -> ")
	return firstCaller
}

func (a *AssertErr) toString() string {
	return fmt.Sprintf("Assert Error on Smoke Test:\n\t%s\n\tCaller: %s", a.Message, a.Caller)
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

func registerError(message string) {
	assertErrs = append(assertErrs, AssertErr{
		Message: message,
		Caller:  getCaller(),
	})
}

func assertEmpty[T struct{} | interface{}](t *testing.T, value T) {
	if ok := assert.Empty(t, value); !ok {
		registerError(fmt.Sprintf("Value is not empty, but must be empty.\n\tExpected: empty\n\tReceived: %v", value))
	}
}

func assertNotEmpty[T struct{} | interface{}](t *testing.T, value T) {
	if ok := assert.NotEmpty(t, value); !ok {
		registerError(fmt.Sprintf("Value is empty, but must be not empty.\n\tExpected: not empty\n\tReceived: %v", value))
	}
}

func assertNotNil[T struct{} | interface{}](t *testing.T, value T) {
	if ok := assert.NotNil(t, value); !ok {
		registerError(fmt.Sprintf("Value is nil, but must be not nil.\n\tExpected: not nil\n\tReceived: %v", value))
	}
}

func assertNil[T struct{} | interface{}](t *testing.T, value T) {
	if ok := assert.Nil(t, value); !ok {
		registerError(fmt.Sprintf("Value is not nil, but must be nil.\n\tExpected: nil\n\tReceived: %v", value))
	}
}

func assertGreater[T struct{} | interface{}](t *testing.T, a T, b T) {
	if ok := assert.Greater(t, a, b); !ok {
		registerError(fmt.Sprintf("The first value is not greater than the second value, but must be nil.\n\tExpected: %v > %v\n\tReceived: %v < %v", a, b, a, b))
	}
}

func assertEqual[T struct{} | interface{}](t *testing.T, a T, b T) {
	if ok := assert.Equal(t, a, b); !ok {
		registerError(fmt.Sprintf("The firstValue is not equal the secondValue, but must be equal.\n\tExpected: %v == %v\n\tReceived: %v != %v", a, b, a, b))
	}
}

func assertTrue(t *testing.T, value bool) {
	if ok := assert.True(t, value); !ok {
		registerError(fmt.Sprintf("Value is not true, but must be true.\n\tExpected: true\n\tReceived: %v", value))
	}
}

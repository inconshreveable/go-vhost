package vhost

import (
	"fmt"
	"testing"
)

// TestErrors ensures that error types for this package are implemented properly
func TestErrors(t *testing.T) {
	// test case for https://github.com/inconshreveable/go-vhost/pull/2
	// create local err vars of error interface type
	var notFoundErr error
	var badRequestErr error
	var closedErr error

	// stuff local error types in to interface values
	notFoundErr = NotFound{fmt.Errorf("test NotFound")}
	badRequestErr = BadRequest{fmt.Errorf("test BadRequest")}
	closedErr = Closed{fmt.Errorf("test Closed")}

	// assert the types
	switch errType := notFoundErr.(type) {
	case NotFound:
	default:
		t.Fatalf("expected NotFound, got: %s", errType)
	}
	switch errType := badRequestErr.(type) {
	case BadRequest:
	default:
		t.Fatalf("expected BadRequest, got: %s", errType)
	}
	switch errType := closedErr.(type) {
	case Closed:
	default:
		t.Fatalf("expected Closed, got: %s", errType)
	}
}

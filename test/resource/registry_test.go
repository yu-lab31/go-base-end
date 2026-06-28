package resource_test

import (
	"errors"
	"go-base-end/logger"
	"go-base-end/resource"
	"testing"
)

func TestGet(t *testing.T) {
	plg, err := resource.Get[*logger.Logger]()
	if err != nil {
		t.Fatalf("failed to get resource: %v", err)
	}

	if plg == nil {
		t.Fatalf("failed to get resource (nil result)")
	}

	_, err = resource.Get[logger.Logger]()
	if !errors.Is(err, resource.ResourceNotExist) {
		t.Fatalf("unexpected error: %v, expected: %v", err, resource.ResourceNotExist)
	}
}

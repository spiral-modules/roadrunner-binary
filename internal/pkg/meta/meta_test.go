package meta_test

import (
	"testing"

	"github.com/spiral/roadrunner-binary/v2/internal/pkg/meta"
)

func TestVersion(t *testing.T) {
	if value := meta.Version(); value != "0.0.0" {
		t.Errorf("Unexpected default version value: %s", value)
	}
}

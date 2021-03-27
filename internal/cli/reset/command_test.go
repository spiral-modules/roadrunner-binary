package reset_test

import (
	"testing"

	"github.com/spiral/roadrunner-binary/v2/internal/cli/reset"

	"github.com/spiral/roadrunner/v2/plugins/config"
	"github.com/stretchr/testify/assert"
)

func TestCommandProperties(t *testing.T) {
	cmd := reset.NewCommand(&config.Viper{})

	assert.Equal(t, "reset", cmd.Use)
	assert.NotNil(t, cmd.RunE)
}

func TestExecution(t *testing.T) {
	t.Skip("Command execution is not implemented yet")
}

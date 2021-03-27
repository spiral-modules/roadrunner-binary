package serve_test

import (
	"testing"

	"github.com/spiral/roadrunner-binary/v2/internal/cli/serve"

	"github.com/spiral/roadrunner/v2/plugins/config"
	"github.com/stretchr/testify/assert"
)

func TestCommandProperties(t *testing.T) {
	cmd := serve.NewCommand(&config.Viper{})

	assert.Equal(t, "serve", cmd.Use)
	assert.NotNil(t, cmd.RunE)
}

func TestExecution(t *testing.T) {
	t.Skip("Command execution is not implemented yet")
}

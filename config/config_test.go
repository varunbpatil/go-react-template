package config_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/varunbpatil/go-react-template/config"
)

func TestConfigParsing(t *testing.T) {
	t.Setenv("USER_PROPERTY1", "foo")

	config, err := config.Parse()
	require.NoError(t, err)
	require.Equal(t, "foo", config.User.UserProperty1)
}

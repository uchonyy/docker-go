package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	conf, err := LoadConfig("docker-go.yaml")
	assert.Nil(t, err)
	assert.NotNil(t, conf)
}

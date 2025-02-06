package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/unicrm/server/internal/config"
	"github.com/unicrm/server/internal/globals"
	"github.com/unicrm/server/internal/initialize"
	"github.com/unicrm/server/internal/models/system/request"
	"github.com/unicrm/server/internal/services/system"
)

func TestInitDB(t *testing.T) {

	globals.UNICRM_CONFIG = &config.UnicrmConfig{
		JWT: config.JWT{},
	}
	initialize.OtherInit()

	config := request.InitDB{}

	service := new(system.InitDBService)
	err := service.InitDB(config)
	assert.Nil(t, err)
}

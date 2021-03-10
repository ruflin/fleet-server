package fleet

import (
	"github.com/elastic/fleet-server/v7/internal/pkg/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertActionsEmpty(t *testing.T) {
	resp, token := convertActions("1234", nil)
	assert.Equal(t, resp, []ActionResp{})
	assert.Equal(t, token, "")
}

func TestConvertActions(t *testing.T) {
	actions := []model.Action{
		{
			ActionId:   "1234",
		},
	}
	resp, token := convertActions("agent-id", actions)
	assert.Equal(t, resp, []ActionResp{
		{
			AgentId: "agent-id",
			Id: "1234",
		},
	})
	assert.Equal(t, token, "")
}

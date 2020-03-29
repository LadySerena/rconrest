package players

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const playerString = "There are 3 of a max 20 players online: bar, foo, baz"

var expectedStruct = playerListResponse{PlayerList: []string{
	"bar",
	"foo",
	"baz",
}}

func TestPlayers(t *testing.T) {
	response := parseListResponse(playerString)
	assert.Equal(t, len(expectedStruct.PlayerList), len(response.PlayerList))
	assert.Equal(t, expectedStruct.PlayerList[0], response.PlayerList[0])
	assert.Equal(t, expectedStruct.PlayerList[1], response.PlayerList[1])
	assert.Equal(t, expectedStruct.PlayerList[2], response.PlayerList[2])
}

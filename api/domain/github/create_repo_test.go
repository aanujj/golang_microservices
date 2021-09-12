package github

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateRepoRequestAsJson(t *testing.T){
	request := CreateRepoRequest{
		 "macbook",
		 "Air M1",
		 "apple.com/github",
		 true,
		 false,
		 true,
		 false,
	}
	//marshal takes an input interface and attempts to create a valid json string
	bytes , error := json.Marshal(request)
	assert.Nil(t, error)
	assert.NotNil(t, bytes)
}

package gridtravelertab

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGridTraveler(t *testing.T) {
	assert := assert.New(t)

	assert.EqualValues(1, GridTraveler(1, 1))
	assert.EqualValues(2, GridTraveler(2, 2))
	assert.EqualValues(3, GridTraveler(2, 3))
	assert.EqualValues(3, GridTraveler(3, 2))
	assert.EqualValues(6, GridTraveler(3, 3))
	assert.EqualValues(2333606220, GridTraveler(18, 18))
}

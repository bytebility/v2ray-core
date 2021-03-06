package platform_test

import (
	"testing"

	. "v2ray.com/core/common/platform"
	"v2ray.com/core/testing/assert"
)

func TestNormalizeEnvName(t *testing.T) {
	assert := assert.On(t)

	cases := []struct {
		input  string
		output string
	}{
		{
			input:  "a",
			output: "A",
		},
		{
			input:  "a.a",
			output: "A_A",
		},
		{
			input:  "A.A.B",
			output: "A_A_B",
		},
	}
	for _, test := range cases {
		assert.String(NormalizeEnvName(test.input)).Equals(test.output)
	}
}

func TestEnvFlag(t *testing.T) {
	assert := assert.On(t)

	assert.Int(EnvFlag{
		Name: "xxxxx.y",
	}.GetValueAsInt(10)).Equals(10)
}

package env

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"strconv"
	"testing"
)

func TestGetEnvWithDefaultAsStringWithEnvValue(t *testing.T) {
	testCases := []struct {
		envKey     string
		envValue   string
		defaultVal string
		expected   string
	}{
		{envKey: "PORT", envValue: "8080", defaultVal: "", expected: "8080"},
		{envKey: "Env", envValue: "prod", defaultVal: "", expected: "prod"},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("'%s' should return '%s'", tC.envValue, tC.expected), func(t *testing.T) {
			err := os.Setenv(tC.envKey, tC.envValue)
			assert.Nil(t, err)
			val := GetEnvWithDefaultAsString(tC.envKey, "")
			assert.Equal(t, tC.envValue, val)
		})
	}
}

func TestGetEnvWithDefaultAsStringWithoutEnvValue(t *testing.T) {
	testCases := []struct {
		envKey     string
		defaultVal string
		expected   string
	}{
		{envKey: "HTTP_PORT", defaultVal: "8080", expected: "8080"},
		{envKey: "Environment", defaultVal: "prod", expected: "prod"},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("'%s' should return '%s'", tC.defaultVal, tC.expected), func(t *testing.T) {
			val := GetEnvWithDefaultAsString(tC.envKey, tC.defaultVal)
			assert.Equal(t, tC.defaultVal, val)
		})
	}
}

func TestGetEnvWithDefaultAsIntWithEnvValue(t *testing.T) {
	testCases := []struct {
		envKey     string
		envValue   string
		defaultVal interface{}
		expected   int
	}{
		{envKey: "PORT", envValue: "8080", defaultVal: nil, expected: 8080},
		{envKey: "Env", envValue: "1", defaultVal: 0, expected: 1},
		{envKey: "Pais", envValue: "Brasil", defaultVal: nil, expected: 0},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("'%s' should return '%d'", tC.envValue, tC.expected), func(t *testing.T) {
			err := os.Setenv(tC.envKey, tC.envValue)
			assert.Nil(t, err)
			val := GetEnvWithDefaultAsInt(tC.envKey, 0)
			value, _ := strconv.Atoi(tC.envValue)
			assert.Equal(t, value, val)
		})
	}
}

func TestGetEnvWithDefaultAsIntWithoutEnvValue(t *testing.T) {
	testCases := []struct {
		envKey     string
		defaultVal string
		expected   int
	}{
		{envKey: "PORT", defaultVal: "8080", expected: 8080},
		{envKey: "Env", defaultVal: "1", expected: 1},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprintf("'%s' should return '%d'", tC.defaultVal, tC.expected), func(t *testing.T) {
			err := os.Setenv(tC.envKey, tC.defaultVal)
			assert.Nil(t, err)
			val := GetEnvWithDefaultAsInt(tC.envKey, 0)
			value, _ := strconv.Atoi(tC.defaultVal)
			assert.Equal(t, value, val)
		})
	}
}

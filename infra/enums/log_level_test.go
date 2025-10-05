package enums

import (
	"fmt"
	"testing"

	"bou.ke/monkey"
	"github.com/stretchr/testify/assert"
)

func TestShouldGetAllLogLevels(t *testing.T) {
	// Arrange
	expected := []LogLevel{DEBUG, INFO, WARN, SUCCESS, ERROR}

	// Act
	actual := GetAllLogLevels()

	// Assert
	for i := range expected {
		assert.Equal(t, expected[i], actual[i])
	}
}

func TestShouldConvertLogLevelToStringForAllPossibleLevels(t *testing.T) {
	// Arrange
	expected := []string{"DEBUG", "INFO", "WARN", "SUCCESS", "ERROR"}

	// Act
	var actual []string
	for _, logLevel := range GetAllLogLevels() {
		logLevelStr, err := ConvertLogLevelToString(logLevel)
		assert.NoError(t, err)
		actual = append(actual, logLevelStr)
	}

	// Assert
	assert.Equal(t, expected, actual)
}

func TestShouldConvertAllPossibleLogLevelsFromStringToEnum(t *testing.T) {
	// Arrange
	expectedLogLevels := []LogLevel{DEBUG, INFO, WARN, SUCCESS, ERROR}
	logLevelStrings := []string{"DEBUG", "INFO", "WARN", "SUCCESS", "ERROR"}

	// Act & Assert
	for i, logLevelStr := range logLevelStrings {
		actualLogLevel, err := ConvertLogLevelFromStringToEnum(logLevelStr)
		assert.NoError(t, err)
		assert.Equal(t, expectedLogLevels[i], actualLogLevel)
	}
}

func TestShouldHandleErrorWhenTryingToConvertInvalidStringToLogLevel(t *testing.T) {
	// Arrange
	invalidLogLevelStr := "INVALID"

	// Act
	actualLogLevel, err := ConvertLogLevelFromStringToEnum(invalidLogLevelStr)

	// Assert
	assert.Error(t, err)
	assert.Equal(t, LogLevel(-1), actualLogLevel)
}

func TestShouldCallGetAllLogLevelsAsStringAndHandleErrorWhenGetAllLogLevelsReturnsInvalidLogLevel(t *testing.T) {
	// Arrange
	patch := monkey.Patch(GetAllLogLevels, func() []LogLevel {
		return []LogLevel{LogLevel(-1)}
	})
	defer patch.Unpatch()

	// Act
	_, err := GetAllLogLevelsAsString()

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid log level: -1")
}

func TestShouldHandlErrorWhenCallingConvertLogLevelFromStringToEnumAndGetAllLogLevelsAsStringReturnsError(t *testing.T) {
	// Arrange
	patch := monkey.Patch(GetAllLogLevelsAsString, func() ([]string, error) {
		return []string{}, fmt.Errorf("some error")
	})
	defer patch.Unpatch()

	// Act
	_, err := ConvertLogLevelFromStringToEnum("AnyInput")

	// Assert
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "some error")
}

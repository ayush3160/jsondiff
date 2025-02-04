package colorisediff

import (
    "testing"
)


// Test generated using Keploy
func TestCompareJSON_InvalidExpectedJSON(t *testing.T) {
    expectedJSON := []byte(`{"key": "value"`) // Invalid JSON
    actualJSON := []byte(`{"key": "value"}`)
    noise := map[string][]string{}
    disableColor := true

    _, err := CompareJSON(expectedJSON, actualJSON, noise, disableColor)
    if err == nil {
        t.Errorf("Expected an error for invalid expected JSON, but got nil")
    }
}

// Test generated using Keploy
func TestCompareJSON_InvalidActualJSON(t *testing.T) {
    expectedJSON := []byte(`{"key": "value"}`)
    actualJSON := []byte(`{"key": "value"`) // Invalid JSON
    noise := map[string][]string{}
    disableColor := true

    _, err := CompareJSON(expectedJSON, actualJSON, noise, disableColor)
    if err == nil {
        t.Errorf("Expected an error for invalid actual JSON, but got nil")
    }
}


// Test generated using Keploy
func TestCompareJSON_DifferentTypes(t *testing.T) {
    expectedJSON := []byte(`{"key": "value"}`)
    actualJSON := []byte(`["value1", "value2"]`) // Different type
    noise := map[string][]string{}
    disableColor := true

    diff, err := CompareJSON(expectedJSON, actualJSON, noise, disableColor)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if diff.Expected == "" || diff.Actual == "" {
        t.Errorf("Expected and Actual diffs should not be empty for different types")
    }
}


// Test generated using Keploy
func TestCompareJSON_NoDifferences(t *testing.T) {
    expectedJSON := []byte(`{"key": "value"}`)
    actualJSON := []byte(`{"key": "value"}`)
    noise := map[string][]string{}
    disableColor := true

    diff, err := CompareJSON(expectedJSON, actualJSON, noise, disableColor)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if diff.Expected != "" || diff.Actual != "" {
        t.Errorf("Expected no differences, but got diffs: %v", diff)
    }
}


// Test generated using Keploy
func TestCompareJSON_KeyValueDifferences(t *testing.T) {
    expectedJSON := []byte(`{"key1": "value1"}`)
    actualJSON := []byte(`{"key1": "value2"}`)
    noise := map[string][]string{}
    disableColor := true

    diff, err := CompareJSON(expectedJSON, actualJSON, noise, disableColor)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if diff.Expected == "" || diff.Actual == "" {
        t.Errorf("Expected and Actual diffs should not be empty for key-value differences")
    }
}


// Test generated using Keploy
func TestCompareJSON_ExpectedHasExtraKeys(t *testing.T) {
    expectedJSON := []byte(`{"key1": "value1", "key2": "value2"}`)
    actualJSON := []byte(`{"key1": "value1"}`)
    noise := map[string][]string{}
    disableColor := true

    diff, err := CompareJSON(expectedJSON, actualJSON, noise, disableColor)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if diff.Expected == "" || diff.Actual == "" {
        t.Errorf("Expected and Actual diffs should not be empty when expected JSON has extra keys")
    }
}


// Test generated using Keploy
func TestCompareJSON_ActualHasExtraKeys(t *testing.T) {
    expectedJSON := []byte(`{"key1": "value1"}`)
    actualJSON := []byte(`{"key1": "value1", "key2": "value2"}`)
    noise := map[string][]string{}
    disableColor := true

    diff, err := CompareJSON(expectedJSON, actualJSON, noise, disableColor)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if diff.Expected == "" || diff.Actual == "" {
        t.Errorf("Expected and Actual diffs should not be empty when actual JSON has extra keys")
    }
}


// Test generated using Keploy
func TestCompareJSON_NestedStructureDifferences(t *testing.T) {
    expectedJSON := []byte(`{"key1": {"nestedKey1": "value1"}}`)
    actualJSON := []byte(`{"key1": {"nestedKey1": "value2"}}`)
    noise := map[string][]string{}
    disableColor := true

    diff, err := CompareJSON(expectedJSON, actualJSON, noise, disableColor)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if diff.Expected == "" || diff.Actual == "" {
        t.Errorf("Expected and Actual diffs should not be empty for nested structure differences")
    }
}


// Test generated using Keploy
func TestCompareJSON_ArrayLengthDifferences(t *testing.T) {
    expectedJSON := []byte(`{"key1": ["value1", "value2"]}`)
    actualJSON := []byte(`{"key1": ["value1"]}`)
    noise := map[string][]string{}
    disableColor := true

    diff, err := CompareJSON(expectedJSON, actualJSON, noise, disableColor)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if diff.Expected == "" || diff.Actual == "" {
        t.Errorf("Expected and Actual diffs should not be empty for array length differences")
    }
}


// Test generated using Keploy
func TestCompareJSON_ArrayOrderDifferences(t *testing.T) {
    expectedJSON := []byte(`{"key1": ["value1", "value2"]}`)
    actualJSON := []byte(`{"key1": ["value2", "value1"]}`)
    noise := map[string][]string{}
    disableColor := true

    diff, err := CompareJSON(expectedJSON, actualJSON, noise, disableColor)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    if diff.Expected == "" || diff.Actual == "" {
        t.Errorf("Expected and Actual diffs should not be empty for array order differences")
    }
}


package colorisediff

import (
    "testing"
    "strings"
)


// Test generated using Keploy
func TestCompareJSON_InvalidExpectedJSON_Error(t *testing.T) {
    expectedJSON := []byte(`{invalid-json}`)
    actualJSON := []byte(`{"key": "value"}`)
    noise := map[string][]string{}
    disableColor := true

    _, err := CompareJSON(expectedJSON, actualJSON, noise, disableColor)
    if err == nil {
        t.Errorf("Expected an error for invalid expected JSON, but got nil")
    }
}

// Test generated using Keploy
func TestCompareJSON_InvalidActualJSON_Error(t *testing.T) {
    expectedJSON := []byte(`{"key": "value"}`)
    actualJSON := []byte(`{invalid-json}`)
    noise := map[string][]string{}
    disableColor := true

    _, err := CompareJSON(expectedJSON, actualJSON, noise, disableColor)
    if err == nil {
        t.Errorf("Expected an error for invalid actual JSON, but got nil")
    }
}


// Test generated using Keploy
func TestCompareJSON_TypeMismatch_Diff(t *testing.T) {
    expectedJSON := []byte(`{"key": "value"}`)
    actualJSON := []byte(`["value1", "value2"]`)
    noise := map[string][]string{}
    disableColor := true

    diff, err := CompareJSON(expectedJSON, actualJSON, noise, disableColor)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    if diff.Expected == "" || diff.Actual == "" {
        t.Errorf("Expected and Actual diffs should not be empty for type mismatch")
    }
}


// Test generated using Keploy
func TestCompareJSON_EmptyJSON_NoDiff(t *testing.T) {
    expectedJSON := []byte(`{}`)
    actualJSON := []byte(`{}`)
    noise := map[string][]string{}
    disableColor := true

    diff, err := CompareJSON(expectedJSON, actualJSON, noise, disableColor)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    if diff.Expected != "" || diff.Actual != "" {
        t.Errorf("Expected no differences for empty JSON inputs, but got diffs")
    }
}


// Test generated using Keploy
func TestCompareJSON_KeyValueDiff(t *testing.T) {
    expectedJSON := []byte(`{"key1": "value1", "key2": "value2"}`)
    actualJSON := []byte(`{"key1": "value1", "key2": "differentValue"}`)
    noise := map[string][]string{}
    disableColor := true

    diff, err := CompareJSON(expectedJSON, actualJSON, noise, disableColor)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    if !strings.Contains(diff.Expected, "value2") || !strings.Contains(diff.Actual, "differentValue") {
        t.Errorf("Expected and Actual diffs should highlight the key-value differences")
    }
}


// Test generated using Keploy
func TestCompareJSON_NestedStructureMismatch(t *testing.T) {
    expectedJSON := []byte(`{"key1": {"nestedKey": "nestedValue"}}`)
    actualJSON := []byte(`{"key1": {"nestedKey": "differentValue"}}`)
    noise := map[string][]string{}
    disableColor := true

    diff, err := CompareJSON(expectedJSON, actualJSON, noise, disableColor)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }

    if !strings.Contains(diff.Expected, "nestedValue") || !strings.Contains(diff.Actual, "differentValue") {
        t.Errorf("Expected and Actual diffs should highlight the nested structure differences")
    }
}


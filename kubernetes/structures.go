package kubernetes

import "strings"

// splitWithEscape splits a string by delimiter, respecting escape character.
// Example: splitWithEscape("spec.s3\.key", '.', '\\') -> ["spec", "s3.key"]
func splitWithEscape(s string, delimiter, escape rune) []string {
	placeholder := "\x00"
	escaped := string(escape) + string(delimiter)
	s = strings.ReplaceAll(s, escaped, placeholder)
	parts := strings.Split(s, string(delimiter))
	for i := range parts {
		parts[i] = strings.ReplaceAll(parts[i], placeholder, string(delimiter))
	}
	return parts
}

func expandStringSlice(s []interface{}) []string {
	result := make([]string, len(s), len(s))
	for k, v := range s {
		// Handle the Terraform parser bug which turns empty strings in lists to nil.
		if v == nil {
			result[k] = ""
		} else {
			result[k] = v.(string)
		}
	}
	return result
}

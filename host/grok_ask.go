package host

// AskFree needs no API key. Local deterministic reply code (1 = ok).
func AskFree(prompt string) int {
	if prompt == "" {
		return 0
	}
	return 1
}

// HasAPIKey reports whether a live API key is configured (host fills this).
func HasAPIKey() int {
	return 0
}

// AskAPI requires host effect ai.query / http.post (stub returns 0 offline).
func AskAPI(prompt string) int {
	return 0
}

// Ask prefers API when a key exists, otherwise AskFree.
func Ask(prompt string) int {
	if HasAPIKey() != 0 {
		return AskAPI(prompt)
	}
	return AskFree(prompt)
}

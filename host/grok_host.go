// grok_host.go — minimal host bridge for Semantic grok module
// Transpile-friendly: small surface, no heavy deps in core paths.
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const (
	defaultEndpoint = "https://api.x.ai/v1/chat/completions"
	defaultModel    = "grok-4"
)

// GrokRequest is the payload forwarded from Semantic ask/forward_to_grok.
type GrokRequest struct {
	Prompt string `json:"prompt"`
	System string `json:"system,omitempty"`
	Model  string `json:"model,omitempty"`
}

// GrokResponse is returned to the Semantic runtime.
type GrokResponse struct {
	OK     bool   `json:"ok"`
	Text   string `json:"text"`
	Source string `json:"source"` // "api" | "free_local" | "error"
	Error  string `json:"error,omitempty"`
}

type chatMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type chatRequest struct {
	Model    string        `json:"model"`
	Messages []chatMessage `json:"messages"`
}

type chatResponse struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}

// AskAPI calls xAI with XAI_API_KEY (required).
func AskAPI(prompt string) GrokResponse {
	key := strings.TrimSpace(os.Getenv("XAI_API_KEY"))
	if key == "" {
		return GrokResponse{OK: false, Source: "error", Error: "XAI_API_KEY missing"}
	}
	model := os.Getenv("GROK_MODEL")
	if model == "" {
		model = defaultModel
	}
	body, _ := json.Marshal(chatRequest{
		Model: model,
		Messages: []chatMessage{
			{Role: "user", Content: prompt},
		},
	})
	req, err := http.NewRequest(http.MethodPost, defaultEndpoint, bytes.NewReader(body))
	if err != nil {
		return GrokResponse{OK: false, Source: "error", Error: err.Error()}
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+key)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return GrokResponse{OK: false, Source: "error", Error: err.Error()}
	}
	defer res.Body.Close()
	raw, _ := io.ReadAll(res.Body)
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return GrokResponse{OK: false, Source: "error", Error: fmt.Sprintf("http %d: %s", res.StatusCode, truncate(string(raw), 200))}
	}
	var parsed chatResponse
	if err := json.Unmarshal(raw, &parsed); err != nil {
		return GrokResponse{OK: false, Source: "error", Error: err.Error()}
	}
	if len(parsed.Choices) == 0 {
		return GrokResponse{OK: false, Source: "error", Error: "empty choices"}
	}
	return GrokResponse{OK: true, Text: parsed.Choices[0].Message.Content, Source: "api"}
}

// AskFree does not require an API key.
// There is no official keyless xAI HTTP API for third-party hosts.
// This path is intentionally local/fail-closed: short deterministic help text
// so Semantic programs can call "free Grok" offline without crashing.
func AskFree(prompt string) GrokResponse {
	p := strings.TrimSpace(prompt)
	if p == "" {
		return GrokResponse{OK: true, Source: "free_local", Text: "empty prompt"}
	}
	// Lightweight local reply — not a live model, no network, no key.
	text := "free_local: no API key mode. You asked: " + truncate(p, 120) +
		". For live Grok answers set XAI_API_KEY and use AskAPI / ask()."
	return GrokResponse{OK: true, Text: text, Source: "free_local"}
}

// Ask prefers API when a key is present, otherwise AskFree.
func Ask(prompt string) GrokResponse {
	if strings.TrimSpace(os.Getenv("XAI_API_KEY")) != "" {
		return AskAPI(prompt)
	}
	return AskFree(prompt)
}

func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: grok_host <ask|ask_free|ask_api> <prompt>\n")
		os.Exit(2)
	}
	cmd := os.Args[1]
	prompt := strings.Join(os.Args[2:], " ")
	var out GrokResponse
	switch cmd {
	case "ask":
		out = Ask(prompt)
	case "ask_free":
		out = AskFree(prompt)
	case "ask_api":
		out = AskAPI(prompt)
	default:
		fmt.Fprintf(os.Stderr, "unknown command %q\n", cmd)
		os.Exit(2)
	}
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	_ = enc.Encode(out)
	if !out.OK {
		os.Exit(1)
	}
}

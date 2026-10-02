# Host API bridge for grok module

When a Semantic program calls `ask(prompt)` or `forward_to_grok(prompt)`:

1. Resolve effect `ai.query` (and optionally `http.post`).
2. Build JSON:
```json
{
  "model": "grok-4",
  "messages": [{"role": "user", "content": "<prompt>"}]
}
```
3. `POST https://api.x.ai/v1/chat/completions`
   Header: `Authorization: Bearer $XAI_API_KEY`
   Header: `Content-Type: application/json`
4. Return `choices[0].message.content` as string to the Semantic runtime.
5. If key/network missing → fail closed (capability error), do not invent answers.

Environment:
- `XAI_API_KEY` required for live forwarding

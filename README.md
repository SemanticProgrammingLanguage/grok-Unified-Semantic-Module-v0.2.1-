# grok — Unified Semantic Module (v0.2.1)

Single module for the Semantic Module Registry combining:

1. **Assist API** — ask, explain, rewrite, summarize, compare
2. **Diagnostics** — undefined/unused/shadowed, effects, types, capabilities
3. **ai.* effects** — query, complete, explain, stream, capability_check
4. **Host contracts** — request/response types + handler registration
5. **Examples** — basic and advanced usage programs

## Design principles

- Meaning first. Syntax is transport.
- Fail closed: missing `ai.*` handlers → capability error, no fake success.
- Effects are explicit on function nodes.
- Local diagnostics work without a host AI backend.

## Symbol overview

### Assist
| Symbol | Role | Effect |
|--------|------|--------|
| ask | Free-form question | ai.query |
| explain / explain_node / explain_relation | Explain program / node / edge | ai.query |
| suggest_rewrite / suggest_lowering | Rewrite or lowering hint | ai.query |
| check_meaning | Semantic health check | ai.query |
| complete | Continue from prefix | ai.complete |
| summarize_program | Short summary | ai.query |
| list_effects | Collect effects (local) | — |
| describe_binding / describe_scope | Binding/scope text | ai.query |
| compare_programs | Diff two programs by meaning | ai.query |
| translate_intent | Intent → Semantic sketch | ai.query |

### Diagnostics (mostly local)
diagnose, find_undefined, find_unused, find_shadowed, effect_report, relation_report, type_gaps, capability_gaps, validate_structure

### Effects (grok-bridge surface, same module)
ai_query, ai_complete, ai_explain, ai_stream, ai_capability_check, ai_batch_query

### Host
HostAIRequest, HostAIResponse, HostAIStreamChunk, register_ai_handler, unregister_ai_handler, list_ai_handlers

## Layout after install

```
Semantic/Modules/grok/
  grok.smod
  semantic/assist/...
  semantic/diagnostics/...
  semantic/effects/...
  semantic/host/...
  semantic/examples/...
```

## Version

0.2.1 — unified package, expanded symbol surface.

# grok capability matrix (v0.2.1)

| Capability | Required by | Host must provide | Fail-closed behavior |
|------------|-------------|-------------------|----------------------|
| ai.query | ask, explain*, check_meaning, diagnose, translate_intent, compare_programs, ai_query, ai_batch_query | Handler returning string answer | capability error |
| ai.complete | complete, ai_complete | Handler returning continuation | capability error |
| ai.explain | explain_node, explain_relation, ai_explain | Handler returning explanation | capability error |
| ai.stream | ai_stream | Streaming chunk iterator | capability error |
| ai.capability_check | ai_capability_check, example_capability_guard, register_ai_handler | Probe → bool | returns false |

## Local (no host AI)
list_effects, find_undefined, find_unused, find_shadowed, effect_report, relation_report, type_gaps, capability_gaps, validate_structure

## Host registration API
register_ai_handler(effect_name, handler) → bool  
unregister_ai_handler(effect_name) → bool  
list_ai_handlers() → [string]

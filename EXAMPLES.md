# grok module examples

## Assist
- ask("What does OperationExpr mean?")
- explain(program) / explain_node(id) / explain_relation(from, to)
- suggest_rewrite(program) / suggest_lowering(program, target)
- complete(prefix) / summarize_program(program)
- compare_programs(a, b) / translate_intent(intent)

## Diagnostics (local)
- find_undefined / find_unused / find_shadowed
- effect_report / relation_report / type_gaps / capability_gaps
- validate_structure

## Effects
- ai_query, ai_complete, ai_explain, ai_stream, ai_batch_query
- ai_capability_check("ai.query")  // fail closed if false

## Host
- register_ai_handler / unregister_ai_handler / list_ai_handlers
- HostAIRequest / HostAIResponse / HostAIStreamChunk

## Example programs in module
- example_ask, example_explain_return42, example_while_if
- example_params_call, example_diagnose_pipeline, example_capability_guard

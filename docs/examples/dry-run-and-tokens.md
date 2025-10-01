# Dry-Run, Tokens, and Truncation

Use `--dry-run` to inspect prompts, token estimates, and retrieval output without calling an AI provider.

## Token breakdown and warnings
```bash
docloom generate -p myproj --dry-run
# Output: Tokens total≈1234 (instructions≈123, docs≈1000, overhead≈111)
```

## Prompt truncation before send
```bash
docloom generate -p myproj --prompt-limit 60000 --print-prompt
```

## Budget guardrail
```bash
docloom generate -p myproj --budget-limit 0.05
# Fails if estimated max cost exceeds $0.05
```

## Retrieval preview (no embeddings sent)
```bash
# See which chunks would be injected without hitting the network
docloom generate -p myproj --retrieval --embed-model openai/text-embedding-3-small --dry-run --print-prompt

# Local embeddings version
docloom generate -p myproj --retrieval --embed-provider ollama --embed-model nomic-embed-text --dry-run
```

## Tune retries and timeouts (flags override config)
```bash
docloom --http-timeout 90 --retry-max 5 --retry-base-ms 750 --retry-max-ms 6000 \
  generate -p myproj --dry-run
```

## Machine-readable dry-run output
```bash
# Quiet log noise and emit JSON
docloom generate -p myproj --dry-run --json --quiet | jq .
```

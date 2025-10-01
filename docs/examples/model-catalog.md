# Model Catalog and Pricing

DocLoom keeps a lightweight catalog to warn about context limits and cost. You can inspect or override it at any time.

## Inspect current catalog
```bash
docloom models show | jq .
```

## Replace or merge from JSON
```bash
docloom models sync --file ./models.json          # replace
docloom models sync --file ./models.json --merge  # merge into existing entries
```

## Fetch from URL or provider preset
```bash
# Remote JSON (optional --output to persist locally)
docloom models fetch --url https://example.com/models.json --merge --output models.json

# Built-in presets (offline)
docloom models fetch --provider openrouter --merge --output models-openrouter.json
```

## Use presets during generate
```bash
# Merge the OpenRouter preset and pick a cheap model automatically
docloom generate -p myproj --model-preset cheap

# Combine provider + tier
docloom generate -p myproj --provider google --model-preset balanced

# Explicit model with preset warnings
docloom generate -p myproj --model openai/gpt-4o-mini --model-preset openrouter --max-tokens 512
```

## Budget awareness during generate
```bash
docloom generate -p myproj --model openai/gpt-4o-mini --max-tokens 512 --budget-limit 0.03
# Prints estimated max cost and fails early if the limit would be exceeded
```

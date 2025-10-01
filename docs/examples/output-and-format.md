# Output Files and Formats

DocLoom lets you control presentation, persistence, and machine-readable output.

## Save responses to files
```bash
# Markdown (default)
docloom generate -p myproj --output out.md --format markdown

# Plain text
docloom generate -p myproj --output out.txt --format text
```

## Emit JSON for automation
```bash
# Structured output to stdout
docloom generate -p myproj --json | jq .

# Write full response payload to a file
docloom generate -p myproj --json --output out.json --format json
```

## Control logging and streaming
```bash
# Quiet mode hides non-essential logs
docloom generate -p myproj --quiet --output out.md

# Real-time deltas (OpenRouter or Ollama runtimes)
docloom generate -p myproj --stream --model openai/gpt-4o-mini --max-tokens 512
```

## Inspect prompts before sending
```bash
docloom generate -p myproj --print-prompt --max-tokens 1024
```

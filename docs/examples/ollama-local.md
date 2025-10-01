# Using DocLoom with Ollama (Local Runtime)

This guide shows how to generate locally using an Ollama runtime.

## Prerequisites
- Install and run Ollama: https://ollama.com
- Ensure the model you want is available, e.g. `ollama pull llama3:latest`

## Examples

- Dry run (no API call) with local provider preset:
```
docloom generate -p myproj --provider ollama --model llama3:latest --dry-run --print-prompt
```

- Streaming output from a local model:
```
docloom generate -p myproj --provider ollama --model llama3:latest --stream
```

- Choose by tier with a preset (if you didn’t set `--model` explicitly):
```
docloom generate -p myproj --provider ollama --model-preset balanced
```

- Retrieval using local embeddings (builds/refreshes `index.json`):
```
docloom generate -p myproj --provider ollama \
  --retrieval --embed-provider ollama --embed-model nomic-embed-text \
  --model llama3:latest --top-k 6 --min-score 0.2
```

- Configure host/timeout via env or config:
```
export DOCLOOM_OLLAMA_HOST="http://127.0.0.1:11434"
export DOCLOOM_OLLAMA_TIMEOUT_SEC=120
```
Or in `~/.docloom-cli/config.yaml`:
```
ollama_host: "http://127.0.0.1:11434"
ollama_timeout_sec: 120
```

If a model is missing, DocLoom will suggest pulling it, e.g.: `ollama pull llama3:latest`.

Tip: To default to local runs, set in `~/.docloom-cli/config.yaml`:
```
default_provider: ollama
```

For retrieval defaults you can also set:
```
embedding_provider: ollama
embedding_model: nomic-embed-text
retrieval_top_k: 6
retrieval_min_score: 0.2
```

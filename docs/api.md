# OpenRouter Integration (DocLoom)

This document describes how DocLoom interacts with the OpenRouter API for chat completions.

## Endpoint (OpenRouter)

- URL: `POST https://openrouter.ai/api/v1/chat/completions`
- Headers:
  - `Authorization: Bearer <OPENROUTER_API_KEY>`
  - `Content-Type: application/json`
  - `HTTP-Referer: https://github.com/KaramelBytes/docloom-cli`
  - `X-Title: DocLoom CLI`
  - On responses, DocLoom captures common request ID headers (e.g., `X-Request-Id`, `OpenAI-Request-ID`) and prints them on success for observability.

## Request Body (MVP)

```json
{
  "model": "openai/gpt-4o-mini",
  "messages": [
    { "role": "user", "content": "[INSTRUCTIONS]... [REFERENCE DOCUMENTS]... [TASK]..." }
  ],
  "max_tokens": 1024,
  "temperature": 0.7
}
```

- `model`: Provider/model identifier as listed by OpenRouter (e.g., `openai/gpt-4o-mini`).
- `messages`: Single-message user prompt containing instructions, documents, and task.
- `max_tokens`: Target maximum response tokens.
- `temperature`: Sampling temperature.

## Response Body (Simplified)

```json
{
  "id": "resp_...",
  "choices": [
    {
      "message": { "role": "assistant", "content": "..." }
    }
  ],
  "usage": {
    "prompt_tokens": 1234,
    "completion_tokens": 456,
    "total_tokens": 1690
  }
}
```

DocLoom prints the first choice's message content.
If a request ID is present in response headers, DocLoom prints it for traceability.

## Errors

- Non-2xx status codes are returned with a short message containing `status` and parsed body (if possible).
- Common issues:
  - Missing or invalid API key
  - Unsupported model name
  - Token/context limit exceeded

## Token and Context Notes

DocLoom estimates prompt tokens locally and prints warnings. Context limits vary by model; verify limits on OpenRouter model docs. Use `--dry-run` to preview prompt and token breakdown without performing an API call.

## Security

- API key is never logged.
- Read from environment (`OPENROUTER_API_KEY`) or config (`~/.docloom-cli/config.yaml`).

## Local Runtime (Ollama)

- Select with `--provider ollama` (alias: `local`) or set `default_provider: ollama` in config.
- Host defaults to `http://127.0.0.1:11434`; configure with `ollama_host` or `DOCLOOM_OLLAMA_HOST`.
- API: `POST /api/chat` with `stream=false`.
- On success, DocLoom prints a local correlation token (simulated Request ID).
 - Streaming: use `--stream` to enable incremental output. Supported for Ollama and OpenRouter.

## Runtime Abstraction

DocLoom defines a small runtime interface so multiple backends can implement the same `Generate` surface. The CLI selects the runtime via `--provider` and applies model presets independently of the runtime choice.

## Model Catalog and Presets

- Inspect current catalog: `docloom models show | jq .`
- Apply a built-in provider preset without network: `docloom models fetch --provider openrouter --merge`
- During generation you can apply presets inline:
  - Provider catalog: `docloom generate -p myproj --model-preset openrouter`
  - Tiered selection (chooses a model if `--model` not set): `docloom generate -p myproj --model-preset cheap|balanced|high-context`
  - Explicit provider guidance: `docloom generate -p myproj --provider google --model-preset balanced`
  - Combined: `docloom generate -p myproj --model-preset openrouter:cheap`
  - Presets merge a curated catalog before generation so pricing/context warnings reflect it.

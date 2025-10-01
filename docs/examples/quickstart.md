# Quickstart: Create, Add, Instruct, Generate

This walkthrough shows a minimal end‑to‑end flow using DocLoom.

1) Initialize a project

```bash
docloom init myproj -d "Docs to merge"
```

2) Add documents

```bash
docloom add -p myproj ./README.md --desc "Main readme"
# add more docs as needed; supports .txt, .md, .docx, .csv, .tsv, .xlsx

# Tip: CSV/TSV/XLSX are summarized instead of printed raw.
# You can pre-check or export a summary with:
#   docloom analyze ./data/metrics.csv --output metrics_summary.md

## Optional: Use analysis instructions

# Include an instruction template to guide the model’s interpretation of the dataset summary:
docloom add -p myproj docs/templates/dataset-analysis.md --desc "Analysis Instructions"

# Or set the project instructions from the file content:
# docloom instruct -p myproj "$(cat docs/templates/dataset-analysis.md)"
```

3) Set instructions

```bash
docloom instruct -p myproj "Summarize key features and provide a short overview."
```

4) (Optional) Enable retrieval for context-aware prompts

```bash
# OpenRouter embeddings (defaults to instructions + top 6 chunks)
docloom generate -p myproj --retrieval --embed-model openai/text-embedding-3-small --dry-run

# Local embeddings via Ollama (requires an embedding model such as nomic-embed-text)
docloom generate -p myproj --retrieval --embed-provider ollama --embed-model nomic-embed-text --dry-run
```

5) Dry run to inspect prompt and token breakdown (no API call)

```bash
docloom generate -p myproj --dry-run --print-prompt
```

6) Real run (requires API key)

```bash
export OPENROUTER_API_KEY=your_key
docloom generate -p myproj --model openai/gpt-4o-mini --max-tokens 512
```

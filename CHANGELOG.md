# Changelog

All notable changes to DocLoom CLI will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.2.0] - 2025-10-15

### 🎉 Added
- **Batch Analysis**: New `analyze-batch` command processes multiple files with `[N/Total]` progress
- **Mixed-Input Batch**: Supports `.csv`, `.tsv`, `.xlsx` (analyzed) + `.yaml`, `.md`, `.txt`, `.docx` (added as docs)
- **Project-Level Sample Control**: `--sample-rows-project` flag to override samples in all summaries (set `0` to disable)
- **Memory Safety**: Hard limits prevent OOM (200k tokens, 20 summaries per project)
- **Context Validation**: Blocks oversized prompts for local LLMs with actionable error messages
- **Timeout Configuration**: `--timeout-sec` flag for generation requests (default 180s)
- **TSV Auto-Delimiter**: Automatically sets tab delimiter for `.tsv` files

### 🐛 Fixed
- **CRITICAL**: XLSX parser returning 0 columns due to absolute relationship paths in ZIP archives
- Unbounded memory accumulation with multiple large files (9.3GB → <2GB peak)
- Duplicate document detection (no more silent overwrites)
- Memory leaks in outlier computation
- Context window overflow causing silent truncation in Ollama
- RAG chunker producing oversized chunks exceeding token limits
- Prompt instruction duplication (40% token reduction)
- Dataset summary basename collisions with disambiguation logic
- Invalid `--sheet-name` silently falling back to first sheet

### ⚡ Performance
- Reduced memory usage by 78% for multi-file projects
- Batched embedding prevents API timeout failures (100 chunks/batch)
- 40% reduction in prompt tokens via deduplication
- Immediate memory release after outlier computation

### 💥 Breaking Changes
- Context overflow now **blocks** execution for Ollama (was warning-only)
- Duplicate files now **error** instead of silently overwriting
- Invalid `--sheet-name` now errors with available sheet list
- Projects enforce maximum 200k token limit (hard cap at 200k)
- Maximum 20 dataset summaries per project (prevents context bloat)

### 📚 Documentation
- Added [docs/examples/analyze-batch.md](cci:7://file:///home/jeremiah/Projects/docloom-cli/docs/examples/analyze-batch.md:0:0-0:0) with batch processing examples
- Updated README with mixed-input batch behavior
- Added XLSX parser fix details and regression test
- Updated quickstart with batch analysis tips

### 🧪 Testing
- Added regression test for XLSX relationship path normalization
- Added integration test for batch analysis with sample suppression
- Memory profiling tests ensure <2GB peak for 10x100k-row files
- Race detector clean across all packages

## [0.1.0] - 2025-10-01

### Added
- Initial release
- Basic project management (`init`, `add`, `list`)
- CSV/TSV/XLSX analysis with schema inference
- OpenRouter, Ollama, and major provider support
- RAG with embedding indexes
- Model catalog management

[0.2.0]: https://github.com/KaramelBytes/docloom-cli/releases/tag/v0.2.0
[0.1.0]: https://github.com/KaramelBytes/docloom-cli/releases/tag/v0.1.0
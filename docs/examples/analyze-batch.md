# Analyze Multiple Datasets with Progress

Use `analyze-batch` to summarize many CSV/TSV/XLSX files with a single command. This prints progress as each file is processed and can attach summaries to a project.

## Examples

- Process a folder of datasets with progress

```bash
docloom analyze-batch "data/*.csv"
```

- Attach all summaries to a project (and suppress sample tables)

```bash
docloom analyze-batch "data/*.xlsx" \
  -p brewlab --desc "Batch dataset summaries" \
  --sample-rows-project 0
```

- Select XLSX sheet and set CSV/locale options

```bash
docloom analyze-batch data/*.xlsx \
  --sheet-name "Aug 2024" \
  --delimiter ',' --decimal dot --thousands ,
```

## Behavior

- Shows progress: `[N/Total] Processing <file>...` (use `--quiet` to suppress)
- Mirrors `analyze` flags (grouping, correlations, outliers, locale)
- When attaching summaries to a project (`-p`):
  - `dataset_summaries/` is created under the project directory
  - Filenames are disambiguated:
    - With `--sheet-name`, sheet slug is added: `name__sheet-sales.summary.md`
    - On collision, an increment is appended: `name__2.summary.md`
  - Use `--sample-rows-project` to override sample rows for all outputs (set `0` to disable sample tables)

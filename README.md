# aoc2024

## Linking Rust subdirs in rust-analyzer via vscode

Gotta do settings in JSON like so (relative path is fine with workspace-specific settings):
```json
{
    "rust-analyzer.linkedProjects": [
        "night_01/rust/Cargo.toml"
    ]
}
```

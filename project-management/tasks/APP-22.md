## Name:
Create Config dependency

---
## Description:
Clean up loose flags and parsing into a single Config struct that can be passed to different places that need it (dependency injection)

---
#### Subtasks:
- [x] Rename `flags` to `config`
- [x] make Configuration struct with all config fields
- [x] do all flag parsing in `config.Parse()` and return Configuration struct

---
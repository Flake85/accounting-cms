## Name:
Create Config dependency

---
## Description:
Clean up loose flags and parsing into a single Config struct that can be passed to different places that need it (dependency injection)

---
#### Subtasks:
- [ ] Rename `flags` to `config`
- [ ] make Configuration struct with all config fields
- [ ] do all flag parsing in `config.Parse()` and return Configuration struct

---
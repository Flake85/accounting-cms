## Name:
Repository DB cleanup

---
## Description:
Use dependency injection to config DB for repositories instead of global DB pollution

---
#### Subtasks:
- [x] Create Repository structs that have a DB field
- [x] Create each repo struct and pass DB as a dependency in constructor func
- [x] Can optionally make only a single Repo struct instead of one for each entity

---
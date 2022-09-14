## Name:
Cleanup DB construction

---
## Description:
Move the DB creation to `db` package, make `NewDB` method that accepts new Config as an argument and then return the `gorm.DB`

---

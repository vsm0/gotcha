**Basic Gacha System**
==========================

A Go + SQLite base for a gacha centered system, providing fundamental models (tables) for a functioning gacha system. This project is based on a previously designed on-paper ERD, adapted to Go.

`model/` contains the model files (tables) for the gacha system:

* `account.go`: Represents user accounts
* `item.go`: Represents global items
* `inventoryitem.go`: Represents user inventory items
* `event.go`: Represents events
* `reward.go`: Represents event rewards

**Tests**
--------

`test/` contains the test files for the gacha system. These tests provide a reference for implementing the functionalities in a full app.

### Testing Requirements

* Go 1.22+ is required to run the tests.

### Running Tests

* Run all tests:
```bash
go test -v ./test
```

### Features in Tests

The following features are implemented or planned:

- [x] Create schema (e.g. on first run)
- [x] Delete schema

- [x] Create user accounts
- [x] Delete user accounts
- [x] Update user accounts (e.g. changing password or other details)
- [x] Query user accounts

- [x] Create global items
- [x] Delete global items (e.g. on system update)
- [x] Query global items

- [x] Create user inventory items
- [x] Delete user inventory items
- [x] Update user inventory items (e.g. when item count changes)
- [x] Query user inventory items

- [x] Create events
- [x] Delete events
- [x] Query events

- [x] Create event rewards
- [x] Delete event rewards
- [x] Query event rewards

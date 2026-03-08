---
stepsCompleted: ['step-01-detect-mode', 'step-02-load-context', 'step-03-risk-and-testability', 'step-04-coverage-plan', 'step-05-generate-output']
lastStep: 'step-05-generate-output'
lastSaved: '2026-03-08T10:22:00Z'
workflowType: 'testarch-test-design'
inputDocuments: ['main.go', 'models.go', 'storage.go', 'commands.go', 'README.md', 'Makefile', 'go.mod']
---

# Test Design for QA: Task Tracker CLI Application

**Purpose:** Test execution recipe for QA team. Defines what to test, how to test it, and what QA needs from other teams.

**Date:** 2026-03-08
**Author:** Test Architect (TEA)
**Status:** Draft
**Project:** Task Tracker

**Related:** See Architecture doc (test-design-architecture.md) for testability concerns and architectural blockers.

---

## Executive Summary

**Scope:** Comprehensive testing of CLI task management application with JSON file storage

**Risk Summary:**

- Total Risks: 6 (2 high-priority score ≥6, 1 medium, 3 low)
- Critical Categories: DATA (data integrity), BUS (business logic)

**Coverage Summary:**

- P0 tests: ~8 (critical paths, core functionality)
- P1 tests: ~7 (important features, integration)
- P2 tests: ~8 (edge cases, performance)
- P3 tests: ~3 (exploratory, benchmarks)
- **Total**: ~26 tests (~3-4 weeks with 1 developer)

---

## Not in Scope

**Components or systems explicitly excluded from this test plan:**

| Item       | Reasoning                   | Mitigation                                                                      |
| ---------- | --------------------------- | ------------------------------------------------------------------------------- |
| **External services** | No external dependencies in application | Not applicable for local-only CLI tool |
| **Multi-user scenarios** | Single-user application by design | Future enhancement consideration |
| **Network operations** | Local filesystem only | Not applicable |

**Note:** Items listed here have been reviewed and accepted as out-of-scope by QA, Dev, and PM.

---

## Dependencies & Test Blockers

**CRITICAL:** QA cannot proceed without these items from other teams.

### Backend/Architecture Dependencies (Pre-Implementation)

**Source:** See Architecture doc "Quick Guide" for detailed mitigation plans

1. **ASR-001: Interactive input interface** - Dev Team - Pre-implementation
   - Extract `bufio.Reader` into interface for mocking
   - Enables automated testing of `addTask` command

2. **ASR-002: Configurable storage path** - Dev Team - Pre-implementation
   - Allow override of `~/.task-tracker` path
   - Enables isolated test environments

### QA Infrastructure Setup (Pre-Implementation)

1. **Test Data Factories** - QA
   - Task factory with faker-based randomization
   - Auto-cleanup fixtures for parallel safety

2. **Test Environments** - QA
   - Local: Temp directory setup for each test
   - CI/CD: Isolated filesystem environments
   - Staging: Not applicable (local-only tool)

**Example factory pattern:**

```go
package testutils

import (
    "testing"
    "time"
    "github.com/google/uuid"
)

func CreateTestTask(t *testing.T, overrides ...func(*Task)) Task {
    task := Task{
        ID:          uuid.New().String(),
        Title:       "Test Task " + uuid.New().String()[:8],
        Description: "Test description",
        Status:      StatusTodo,
        CreatedAt:   time.Now(),
        UpdatedAt:   time.Now(),
    }
    
    for _, override := range overrides {
        override(&task)
    }
    
    return task
}
```

---

## Risk Assessment

**Note:** Full risk details in Architecture doc. This section summarizes risks relevant to QA test planning.

### High-Priority Risks (Score ≥6)

| Risk ID    | Category | Description         | Score       | QA Test Coverage             |
| ---------- | -------- | ------------------- | ----------- | ---------------------------- |
| **DATA-001** | DATA    | JSON file corruption | **6** | JSON validation tests, backup/restore scenarios |
| **BUS-001** | BUS     | Data loss during operations | **6** | Atomic operation tests, crash recovery scenarios |

### Medium/Low-Priority Risks

| Risk ID | Category | Description         | Score   | QA Test Coverage             |
| ------- | -------- | ------------------- | ------- | ---------------------------- |
| PERF-001  | PERF    | Performance with large task lists | 4 | Performance benchmarks with 1000+ tasks |
| SEC-001  | SEC     | File permission issues | 2 | Permission error scenarios |
| OPS-001  | OPS     | Cross-platform deployment | 1 | OS-specific path handling tests |
| TECH-001  | TECH    | Dependency breaking changes | 2 | Dependency version validation |

---

## Entry Criteria

**QA testing cannot begin until ALL of the following are met:**

- [ ] All requirements and assumptions agreed upon by QA, Dev, PM
- [ ] Test environments provisioned and accessible
- [ ] Test data factories ready or seed data available
- [ ] Pre-implementation blockers resolved (see Dependencies section)
- [ ] Feature buildable via `make build`
- [ ] Interactive input interface extracted

## Exit Criteria

**Testing phase is complete when ALL of the following are met:**

- [ ] All P0 tests passing
- [ ] All P1 tests passing (or failures triaged and accepted)
- [ ] No open high-priority / high-severity bugs
- [ ] Test coverage ≥80% agreed as sufficient by QA Lead and Dev Lead
- [ ] Performance baselines met (<2s for 1000 tasks)
- [ ] Atomic file operations validated

---

## Test Coverage Plan

**IMPORTANT:** P0/P1/P2/P3 = **priority and risk level** (what to focus on if time-constrained), NOT execution timing. See "Execution Strategy" for when tests run.

### P0 (Critical)

**Criteria:** Blocks core functionality + High risk (≥6) + No workaround + Affects majority of users

| Test ID    | Requirement   | Test Level | Risk Link | Notes   |
| ---------- | ------------- | ---------- | --------- | ------- |
| **P0-001** | Create task with valid data | Unit | DATA-001 | Core business logic |
| **P0-002** | Create task with empty title (error) | Unit | DATA-001 | Input validation |
| **P0-003** | Load tasks from non-existent file | Unit | BUS-001 | Edge case handling |
| **P0-004** | Atomic file write operations | Integration | BUS-001 | Data loss mitigation |
| **P0-005** | `init` command creates directory and file | Integration | - | Setup prerequisite |
| **P0-006** | `add` command with interactive input | Integration | - | Core user journey |
| **P0-007** | `list` command displays tasks correctly | Integration | - | Primary read operation |
| **P0-008** | `complete` command updates task status | Integration | - | State transition |

**Total P0:** ~8 tests

---

### P1 (High)

**Criteria:** Important features + Medium risk (3-4) + Common workflows + Workaround exists but difficult

| Test ID    | Requirement   | Test Level | Risk Link | Notes   |
| ---------- | ------------- | ---------- | --------- | ------- |
| **P1-001** | Save tasks with invalid JSON structure | Unit | DATA-001 | Data corruption risk |
| **P1-002** | `delete` command removes task | Integration | - | Data removal |
| **P1-003** | Command with invalid task ID | Integration | - | Error handling |
| **P1-004** | Task struct validation | Unit | - | Data integrity |
| **P1-005** | JSON serialization/deserialization | Unit | DATA-001 | Data persistence risk |
| **P1-006** | Complete task lifecycle (add→complete→delete) | E2E | - | Critical user journey |
| **P1-007** | Concurrent file access handling | Integration | BUS-001 | Race condition prevention |

**Total P1:** ~7 tests

---

### P2 (Medium)

**Criteria:** Secondary features + Low risk (1-2) + Edge cases + Regression prevention

| Test ID    | Requirement   | Test Level | Risk Link | Notes   |
| ---------- | ------------- | ---------- | --------- | ------- |
| **P2-001** | Status constants validation | Unit | - | Enum safety |
| **P2-002** | Timestamp handling (created/updated) | Unit | - | Temporal logic |
| **P2-003** | File permission error handling | Integration | SEC-001 | Security boundary |
| **P2-004** | Command aliases work (a, l, c, d) | Integration | - | UX convenience |
| **P2-005** | Load 1000+ tasks performance | Integration | PERF-001 | Scale requirement |
| **P2-006** | Large task list display performance | Integration | PERF-001 | UX responsiveness |
| **P2-007** | Different OS path handling | Integration | OPS-001 | Portability |
| **P2-008** | Multiple tasks management workflow | E2E | - | Realistic usage |

**Total P2:** ~8 tests

---

### P3 (Low)

**Criteria:** Nice-to-have + Exploratory + Performance benchmarks + Documentation validation

| Test ID    | Requirement   | Test Level | Notes   |
| ---------- | ------------- | ---------- | ------- |
| **P3-001** | File size impact on operations | Integration | Resource limits |
| **P3-002** | Error recovery workflow | E2E | Resilience testing |
| **P3-003** | Memory usage with 10k tasks | Integration | Resource monitoring |

**Total P3:** ~3 tests

---

## Execution Strategy

**Philosophy:** Run everything in PRs unless there's significant infrastructure overhead. Go tests are extremely fast (100s of tests in ~5-10 min).

**Organized by TOOL TYPE:**

### Every PR: Go Tests (~8-12 minutes)

**All functional tests** (from any priority level):

- All unit, integration, and E2E tests using Go testing package
- Parallelized across available CPU cores
- Total: ~26 Go tests (includes P0, P1, P2, P3)

**Why run in PRs:** Fast feedback, no expensive infrastructure

### Nightly: Performance Tests (~20-30 minutes)

**All performance tests** (from any priority level):

- Large dataset benchmarks (1000+, 10000+ tasks)
- Memory usage profiling
- File I/O performance analysis
- Total: ~3 performance tests (P2-P3)

**Why defer to nightly:** Longer execution time, resource-intensive

### Weekly: Chaos Engineering (~45-60 minutes)

**Special infrastructure tests** (from any priority level):

- File corruption simulation
- Crash during write operations
- Backup/restore validation
- Total: ~2 chaos tests (validating high-risk scenarios)

**Why defer to weekly:** Destructive testing, infrequent validation sufficient

**Manual tests** (excluded from automation):

- Usability testing with real users
- Documentation validation
- Cross-platform manual verification

---

## QA Effort Estimate

**QA test development effort only** (excludes DevOps, Backend, Data Eng, Finance work):

| Priority  | Count | Effort Range       | Notes                                             |
| --------- | ----- | ------------------ | ------------------------------------------------- |
| P0        | ~8    | ~25-40 hours      | Complex setup (atomic operations, file handling)   |
| P1        | ~7    | ~20-35 hours      | Standard coverage (integration, API tests)        |
| P2        | ~8    | ~15-25 hours      | Edge cases, performance validation                 |
| P3        | ~3    | ~5-10 hours       | Exploratory, benchmarks                           |
| **Total** | ~26   | **~65-110 hours** | **1 developer, 3-4 weeks full-time**              |

**Assumptions:**

- Includes test design, implementation, debugging, CI integration
- Excludes ongoing maintenance (~2-4 hours per sprint)
- Assumes test infrastructure (factories, fixtures) ready

**Dependencies from other teams:**

- See "Dependencies & Test Blockers" section for what QA needs from Dev Team

---

## Implementation Planning Handoff

**Use this to inform implementation planning; assign to Dev owners.**

| Work Item   | Owner        | Target Milestone | Dependencies/Notes |
| ----------- | ------------ | ----------------- | ------------------ |
| Extract interactive input interface | Dev Team | Pre-implementation | Required for P0-006 test automation |
| Implement configurable storage path | Dev Team | Pre-implementation | Required for isolated test environments |
| Add atomic file write operations | Dev Team | Implementation | Mitigates BUS-001 risk |
| Implement JSON validation | Dev Team | Implementation | Mitigates DATA-001 risk |

---

## Tooling & Access

**Standard Go tooling - no special access required:**

| Tool or Service   | Purpose   | Access Required | Status            |
| ----------------- | --------- | --------------- | ----------------- |
| Go testing package | Unit/Integration tests | Standard installation | Ready |
| Testify library | Enhanced assertions | Go module dependency | Ready |
| Temp directory | Test isolation | OS filesystem access | Ready |

**Access requests needed (if any):**

- [ ] None required (standard development environment sufficient)

---

## Interworking & Regression

**Services and components impacted by this feature:**

| Service/Component | Impact              | Regression Scope                | Validation Steps              |
| ----------------- | ------------------- | ------------------------------- | ----------------------------- |
| **File System**     | Direct interaction | All file I/O operations must work | Verify file permissions, atomic writes |
| **JSON Parser**      | Data serialization | JSON parsing must handle edge cases | Test malformed JSON, empty files |
| **CLI Framework**    | Command processing | All CLI commands must function | Test all commands and aliases |

**Regression test strategy:**

- All existing CLI commands must continue to work
- JSON file format must remain backward compatible
- Performance must not degrade for typical usage (<100 tasks)
- No cross-team coordination needed (single-component application)

---

## Appendix A: Code Examples & Tagging

**Go Test Tags for Selective Execution:**

```go
// P0 critical test
func TestCreateTaskWithValidData(t *testing.T) {
    // Setup
    storage := &TaskStorage{}
    
    // Execute
    task := createTask("Test Task", "Test Description")
    
    // Validate
    if task.Title != "Test Task" {
        t.Errorf("Expected title 'Test Task', got '%s'", task.Title)
    }
    if task.Status != StatusTodo {
        t.Errorf("Expected status '%s', got '%s'", StatusTodo, task.Status)
    }
}

// P1 integration test
func TestAddCommandIntegration(t *testing.T) {
    // Setup temp directory
    tempDir := t.TempDir()
    dataFile := filepath.Join(tempDir, "tasks.json")
    
    // Execute
    err := initStorage(tempDir, dataFile)
    if err != nil {
        t.Fatalf("Failed to initialize storage: %v", err)
    }
    
    // Validate file exists and is valid JSON
    data, err := os.ReadFile(dataFile)
    if err != nil {
        t.Fatalf("Failed to read data file: %v", err)
    }
    
    var storage TaskStorage
    err = json.Unmarshal(data, &storage)
    if err != nil {
        t.Fatalf("Failed to unmarshal JSON: %v", err)
    }
    
    if len(storage.Tasks) != 0 {
        t.Errorf("Expected 0 tasks, got %d", len(storage.Tasks))
    }
}
```

**Run specific tags:**

```bash
# Run only P0 tests
go test -run P0 ./...

# Run P0 + P1 tests
go test -run "(P0|P1)" ./...

# Run performance tests
go test -run Perf ./...

# Run all tests in PR (default)
go test ./...
```

---

## Appendix B: Knowledge Base References

- **Risk Governance**: `risk-governance.md` - Risk scoring methodology
- **Test Priorities Matrix**: `test-priorities-matrix.md` - P0-P3 criteria
- **Test Levels Framework**: `test-levels-framework.md` - E2E vs API vs Unit selection
- **Test Quality**: `test-quality.md` - Definition of Done (no hard waits, <300 lines, <1.5 min)

---

**Generated by:** BMad TEA Agent
**Workflow:** `_bmad/tea/testarch/test-design`
**Version:** 4.0 (BMad v6)

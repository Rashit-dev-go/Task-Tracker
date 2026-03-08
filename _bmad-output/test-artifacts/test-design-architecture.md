---
stepsCompleted: ['step-01-detect-mode', 'step-02-load-context', 'step-03-risk-and-testability', 'step-04-coverage-plan', 'step-05-generate-output']
lastStep: 'step-05-generate-output'
lastSaved: '2026-03-08T10:22:00Z'
workflowType: 'testarch-test-design'
inputDocuments: ['main.go', 'models.go', 'storage.go', 'commands.go', 'README.md', 'Makefile', 'go.mod']
---

# Test Design for Architecture: Task Tracker CLI Application

**Purpose:** Architectural concerns, testability gaps, and NFR requirements for review by Architecture/Dev teams. Serves as a contract between QA and Engineering on what must be addressed before test development begins.

**Date:** 2026-03-08
**Author:** Test Architect (TEA)
**Status:** Architecture Review Pending
**Project:** Task Tracker
**PRD Reference:** README.md (functional requirements)
**ADR Reference:** Source code analysis (implicit architecture)

---

## Executive Summary

**Scope:** Simple CLI task management application with JSON file-based storage

**Business Context** (from README.md):

- **Revenue/Impact:** Personal productivity tool (no direct revenue)
- **Problem:** Personal task tracking without external dependencies
- **GA Launch:** Ready for testing (implementation complete)

**Architecture** (from source code analysis):

- **Key Decision 1:** JSON file-based persistence in user home directory
- **Key Decision 2:** CLI framework using urfave/cli/v2
- **Key Decision 3:** Pure Go implementation with no external services

**Expected Scale** (from requirements):

- Single user, local filesystem storage
- Task count: <10,000 tasks per user
- CLI response time: <500ms

**Risk Summary:**

- **Total risks**: 6
- **High-priority (≥6)**: 2 risks requiring immediate mitigation
- **Test effort**: ~26 tests (~3-4 weeks for 1 developer)

---

## Quick Guide

### 🚨 BLOCKERS - Team Must Decide (Can't Proceed Without)

**Pre-Implementation Critical Path** - These MUST be completed before QA can write integration tests:

1. **ASR-001: Extract interactive input interface** - Make `addTask` testable by extracting `bufio.Reader` into interface (recommended owner: Development Team)
2. **ASR-002: Configurable storage path** - Allow override of `~/.task-tracker` for test environments (recommended owner: Development Team)

**What we need from team:** Complete these 2 items pre-implementation or test development is blocked.

---

### ⚠️ HIGH PRIORITY - Team Should Validate (We Provide Recommendation, You Approve)

1. **DATA-001: JSON file corruption** - Implement atomic file writes and JSON validation (implementation phase)
2. **BUS-001: Data loss during operations** - Add backup/restore mechanisms (implementation phase)

**What we need from team:** Review recommendations and approve (or suggest changes).

---

### 📋 INFO ONLY - Solutions Provided (Review, No Decisions Needed)

1. **Test strategy**: Unit-heavy approach with integration for file operations (file system dependencies require integration testing)
2. **Tooling**: Go standard testing package + testify for assertions
3. **Tiered CI/CD**: PR (<15min), Nightly (comprehensive), Weekly (deep validation)
4. **Coverage**: ~26 test scenarios prioritized P0-P3 with risk-based classification
5. **Quality gates**: P0 100% pass, P1 ≥95%, coverage ≥80%

**What we need from team:** Just review and acknowledge (we already have the solution).

---

## For Architects and Devs - Open Topics 👷

### Risk Assessment

**Total risks identified**: 6 (2 high-priority score ≥6, 1 medium, 3 low)

#### High-Priority Risks (Score ≥6) - IMMEDIATE ATTENTION

| Risk ID    | Category  | Description   | Probability | Impact | Score       | Mitigation            | Owner   | Timeline |
| ---------- | --------- | ------------- | ----------- | ------ | ----------- | --------------------- | ------- | -------- |
| **DATA-001** | **DATA** | JSON file corruption | 2 | 3 | **6** | JSON validation, backup tests | Dev Team | Implementation |
| **BUS-001** | **BUS** | Data loss during file operations | 2 | 3 | **6** | Atomic file operations | Dev Team | Implementation |

#### Medium-Priority Risks (Score 3-5)

| Risk ID | Category | Description   | Probability | Impact | Score   | Mitigation   | Owner   |
| ------- | -------- | ------------- | ----------- | ------ | ------- | ------------ | ------- |
| PERF-001  | PERF    | Performance with large task lists | 2 | 2 | 4 | Benchmark tests | Dev Team |

#### Low-Priority Risks (Score 1-2)

| Risk ID | Category | Description   | Probability | Impact | Score   | Action  |
| ------- | -------- | ------------- | ----------- | ------ | ------- | ------- |
| SEC-001  | SEC    | File permission issues | 1 | 2 | 2 | Monitor |
| OPS-001  | OPS    | Cross-platform deployment | 1 | 1 | 1 | Monitor |
| TECH-001  | TECH    | Dependency breaking changes | 1 | 2 | 2 | Monitor |

#### Risk Category Legend

- **TECH**: Technical/Architecture (flaws, integration, scalability)
- **SEC**: Security (access controls, auth, data exposure)
- **PERF**: Performance (SLA violations, degradation, resource limits)
- **DATA**: Data Integrity (loss, corruption, inconsistency)
- **BUS**: Business Impact (UX harm, logic errors, revenue)
- **OPS**: Operations (deployment, config, monitoring)

---

### Testability Concerns and Architectural Gaps

**🚨 ACTIONABLE CONCERNS - Architecture Team Must Address**

#### 1. Blockers to Fast Feedback (WHAT WE NEED FROM ARCHITECTURE)

| Concern            | Impact              | What Architecture Must Provide         | Owner  | Timeline    |
| ------------------ | ------------------- | -------------------------------------- | ------ | ----------- |
| **Interactive input dependency** | Cannot automate `addTask` testing | Extract `bufio.Reader` into interface for mocking | Dev Team | Pre-implementation |
| **Hard-coded storage path** | Tests require specific home directory setup | Make storage path configurable via parameter | Dev Team | Pre-implementation |

#### 2. Architectural Improvements Needed (WHAT SHOULD BE CHANGED)

1. **Interface extraction for CLI input**
   - **Current problem**: `addTask` directly uses `bufio.NewReader(os.Stdin)`
   - **Required change**: Extract to `InputReader` interface for test mocking
   - **Impact if not fixed**: Core functionality cannot be automatically tested
   - **Owner**: Dev Team
   - **Timeline**: Pre-implementation

2. **Configurable storage location**
   - **Current problem**: Hard-coded `~/.task-tracker` path
   - **Required change**: Accept storage path via parameter or environment variable
   - **Impact if not fixed**: Tests require specific filesystem setup
   - **Owner**: Dev Team
   - **Timeline**: Pre-implementation

---

### Testability Assessment Summary

**📊 CURRENT STATE - FYI**

#### What Works Well

- ✅ Pure Go implementation (no external services to mock)
- ✅ Simple JSON storage format (easy to validate)
- ✅ Clear separation of concerns (storage, commands, models)
- ✅ Deterministic file-based operations (reproducible tests)

#### Accepted Trade-offs (No Action Required)

For Task Tracker v1.0, the following trade-offs are acceptable:

- **File system dependencies** - Acceptable for CLI tool; integration tests will handle
- **No external service mocks** - Not needed for local-only application
- **Manual test data setup** - Acceptable for simple file-based storage

---

### Risk Mitigation Plans (High-Priority Risks ≥6)

**Purpose**: Detailed mitigation strategies for all 2 high-priority risks (score ≥6). These risks MUST be addressed before release.

#### DATA-001: JSON file corruption (Score: 6) - HIGH

**Mitigation Strategy:**

1. Implement JSON schema validation before file writes
2. Add backup file creation before modifications
3. Create recovery tests for corrupted files

**Owner:** Dev Team
**Timeline:** Implementation phase
**Status:** Planned
**Verification:** Backup/restore tests pass, corruption detection works

#### BUS-001: Data loss during file operations (Score: 6) - HIGH

**Mitigation Strategy:**

1. Implement atomic file writes (write to temp, then rename)
2. Add file integrity checks (checksums)
3. Test crash scenarios during write operations

**Owner:** Dev Team
**Timeline:** Implementation phase
**Status:** Planned
**Verification:** Atomic operation tests pass, crash recovery validated

---

### Assumptions and Dependencies

#### Assumptions

1. Single-user local filesystem access is sufficient
2. Go 1.21+ runtime environment available
3. Standard filesystem permissions are adequate

#### Dependencies

1. Go standard library - Always available
2. urfave/cli/v2 framework - Stable dependency
3. github.com/google/uuid - External dependency (stable)

#### Risks to Plan

- **Risk**: Interactive input refactoring may take longer than expected
  - **Impact**: Delays test automation for core functionality
  - **Contingency**: Create wrapper functions initially, refactor later

---

**End of Architecture Document**

**Next Steps for Architecture Team:**

1. Review Quick Guide (🚨/⚠️/📋) and prioritize blockers
2. Assign owners and timelines for high-priority risks (≥6)
3. Validate assumptions and dependencies
4. Provide feedback to QA on testability gaps

**Next Steps for QA Team:**

1. Wait for pre-implementation blockers to be resolved
2. Refer to companion QA doc (test-design-qa.md) for test scenarios
3. Begin test infrastructure setup (factories, fixtures, environments)

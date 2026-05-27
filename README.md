# Keen

**The complete Atlassian platform CLI.** 1,091 API endpoints across Jira, Confluence, Agile, Marketplace, Service Desk, and Admin — in a single Go binary. Built for AI agents first, humans second.

Built by [Imogen Labs](https://imogenlabs.ai) using [CLI Printing Press](https://github.com/mvanhorn/cli-printing-press).

## Why

Every Atlassian workflow we automate used to be a wall of `curl` calls with hand-rolled JSON payloads and fragile `jq` parsing. This CLI replaces all of that with typed commands, structured output (`--agent` mode), local SQLite caching, and offline search.

## API Coverage

| API | Endpoints | What you get |
|-----|-----------|-------------|
| **Jira Platform v3** | 692 | Issues, transitions, search, bulk ops, fields, workflows, permissions |
| **Jira Agile** | 105 | Sprints, boards, backlogs, epics, ranking, velocity |
| **Confluence** | 100 | Pages, blog posts, spaces, search, templates, permissions |
| **Marketplace v2** | 149 | Addons, versions, vendors, distribution, reviews, reporting |
| **Service Desk** | 70 | Requests, queues, customers, SLAs, approvals |
| **Admin Orgs** | 61 | Organizations, directories, groups, policies, domains |

## Install

### From source (recommended)

```bash
git clone https://github.com/smashingtags/jira-pp-cli.git
cd jira-pp-cli
go build -o keen ./cmd/jira-pp-cli-pp-cli/
```

### MCP Server

```bash
go build -o keen-mcp ./cmd/jira-pp-cli-pp-mcp/
```

## Auth

```bash
export JIRA_CLOUD_PLATFORM_USERNAME="you@example.com"
export JIRA_CLOUD_PLATFORM_PASSWORD="your-api-token"
```

Generate a token at [id.atlassian.com/manage-profile/security/api-tokens](https://id.atlassian.com/manage-profile/security/api-tokens).

Verify: `keen doctor`

## Quick Start

```bash
# Search issues
keen jira-cloud-platform-search and-reconsile-issues-using-jql-post \
  --jql "project = EOS AND status = 'In Progress'" \
  --fields '["key","summary","status","assignee"]' --agent

# Transition an issue to Done
keen issue transitions do EOS-42 --transition-id 41 --agent

# List sprints
keen agile get-all-sprints 216 --state active --agent

# Create a sprint
keen agile create-sprint --stdin <<< '{"name":"Sprint 5","originBoardId":216}'

# Move issues into a sprint
keen agile move-issues-to-sprint-and-rank 801 \
  --stdin <<< '{"issues":["EOS-42","EOS-43"]}'

# Search Confluence
keen wiki search-content-by-cql --cql "type=page AND space=EOS" --limit 10 --agent

# Create a Confluence page
keen wiki create-or-update-content --stdin <<< '{
  "type": "page",
  "title": "Sprint 5 Review",
  "space": {"key": "EOS"},
  "body": {"storage": {"value": "<p>Content here</p>", "representation": "storage"}}
}'

# Get a Marketplace addon
keen addons get io.imogenlabs.story-points-auditor --agent

# Check the weather (why not)
keen which "weather"
```

## Agent Mode

Every command supports `--agent` which sets `--json --compact --no-input --no-color` in one flag. Designed for AI agent consumption — structured JSON output, no interactive prompts, minimal token usage.

```bash
# Agent-friendly output
keen agile get-all-boards --agent --compact

# Pipe to jq
keen issue transitions get EOS-42 --agent | jq '.results.transitions[] | {id, name}'
```

## Key Flags

| Flag | What it does |
|------|-------------|
| `--agent` | JSON + no prompts + compact (use this for all scripted calls) |
| `--compact` | Only key fields (id, name, status, timestamps) |
| `--json` | Force JSON output |
| `--csv` | CSV output |
| `--dry-run` | Show what would be sent without sending |
| `--stdin` | Read request body as JSON from stdin |
| `--select "id,name"` | Pick specific fields from response |
| `--quiet` | One value per line |
| `--data-source local` | Query local SQLite cache only |

## Command Discovery

```bash
# Find commands by keyword
keen which "sprint"
keen which "create page"
keen which "addon version"
keen which "board backlog"

# Browse all API interfaces
keen api

# Browse a specific interface
keen api wiki
keen api servicedeskapi
keen api agile
```

## Common Workflows

### Sprint Management

```bash
# List boards
keen agile get-all-boards --agent

# Get sprints for a board
keen agile get-all-sprints BOARD_ID --state active,future --agent

# Create sprint
keen agile create-sprint --stdin <<< '{
  "name": "Sprint 5",
  "originBoardId": 216,
  "startDate": "2026-06-01T09:00:00.000Z",
  "endDate": "2026-06-14T17:00:00.000Z",
  "goal": "Ship feature X"
}'

# Move issues to sprint
keen agile move-issues-to-sprint-and-rank SPRINT_ID \
  --stdin <<< '{"issues": ["EOS-42", "EOS-43", "EOS-44"]}'

# Start sprint (partial update)
keen agile partially-update-sprint SPRINT_ID \
  --stdin <<< '{"state": "active"}'

# Close sprint
keen agile partially-update-sprint SPRINT_ID \
  --stdin <<< '{"state": "closed"}'

# Get issues in sprint
keen agile get-issues-for-sprint SPRINT_ID --agent
```

### Issue Lifecycle

```bash
# Search
keen jira-cloud-platform-search and-reconsile-issues-using-jql-post \
  --jql "project = EOS AND status != Done ORDER BY rank" --agent

# Get available transitions
keen issue transitions get EOS-42 --agent

# Transition (e.g., to Done)
keen issue transitions do EOS-42 --transition-id 41 --agent

# Edit fields
keen issue edit EOS-42 --fields '{"priority":{"name":"High"}}' --agent

# Add comment
keen issue comment create EOS-42 --stdin <<< '{
  "body": {"type":"doc","version":1,"content":[
    {"type":"paragraph","content":[{"type":"text","text":"Done. Deployed to prod."}]}
  ]}
}'
```

### Confluence Documentation

```bash
# List spaces
keen wiki get-spaces --agent

# Search pages
keen wiki search-content-by-cql --cql "type=page AND space=EOS" --agent

# Create page
keen wiki create-or-update-content --stdin <<< '{
  "type": "page",
  "title": "Architecture Decision Record: API Gateway",
  "space": {"key": "EOS"},
  "body": {"storage": {"value": "<h2>Context</h2><p>...</p>", "representation": "storage"}}
}'

# Create blog post (sprint review, retro)
keen wiki create-or-update-content --stdin <<< '{
  "type": "blogpost",
  "title": "EOS Sprint 5 Review - Jun 1-14",
  "space": {"key": "EOS"},
  "body": {"storage": {"value": "<h2>Sprint Goal</h2><p>...</p>", "representation": "storage"}}
}'
```

### Marketplace Management

```bash
# Get addon details
keen addons get io.imogenlabs.story-points-auditor --agent

# List vendor's addons
keen addons get-vendor 183012220 --agent

# Search marketplace
keen addons list-search --q "story points" --agent

# Get addon versions
keen addons versions get io.imogenlabs.story-points-auditor --agent

# Create new version
keen addons versions create io.imogenlabs.story-points-auditor --stdin < version.json
```

## Local Sync & Offline Search

```bash
# Sync all Jira data to local SQLite
keen sync

# Search locally (no API calls)
keen search "authentication bug" --data-source local

# Analytics on synced data
keen analytics

# Find stale items
keen stale --days 30

# Find orphaned items
keen orphans

# Workload distribution
keen load
```

## Regeneration

Generated from 6 official Atlassian OpenAPI specs using [CLI Printing Press](https://github.com/mvanhorn/cli-printing-press) v4.15.0. Source specs stored in `specs/`.

```bash
cli-printing-press generate \
  --spec specs/jira-platform.json \
  --spec specs/jira-agile.json \
  --spec specs/confluence.json \
  --spec specs/marketplace.json \
  --spec specs/jira-service-desk.json \
  --spec specs/admin-orgs.json \
  --name jira-pp-cli --owner smashingtags --force --lenient \
  --auth-preference basicAuth --max-resources 1500
```

**Do not hand-edit generated Go files.** They are regenerated from specs.

## Deployed On

| Machine | Platform | Path |
|---------|----------|------|
| Mac Mini (dev workstation) | macOS arm64 | `~/repos/jira-pp-cli/` + `~/go/bin/` |
| iMac (build machine) | macOS arm64 | `~/repos/jira-pp-cli/` + `~/go/bin/` |
| MacBook Pro | macOS arm64 | `~/repos/jira-pp-cli/` + `~/go/bin/` |
| Windows Desktop | Windows amd64 | `C:\Users\micha\go\bin\` |

## License

Apache-2.0

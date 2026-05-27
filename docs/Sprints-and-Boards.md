# Sprints & Boards

## Boards

```bash
# List all boards
atlas-cli agile get-all-boards --agent

# Get board details
atlas-cli agile get-board 216 --agent

# Get board configuration
atlas-cli agile get-configuration 216 --agent

# Get board backlog
atlas-cli agile get-issues-for-backlog 216 --agent
```

## Sprints

```bash
# List sprints (active, future, closed)
atlas-cli agile get-all-sprints 216 --state active,future --agent

# Create sprint
atlas-cli agile create-sprint --stdin <<< '{"name":"Sprint 5","originBoardId":216}'

# Start sprint
atlas-cli agile partially-update-sprint 801 --stdin <<< '{"state":"active"}'

# Close sprint
atlas-cli agile partially-update-sprint 801 --stdin <<< '{"state":"closed"}'
```

## Moving Issues

```bash
# Move issues to sprint
atlas-cli agile move-issues-to-sprint-and-rank 801 --stdin <<< '{"issues":["EOS-42","EOS-43"]}'

# Move to backlog
atlas-cli agile move-issues-to-backlog --stdin <<< '{"issues":["EOS-42"]}'

# Get sprint issues
atlas-cli agile get-issues-for-sprint 801 --agent
```

## Board IDs (Imogen Labs)

| Project | Board ID |
|---------|----------|
| MPAC | 294 |
| EOS | 216 |
| EC | 217 |
| EBS | 218 |
| EA | 219 |
| SITE | 220 |
| MOB | 221 |
| INFRA | 222 |
| TMPL | 223 |
| EP | 224 |

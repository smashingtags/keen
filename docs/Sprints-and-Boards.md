# Sprints & Boards

## Boards

```bash
# List all boards
keen agile get-all-boards --agent

# Get board details
keen agile get-board 216 --agent

# Get board configuration
keen agile get-configuration 216 --agent

# Get board backlog
keen agile get-issues-for-backlog 216 --agent
```

## Sprints

```bash
# List sprints (active, future, closed)
keen agile get-all-sprints 216 --state active,future --agent

# Create sprint
keen agile create-sprint --stdin <<< '{"name":"Sprint 5","originBoardId":216}'

# Start sprint
keen agile partially-update-sprint 801 --stdin <<< '{"state":"active"}'

# Close sprint
keen agile partially-update-sprint 801 --stdin <<< '{"state":"closed"}'
```

## Moving Issues

```bash
# Move issues to sprint
keen agile move-issues-to-sprint-and-rank 801 --stdin <<< '{"issues":["EOS-42","EOS-43"]}'

# Move to backlog
keen agile move-issues-to-backlog --stdin <<< '{"issues":["EOS-42"]}'

# Get sprint issues
keen agile get-issues-for-sprint 801 --agent
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

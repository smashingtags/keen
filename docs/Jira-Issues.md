# Jira Issues

## Search (JQL)

```bash
# Basic search
atlas-cli jira-cloud-platform-search and-reconsile-issues-using-jql-post \
  --jql "project = EOS AND status = 'In Progress'" --agent

# With specific fields
atlas-cli jira-cloud-platform-search and-reconsile-issues-using-jql-post \
  --jql "project = MPAC ORDER BY rank" \
  --fields '["key","summary","status","priority","customfield_10039"]' \
  --max-results 50 --agent

# All issues in a sprint
atlas-cli jira-cloud-platform-search and-reconsile-issues-using-jql-post \
  --jql "sprint = 801" --agent
```

## Get Issue

```bash
atlas-cli issue get EOS-42 --agent
```

## Edit Issue

```bash
# Change priority
atlas-cli issue edit EOS-42 --fields '{"priority":{"name":"High"}}' --agent

# Set story points (customfield_10039)
atlas-cli issue edit EOS-42 --fields '{"customfield_10039":3}' --agent

# Set resolution
atlas-cli issue edit EOS-42 --fields '{"resolution":{"name":"Done"}}' --agent
```

## Transitions

```bash
# Get available transitions
atlas-cli issue transitions get EOS-42 --agent

# Transition to a status (by transition ID)
atlas-cli issue transitions do EOS-42 --transition-id 41 --agent
```

Common transition IDs (EOS/MPAC workflow):
- 11: Backlog
- 21: Selected for Development
- 31: In Progress
- 41: Done
- 51: QA
- 61: QA Passed
- 71: Flopped
- 81: Blocked

## Comments

```bash
# Add comment (ADF format)
atlas-cli issue comment create EOS-42 --stdin <<< '{
  "body": {
    "type": "doc",
    "version": 1,
    "content": [{"type":"paragraph","content":[{"type":"text","text":"Deployed to prod."}]}]
  }
}'
```

## Bulk Operations

```bash
# Bulk edit
atlas-cli bulk submit-edit --stdin < bulk-edit.json

# Bulk transition
atlas-cli bulk submit-transition --stdin < bulk-transition.json
```

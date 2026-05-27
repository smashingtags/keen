# Confluence

## Search

```bash
# Search by CQL
keen wiki search-content-by-cql --cql "type=page AND space=EOS" --limit 10 --agent

# Full-text search
keen wiki search-by-cql --cql "text ~ 'deployment'" --agent
```

## Spaces

```bash
# List spaces
keen wiki get-spaces --agent
```

## Pages

```bash
# Create page
keen wiki create-or-update-content --stdin <<< '{
  "type": "page",
  "title": "My New Page",
  "space": {"key": "EOS"},
  "body": {
    "storage": {
      "value": "<h2>Overview</h2><p>Content here</p>",
      "representation": "storage"
    }
  }
}'

# Create blog post
keen wiki create-or-update-content --stdin <<< '{
  "type": "blogpost",
  "title": "Sprint 5 Retrospective - Jun 1-14",
  "space": {"key": "EOS"},
  "body": {
    "storage": {
      "value": "<h2>What went well</h2><p>...</p><h2>What to improve</h2><p>...</p>",
      "representation": "storage"
    }
  }
}'

# Get page by ID
keen wiki get-content-by-id 12345 --agent

# Copy page
keen wiki copy-page --stdin <<< '{"copyAttachments":true,"copyPermissions":true,"destination":{"type":"parent_page","value":"67890"}}'
```

## Labels

```bash
# Add labels
keen wiki add-labels-to-content --stdin <<< '[{"prefix":"global","name":"sprint-review"}]'

# Get labels for space
keen wiki get-labels-for-space EOS --agent
```

## Confluence Spaces (Imogen Labs)

| Key | Name |
|-----|------|
| EOS | Eight.ly OS |
| EBS | Eight.ly Backup Server |
| EA | Eight.ly Agent |
| SITES | Marketing Sites |
| MOB | Mobile |
| TMPL | Templates |
| INFRA | Infrastructure |
| EP | Eight.ly Professional |
| HC | HomelabARR CE |
| ENGOPS | Engineering Operations (SDLC) |

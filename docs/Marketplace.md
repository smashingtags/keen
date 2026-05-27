# Marketplace

## Addons

```bash
# Get addon
atlas-cli addons get io.imogenlabs.story-points-auditor --agent

# List vendor's addons
atlas-cli addons get-vendor 183012220 --agent

# Search marketplace
atlas-cli addons list-search --q "story points" --agent

# Create addon
atlas-cli addons create --stdin < addon.json

# Update addon (JSON Patch)
atlas-cli addons update io.imogenlabs.story-points-auditor --stdin < patch.json
```

## Versions

```bash
# Get versions
atlas-cli addons versions get io.imogenlabs.story-points-auditor --agent

# Create version
atlas-cli addons versions create io.imogenlabs.story-points-auditor --stdin < version.json
```

## Distribution

```bash
# Get distribution status
atlas-cli addons distribution get io.imogenlabs.story-points-auditor --agent
```

## Imogen Labs Apps

| App | Key | Marketplace |
|-----|-----|-------------|
| Point Health | io.imogenlabs.story-points-auditor | [3545512218](https://marketplace.atlassian.com/apps/3545512218) |
| Ready Check | io.imogenlabs.dor-gate-enforcer | [1212055508](https://marketplace.atlassian.com/apps/1212055508) |
| Sprint Spillover | io.imogenlabs.sprint-carry-over-manager | [2973651844](https://marketplace.atlassian.com/apps/2973651844) |
| CommitGuard | io.imogenlabs.smart-commit-validator | [3374376546](https://marketplace.atlassian.com/apps/3374376546) |
| Sprint Health Dashboard | io.imogenlabs.sprint-health-dashboard | sharing enabled |
| Retro Automation | io.imogenlabs.retro-automation | sharing enabled |
| Dependency Mapper | io.imogenlabs.dependency-mapper | sharing enabled |
| Bulk Field Editor | io.imogenlabs.bulk-field-editor | sharing enabled |
| Workflow Transition Auditor | io.imogenlabs.workflow-transition-auditor | sharing enabled |

Vendor: **Imogen Labs** (183012220)

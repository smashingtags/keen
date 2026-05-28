// Local customization: top-level shortcut aliases for common agent operations.
// See .printing-press-patches.json for rationale.

package cli

import (
	"github.com/spf13/cobra"
)

func addAliasCommands(rootCmd *cobra.Command, flags *rootFlags) {
	// keen find-issues → keen jira-cloud-platform-search and-reconsile-issues-using-jql-post
	// Wraps with default fields so --agent returns key/summary/status, not just id
	findCmd := newJiraCloudPlatformSearchAndReconsileIssuesUsingJqlPostCmd(flags)
	origRunE := findCmd.RunE
	findCmd.RunE = func(cmd *cobra.Command, args []string) error {
		if !cmd.Flags().Changed("fields") && flags.compact {
			_ = cmd.Flags().Set("fields", `["key","summary","status","priority","issuetype","project"]`)
		}
		return origRunE(cmd, args)
	}
	findCmd.Use = "find-issues"
	findCmd.Aliases = []string{"fi"}
	findCmd.Short = "Search Jira issues by JQL (includes key/summary/status by default)"
	findCmd.Example = "  keen find-issues --jql \"project = INFRA AND status = 'In Progress'\" --agent"
	rootCmd.AddCommand(findCmd)

	// keen transition → keen issue transitions do
	transitionCmd := newIssueTransitionsDoCmd(flags)
	transitionCmd.Use = "transition [issueKey]"
	transitionCmd.Short = "Transition a Jira issue"
	transitionCmd.Example = "  keen transition INFRA-46 --transition-id 51 --agent"
	rootCmd.AddCommand(transitionCmd)

	// keen sprint → keen agile get-all-sprints
	sprintListCmd := newAgileGetAllSprintsCmd(flags)
	sprintListCmd.Use = "sprints [boardId]"
	sprintListCmd.Short = "List sprints for a board"
	sprintListCmd.Example = "  keen sprints 222 --agent"
	rootCmd.AddCommand(sprintListCmd)

	// keen get → keen issue get
	getIssueCmd := newIssueGetCmd(flags)
	getIssueCmd.Use = "get [issueKey]"
	getIssueCmd.Short = "Get a Jira issue"
	getIssueCmd.Example = "  keen get INFRA-46 --agent"
	rootCmd.AddCommand(getIssueCmd)
}

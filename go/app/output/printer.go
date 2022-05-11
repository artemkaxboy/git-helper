package output

import (
	"app/common"
	"app/git"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"os"
	"sort"
	"strconv"
	"time"
)

func PrintRemoteBranches(remote string, branches []*git.RemoteBranch) {

	headers := git.RemoteBranchHeaders()
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	table.SetAutoFormatHeaders(false)
	table.SetAutoWrapText(false)

	sort.Slice(branches, func(i, j int) bool {
		return branches[i].LastCommitTime(time.Now()).After(branches[j].LastCommitTime(time.Now()))
	})

	branchRows := make([][]string, len(branches))
	for i, branch := range branches {
		branchRows[i] = branch.ToRow()
	}

	table.AppendBulk(branchRows)

	fmt.Printf("%v:\n", remote)
	table.Render()

}

func PrintAuthors(remote string, authors []common.StringIntPair) {

	headers := []string{"#", "Author", "Branches"}
	table := tablewriter.NewWriter(os.Stdout)

	sort.Slice(authors, func(i, j int) bool {
		return authors[i].Value > authors[j].Value
	})

	authorRows := make([][]string, len(authors))
	sum := 0
	for i, author := range authors {

		sum += author.Value
		row := []string{strconv.Itoa(i + 1), author.Key, strconv.Itoa(author.Value)}
		authorRows[i] = common.RoundStrings(row, []int{5, 60, 5})
	}

	table.SetHeader(headers)
	table.SetAutoFormatHeaders(false)
	table.SetAutoWrapText(false)
	table.AppendBulk(authorRows)
	table.SetFooterAlignment(tablewriter.ALIGN_RIGHT)
	table.SetFooter([]string{"", fmt.Sprintf("Sum of %d items", table.NumLines()), strconv.Itoa(sum)})

	fmt.Printf("Remote = %v:\n", remote)
	table.Render()
}

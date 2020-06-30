package parser

import (
	"regexp"
	"strings"
)

func ParseInnoDBDeadlocks(innodbStatusText string) map[string][]Deadlock {
	results := make(map[string][]Deadlock)

	for _, dl := range parseDeadlocks(innodbStatusText) {
		if _, ok := results[dl.GroupKey]; !ok {
			results[dl.GroupKey] = make([]Deadlock, 0)
		}

		results[dl.GroupKey] = append(results[dl.GroupKey], dl)
	}

	return results
}

func parseDeadlocks(data string) []Deadlock {
	deadlocks := make([]Deadlock, 0)

	sectionContent := findDeadlockSection(data)
	compile := regexp.MustCompile(`(\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}) ([A-Za-z0-9]{12})`)

	segsByDate := compile.Split(sectionContent, -1)
	for i, s := range compile.FindAllString(sectionContent, -1) {
		for _, dl := range extractDeadlock(segsByDate[i+1])[1:] {
			deadlocks = append(deadlocks, Deadlock{
				GroupKey: s,
				Sections: extractDeadlockSegs(dl),
			})
		}

	}
	return deadlocks
}

type Deadlock struct {
	GroupKey string
	Sections []Section
}

type Section struct {
	Title   string
	Content string
}

func extractDeadlockSegs(dl string) []Section {
	compile := regexp.MustCompile(`\*\*\* \(\d+\) `)
	segs := compile.Split(dl, -1)
	segs[0] = "TRANSACTION" + segs[0]

	sections := make([]Section, 0)
	for _, s := range segs {
		ss := strings.SplitN(s, "\n", 2)
		sections = append(sections, Section{
			Title:   strings.TrimSuffix(ss[0], ":"),
			Content: ss[1],
		})
	}

	return sections
}

func extractDeadlock(deadlockSection string) []string {
	compile := regexp.MustCompile(`\*\*\* \(\d+\) TRANSACTION`)
	return compile.Split(deadlockSection, -1)
}

func findDeadlockSection(data string) string {
	compile := regexp.MustCompile(`---+\n([A-Z /]+)\n---+\n`)
	index := deadlockIndex(compile, data)

	return compile.Split(data, -1)[index+1]
}

func deadlockIndex(compile *regexp.Regexp, data string) int {
	for i, s := range compile.FindAllStringSubmatch(data, -1) {
		for _, ss := range s {
			if ss == "LATEST DETECTED DEADLOCK" {
				return i
			}
		}
	}
	return -1
}

// Tool to generate benchmark ranks and output back to README using templating
package rank

import (
	"bytes"
	"fmt"
	"golang.org/x/tools/benchmark/parse"
	"io"
	"os"
	"regexp"
	"sort"
	"strings"
	"text/template"
)

// Maps of library names (eg "ffjson") to a map of actions (e.g. "Marshal").
// Actions map to a list of recorded results (one for each data fixture).
type BenchmarkMap map[string]map[string][]int

// ParseBenchmarkSuite parses out each test line from the benchmarking
// output and reorders it by library. For each library, the results are
// ordered by action.
func ParseBenchmarkSuite(r io.Reader) *BenchmarkMap {

	lines, _ := parse.ParseSet(r)
	parsed := BenchmarkMap{}
	re1, _ := regexp.Compile("/(\\w+)")
	re2, _ := regexp.Compile("([A-Z][a-z]+)")

	for name, data := range lines {
		matches := re1.FindAllStringSubmatch(name, -1)
		library, action := strings.ToLower(matches[0][1]), matches[1][1]
		action = re2.FindStringSubmatch(action)[1]
		if _, ok := parsed[library]; !ok {
			parsed[library] = make(map[string][]int)
		}
		parsed[library][action] = append(parsed[library][action], int(data[0].NsPerOp)) // TODO right now only grabbing NsPrOp, but need to extend to other metrics
	}
	return &parsed
}

type RankData struct {
	Library string
	Value   int
}

// Implements sort.Interface
type RankTable []RankData

func (t RankTable) Len() int           { return len(t) }
func (t RankTable) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t RankTable) Less(i, j int) bool { return t[i].Value < t[j].Value }

// BuildRankTable transforms a BenchmarkMap
func BuildRankTable(parsed *BenchmarkMap) map[string]RankTable {
	tables := make(map[string]RankTable)
	for library, data := range *parsed {
		for benchmark, result := range data {
			sum, size := 0, 0
			for i, num := range result {
				sum += num
				size = i
			}
			tables[benchmark] = append(tables[benchmark], RankData{library, sum / size})
		}
	}
	for _, table := range tables {
		sort.Sort(table)
	}
	return tables
}

// README template data
type READMEData struct {
	Action string
	Chart  string
}

func BuildRankTableString(t RankTable) string {
	var b bytes.Buffer
	for i, row := range t {
		fmt.Fprintf(&b, "%d | %v | %d ns/op\n", i+1, row.Library, row.Value)
	}
	return b.String()
}

func GenerateREADME(dataPath, templatePath, readmePath string) {
	tables := make([]READMEData, 2)

	df, err := os.Open(dataPath)
	if err != nil {
		panic(err)
	}
	defer df.Close()

	data := ParseBenchmarkSuite(df)
	ranks := BuildRankTable(data)
	i := 0
	for action, rows := range ranks {
		chart := BuildRankTableString(rows)
		tables[i] = READMEData{action, chart}
		i += 1
	}

	tf, err := template.ParseFiles(templatePath)
	if err != nil {
		panic(err)
	}
	rf, err := os.Create(readmePath)
	if err != nil {
		panic(err)
	}
	err = tf.Execute(rf, tables)
	if err != nil {
		panic(err)
	}
	rf.Close()
}

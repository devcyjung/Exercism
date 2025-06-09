package tournament

import (
    "bufio"
    "cmp"
    "errors"
    "fmt"
    "io"
    "slices"
    "strings"
    "text/tabwriter"
)

const (
    rowFormat = "%v\t\t\t|%v\t|%v\t|%v\t|%v\t|%v\n"
    header1 = "Team"
    header2 = "MP"
    header3 = "W"
    header4 = "D"
    header5 = "L"
    header6 = "P"
)

var InvalidInputError = errors.New("Invalid input format")

type score struct {
    win, draw, loss, match, point int
    name string
}

type scoreMap map[string]*score

func Tally(reader io.Reader, writer io.Writer) error {
	scoreboard, err := generateScoreboard(reader)
    sortedScoreList := sortedScores(scoreboard)
    err = errors.Join(err, writeScores(writer, sortedScoreList))
    return err
}

func generateScoreboard(reader io.Reader) (scoreMap, error) {
    scanner := bufio.NewScanner(reader)
    scoreboard := make(scoreMap)
    var trimmed string
    var fields []string
    var err error
    var ok bool
    for scanner.Scan() {
        trimmed = strings.TrimSpace(scanner.Text())
        if len(trimmed) == 0 || strings.HasPrefix(trimmed, "#") {
            continue
        }
        fields = strings.FieldsFunc(trimmed, semiColonDelim)
        if len(fields) != 3 {
            err = errors.Join(err, InvalidInputError)
            continue
        }
        _, ok = scoreboard[fields[0]]
        if !ok {
            scoreboard[fields[0]] = &score{}
        }
        _, ok = scoreboard[fields[1]]
        if !ok {
            scoreboard[fields[1]] = &score{}
        }
        switch fields[2] {
        case "win":
            scoreboard[fields[0]].win++
            scoreboard[fields[1]].loss++
        case "loss":
            scoreboard[fields[0]].loss++
            scoreboard[fields[1]].win++
        case "draw":
            scoreboard[fields[0]].draw++
            scoreboard[fields[1]].draw++
        default:
            err = errors.Join(err, InvalidInputError)
        }
    }
    err = errors.Join(err, scanner.Err())
    return scoreboard, err
}

func semiColonDelim(r rune) bool {
    return r == ';'
}

func sortedScores(scoreboard scoreMap) []*score {
    scores := make([]*score, 0, len(scoreboard))
    for k, v := range scoreboard {
        v.match = v.win + v.draw + v.loss
        v.point = 3 * v.win + v.draw
        v.name = k
        scores = append(scores, v)
    }
    slices.SortFunc(scores, scoreSorter)
    return scores
}

func scoreSorter(a, b *score) int {
    if c := cmp.Compare(b.point, a.point); c != 0 {
        return c
    }
    return strings.Compare(a.name, b.name)
}

func writeScores(writer io.Writer, scores []*score) error {
    tabWriter := tabwriter.NewWriter(writer, 4, 0, 0, ' ', 0)
    err := writeRow(tabWriter, header1, header2, header3, header4, header5, header6)
    for _, sc := range scores {
        err = errors.Join(err, writeRow(tabWriter, sc.name, sc.match, sc.win, sc.draw, sc.loss, sc.point))
    }
    err = errors.Join(err, tabWriter.Flush())
    return err
}

func writeRow(writer io.Writer, col1, col2, col3, col4, col5, col6 any) error {
    _, err := fmt.Fprintf(writer, rowFormat, cell(col1), cell(col2), cell(col3), cell(col4), cell(col5), lastCell(col6))
    return err
}

func cell(value any) string {
    return center(value, 4)
}

func lastCell(value any) string {
    return strings.TrimRight(center(value, 4), " ")
}

func center(value any, width int) string {
    s := fmt.Sprintf("%v", value)
    return fmt.Sprintf("%*s", -width, fmt.Sprintf("%*s", width-(width - len(s))/2, s))
}
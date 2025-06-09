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

const rowFormat = "%v\t\t\t|%v\t|%v\t|%v\t|%v\t|%v\n"

var (
    InvalidInputError = errors.New("Invalid input format")
    delim = func(r rune) bool {
        return r == ';'
    }
    sorter = func(a, b *score) int {
        if c := cmp.Compare(b.point, a.point); c != 0 {
            return c
        }
        return strings.Compare(a.name, b.name)
    }
    center = func(value any, width int) string {
        s := fmt.Sprintf("%v", value)
        return fmt.Sprintf("%*s", -width, fmt.Sprintf("%*s", width-(width - len(s))/2, s))
    }
)

type score struct {
    win, draw, loss, match, point int
    name string
}

func Tally(reader io.Reader, writer io.Writer) (rErr error) {
    var fields []string
    var ok bool
    var err error
    var trimmed string
	scanner := bufio.NewScanner(reader)
    scoreboard := make(map[string]*score)
    for scanner.Scan() {
        trimmed = strings.TrimSpace(scanner.Text())
        if strings.HasPrefix(trimmed, "#") || len(trimmed) == 0 {
            continue
        }
        fields = strings.FieldsFunc(trimmed, delim)
        if len(fields) != 3 {
            rErr = errors.Join(rErr, InvalidInputError)
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
            rErr = errors.Join(rErr, InvalidInputError)
        }
    }
    rErr = errors.Join(rErr, scanner.Err())
    scores := make([]*score, 0, len(scoreboard))
    for k, v := range scoreboard {
        v.match = v.win + v.draw + v.loss
        v.point = 3 * v.win + v.draw
        v.name = k
        scores = append(scores, v)
    }
    slices.SortFunc(scores, sorter)
	tabWriter := tabwriter.NewWriter(writer, 4, 0, 0, ' ', 0)
    _, err = fmt.Fprintf(tabWriter, rowFormat,
    	center("Team", 4), center("MP", 4),
        center("W", 4), center("D", 4), center("L", 4), strings.TrimRight(center("P", 4), " "))
    rErr = errors.Join(rErr, err)
    for _, row := range scores {
        _, err = fmt.Fprintf(tabWriter, rowFormat,
        	center(row.name, 4), center(row.match, 4), center(row.win, 4),
        	center(row.draw, 4), center(row.loss, 4), strings.TrimRight(center(row.point, 4), " "))
        rErr = errors.Join(rErr, err)
    }
    rErr = errors.Join(rErr, tabWriter.Flush())
    return
}
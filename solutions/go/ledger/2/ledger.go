package ledger

import (
    "cmp"
	"errors"
    "fmt"
    "slices"
	"strconv"
	"strings"
    "time"
    "unicode/utf8"
)

type Entry struct {
	Date, Description	string
	Change				int
}

type localeType string
type currencyType string

const (
    headerFmt = "%-10s | %-25s | %s\n"
    rowFmt = "%-10s | %-25s | %13s\n"
    entryDateFmt = "2006-01-02"
    localeNL localeType = "nl-NL"
    localeUS localeType = "en-US"
    currencyEUR currencyType = "€"
    currencyUSD currencyType = "$"
)

var (
    ErrInvalidDateFormat = errors.New("incorrect date format")
    ErrUnsupportedLocale = errors.New("unsupported locale")
    ErrUnsupportedCurrency = errors.New("unsupported currency")
)

func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
    loc, ok := getLocaleType(locale)
    if !ok {
        return "", ErrUnsupportedLocale
    }
    cur, ok := getCurrencyType(currency)
    if !ok {
        return "", ErrUnsupportedCurrency
    }
    headers := getLocaleHeaders(loc)
    var b strings.Builder
    fmt.Fprintf(&b, headerFmt, headers[0], headers[1], headers[2])
    clonedEntries := slices.Clone(entries)
	slices.SortFunc(clonedEntries, getEntrySorter(loc))
    for _, entry := range clonedEntries {
        date, ok := getDate(loc, entry.Date)
        if !ok {
            return "", ErrInvalidDateFormat
        }
        summary := getSummary(entry.Description)
        change := getChange(loc, cur, entry.Change)
        fmt.Fprintf(&b, rowFmt, date, summary, change)
    }
    return b.String(), nil
}

func getLocaleType(locale string) (l localeType, ok bool) {
    switch locale {
    case "nl-NL":
        l, ok = localeNL, true
        return
    case "en-US":
        l, ok = localeUS, true
        return
    default:
        return
    }
}

func getCurrencyType(currency string) (c currencyType, ok bool) {
    switch currency {
    case "EUR":
        c, ok = currencyEUR, true
        return
    case "USD":
        c, ok = currencyUSD, true
        return
    default:
        return
    }
}

func getLocaleHeaders(locale localeType) []string {
    switch locale {
    case localeNL:
        return []string{"Datum", "Omschrijving", "Verandering"}
    case localeUS:
        return []string{"Date", "Description", "Change"}
    default:
        return []string{}
    }
}

func getEntrySorter(locale localeType) func(Entry, Entry) int {
    return func(a, b Entry) int {
        aDate, _ := time.Parse(entryDateFmt, a.Date)
        bDate, _ := time.Parse(entryDateFmt, b.Date)
        switch {
        case aDate.Before(bDate):
            return -1
        case aDate.After(bDate):
            return 1
        default:
            descCmp := cmp.Compare(a.Description, b.Description)
            if descCmp != 0 {
                return descCmp
            }
            return cmp.Compare(a.Change, b.Change)
        }
    }
}

func getLocaleDateFmt(locale localeType) string {
    switch locale {
    case localeNL:
        return "02-01-2006"
    case localeUS:
        return "01/02/2006"
    default:
        return ""
    }
}

func getDate(locale localeType, dateString string) (string, bool) {
    t, err := time.Parse(entryDateFmt, dateString)
    if err != nil {
        return "", false
    }
    return t.Format(getLocaleDateFmt(locale)), true
}

func getSummary(description string) string {
    c := utf8.RuneCountInString(description)
    if c <= 25 {
        return description
    }
    var b strings.Builder
    for i, r := range description {
        if i == 22 {
            break
        }
        b.WriteRune(r)
    }
    b.WriteString("...")
    return b.String()
}

func getChangeFmt(locale localeType, isPositive bool) string {
    switch {
    case locale == localeNL && !isPositive:
        return "%s %s-"
    case locale == localeNL && isPositive:
        return "%s %s "
    case locale == localeUS && !isPositive:
        return "(%s%s)"
    case locale == localeUS && isPositive:
        return "%s%s "
    default:
        return ""
    }
}

func getDigitSym(locale localeType) (string, string) {
    switch locale {
    case localeNL:
        return ".", ","
    case localeUS:
        return ",", "."
    default:
        return "", ""
    }
}

func getAbsChangeString(locale localeType, absCentAmount int) string {
    result := make([]string, 0, 5)
    cents := absCentAmount % 100
    absCentAmount /= 100
    if absCentAmount == 0 {
        result = append(result, "0")
    }
    for absCentAmount > 0 {
        result = append(result, strconv.Itoa(absCentAmount % 1000))
        absCentAmount /= 1000
    }
    slices.Reverse(result)
    thousands, decimal := getDigitSym(locale)
    return strings.Join(result, thousands) + decimal + fmt.Sprintf("%02d", cents)
}

func getChange(locale localeType, currency currencyType, centAmount int) string {
    if centAmount >= 0 {
        return fmt.Sprintf(getChangeFmt(locale, true), currency, getAbsChangeString(locale, centAmount))
    } else {
        return fmt.Sprintf(getChangeFmt(locale, false), currency, getAbsChangeString(locale, -centAmount))
    }
}
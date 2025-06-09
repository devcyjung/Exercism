package grep

import (
    "bufio"
    "fmt"
    "log/slog"
    "os"
    "regexp"
    "slices"
)

var logger = slog.New(slog.NewJSONHandler(os.Stderr, nil))

type ctxKey string

type matcher struct {
    re 		*regexp.Regexp
    invert	bool
}

type matcherOptions struct {
    caseInsensitive, invert, wholeLine bool
}

func newMatcher(pattern string, matchOpt matcherOptions) (*matcher, error) {
    m := &matcher{invert: matchOpt.invert}
    if matchOpt.caseInsensitive {
        pattern = "(?i)" + pattern
    }
    if matchOpt.wholeLine {
        pattern = "^" + pattern + "$"
    }
    re, err := regexp.Compile(pattern)
    if err != nil {
        return nil, err
    }
    m.re = re
    return m, nil
}

func (m *matcher) matchLine(line string) bool {
    return m.re.MatchString(line) != m.invert
}

type writer func(string, int, string) string

type writerOptions struct {
    writeLineNumber, onlyFileNames, writeFileNames bool
}

func newWriter(writeOpt writerOptions) writer {
    switch {
    case writeOpt.onlyFileNames:
        return func(file string, line int, str string) string {
            return file
        }
    case writeOpt.writeLineNumber && writeOpt.writeFileNames:
        return func(file string, line int, str string) string {
            return fmt.Sprintf("%v:%v:%v", file, line, str)
        }
    case writeOpt.writeLineNumber:
        return func(file string, line int, str string) string {
            return fmt.Sprintf("%v:%v", line, str)
        }
    case writeOpt.writeFileNames:
        return func(file string, line int, str string) string {
            return fmt.Sprintf("%v:%v", file, str)
        }
    default:
        return func(file string, line int, str string) string {
            return str
        }
    }
}

func Search(pattern string, flags, files []string) []string {
    matchOpt := matcherOptions{}
    writeOpt := writerOptions{}
    for _, flag := range flags {
        switch flag {
            case "-n":
            	writeOpt.writeLineNumber = true
            case "-l":
            	writeOpt.onlyFileNames = true
            case "-i":
            	matchOpt.caseInsensitive = true
            case "-v":
            	matchOpt.invert = true
            case "-x":
            	matchOpt.wholeLine = true
        }
    }
    writeOpt.writeFileNames = len(files) > 1
    lineMatcher, err := newMatcher(pattern, matchOpt)
    if err != nil {
        logger.Error("Regex generation failed", "error", err)
        return nil
    }
	lineWriter := newWriter(writeOpt)
    return matchResult(files, lineMatcher, lineWriter, writeOpt.onlyFileNames)
}

func matchResult(files []string, lineMatcher *matcher, lineWriter writer, shortCircuit bool) []string {
    result := make([]string, 0, 100)
    for _, filename := range files {
        file, err := os.Open(filename)
        if err != nil {
            logger.Error("File open failed", "error", err, "name", filename)
            continue
        }
        scanner := bufio.NewScanner(file)
        lineNumber := 1
        for scanner.Scan() {
            lineText := scanner.Text()
            if lineMatcher.matchLine(lineText) {
                result = append(result, lineWriter(filename, lineNumber, lineText))
                if shortCircuit {
                    break
                }
            }
            lineNumber++
        }
        if err = scanner.Err(); err != nil {
            logger.Error("File scan failed", "error", err, "name", filename)
        }
        if err = file.Close(); err != nil {
            logger.Error("File close failed", "error", err, "name", filename)
        }
    }
    slices.Clip(result)
    return result
}
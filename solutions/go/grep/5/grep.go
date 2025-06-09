package grep

import (
    "bufio"
    "log"
    "os"
    "regexp"
    "strconv"
    "strings"
)

func Search(pattern string, flags, files []string) (result []string) {
    var prependLineNumber, onlyFileNames, caseInsensitive, invert, wholeLineMatch bool
    
    for _, flag := range flags {
        switch flag {
            case "-n": prependLineNumber = true
            case "-l": onlyFileNames = true
            case "-i": caseInsensitive = true
            case "-v": invert = true
            case "-x": wholeLineMatch = true
        }
    }
    
    var file *os.File
    var err error
    var scanner *bufio.Scanner
    var line string
    var lineNumber int
    var shouldAppendResult bool
    resultElems := make([]string, 0, 3)
    isMultiFile := len(files) > 1

	if wholeLineMatch {
        pattern = "^" + pattern + "$"
    }
    
    if caseInsensitive {
        pattern = "(?i)" + pattern
    }
    
    re, err := regexp.Compile(pattern)
    if err != nil {
        log.Fatalf("error parsing regex pattern %v: %v", pattern, err)
    }
    re.Longest()
    
	for _, filename := range files {
        file, err = os.Open(filename)
        
		if err != nil {
            log.Fatalf("error opening %v: %v", filename, err)
        }
        
        lineNumber = 0
        scanner = bufio.NewScanner(file)
        
        for scanner.Scan() {
            lineNumber++
            line = scanner.Text()
            
            if !invert {
                shouldAppendResult = re.MatchString(line)
            } else {
                shouldAppendResult = !re.MatchString(line)
            }

            if !shouldAppendResult {
                continue
            }

            if onlyFileNames {
                result = append(result, filename)
                break
            }

            if isMultiFile {
                resultElems = append(resultElems, filename)
            }
            
            if prependLineNumber {
                resultElems = append(resultElems, strconv.Itoa(lineNumber))
            }

            resultElems = append(resultElems, line)
            result = append(result, strings.Join(resultElems, ":"))
            resultElems = resultElems[:0:3]
        }
        
        if err = scanner.Err(); err != nil {
            log.Fatalf("error reading %v: %v", filename, err)
        }

        err = file.Close()
        if err != nil {
            log.Fatalf("error closing %v: %v", filename, err)
        }
    }

    return
}

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
        "runtime"
        "runtime/pprof"
        "sort"
	"strings"
	"sync"
	"time"

	fits "github.com/astrogo/fitsio"
)

func getExptimeInFile(filePath string) (float64, error) {
	var exptimeValue float64

	fh, err := os.Open(filePath)
	if err != nil {
		return 0, fmt.Errorf("failed to open file %s: %v", filePath, err)
	}
	defer fh.Close()

	f, err := fits.Open(fh)
	if err != nil {
		return 0, fmt.Errorf("failed to process file %s: %v", filePath, err)
	}
	defer f.Close()

	hdr := f.HDU(0).Header()
	exptime := hdr.Get("EXPTIME")
	if exptime != nil {
		exptimeValue = exptime.Value.(float64)
	}

	return exptimeValue, nil
}

func worker(files <-chan string, results chan<- ExptimeResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for path := range files {
		exptime, err := getExptimeInFile(path)
		if err != nil {
			fmt.Printf("Error processing file %s: %v\n", path, err)
			continue
		}
		dir := filepath.Dir(path)
		results <- ExptimeResult{dir, exptime}
	}
}

func secondsToHoursMinutes(seconds float64) (int, int) {
	hours := int(seconds) / 3600
	minutes := (int(seconds) % 3600) / 60
	return hours, minutes
}

type ExptimeResult struct {
	Directory string
	Exptime   float64
}

func main() {
        // Flags
	cpuProfile := flag.String("cpuprofile", "", "write cpu profile to file")
	startDir := flag.String("dir", ".", "directory to start scanning")
	ignoreDir := flag.String("ignore", "", "directory tree to ignore")
        cpuThreads := flag.Int("threads", runtime.NumCPU(), "number of threads")
	flag.Parse()

	// Enable CPU profiling if specified
	if *cpuProfile != "" {
		f, err := os.Create(*cpuProfile)
		if err != nil {
			fmt.Printf("Error creating CPU profile: %v\n", err)
			return
		}
		defer f.Close()
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	startTime := time.Now()

	files := make(chan string, 100) // Buffered channel to hold file paths
	results := make(chan ExptimeResult, 100)

	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 0; i < *cpuThreads; i++ {
		wg.Add(1)
		go worker(files, results, &wg)
	}

	// This goroutine will walk the directory and send file paths to the channel
        go func() {
		filepath.Walk(*startDir, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			// Check if the path contains the ignore directory
			if *ignoreDir != "" && strings.Contains(path, *ignoreDir) {
				if info.IsDir() {
					return filepath.SkipDir // skip the entire directory
				}
				return nil // skip the current file
			}
			if !info.IsDir() && strings.HasSuffix(info.Name(), ".fits") {
				files <- path
			}
			return nil
		})
		close(files) // Close the channel once all file paths have been sent
	}()

	// This goroutine will collect the results
	go func() {
		wg.Wait()
		close(results) // Close the results channel once all workers are done
	}()

	directoryExptime := make(map[string]float64)
	for result := range results {
		directoryExptime[result.Directory] += result.Exptime
	}

        // Convert map keys to a slice for sorting
        var dirs []string
        for dir := range directoryExptime {
            dirs = append(dirs, dir)
        }

        // Sort the directory names
        sort.Strings(dirs)



	var totalExptime float64
	fmt.Printf("Exptime\t\tDirectory (hours:minutes)\n")
	fmt.Printf("-----------------------------------------------------\n")
        for _, dir := range dirs {
                exptime := directoryExptime[dir]
		hours, minutes := secondsToHoursMinutes(exptime)
		fmt.Printf("%d:%02d\t\t%s\n", hours, minutes, dir)
		totalExptime += exptime
	}

	totalHours, totalMinutes := secondsToHoursMinutes(totalExptime)
	fmt.Printf("\nTotal EXPTIME: %d hours and %02d minutes\n", totalHours, totalMinutes)

	// Profiling results
	elapsedTime := time.Since(startTime)
	fmt.Printf("\nTime taken: %v\n", elapsedTime)
}


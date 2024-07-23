// GOROUTINES MENTORING SESSION, STARTING POINT
// ━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
// ────
// Goals
// ────
// This is the starting point for a mentoring session in which we
// explore the implementation of a simple map-reduce pattern and an
// aggregator in Go, utilising goroutines, channels and wait groups.
//
// ────────
// Non-goals
// ────────
// The goal of this session is not to explore map-reduce, per se, but
// to gain experience with goroutines, channels and sync.
//
// This exercise is *not* intended as 'homework' but should be
// completed in a pairing session with you mentor. It is *explicitly*
// not a test.
//
// ───────────────────
// Purpose of the program
// ───────────────────
// This program we're working on is minimal word counter, you can
// invoke it by executing:
//
//    ./mapreduce <path to file>
//
// The output you see should be directly comparable to the output of
// the shell command:
//
//    wc -w <path to file>
//
// ───────────────
// Steps to complete
// ───────────────
// Despite the naming of the repository, and resulting binary, the
// starting condition is *not* a map-reduce pattern, but rather a
// simple, single-threaded imperative solution.
//
//  To convert this to a map-reduce pattern we must:
//
//   1. Make each output of scanner.Text() feed a channel that is
//   consumed by mapper function.
//
//   2. Make a mapper function that consumes lines of text from a
//   channel and splits them into slices of words, which it then
//   outputs on a channel consumed by a reducer function.
//
//   3. Make a reducer function that consumers slices of strings, and
//   outputs their lengths onto a channel for the aggregator function
//
//   4. Make an aggregator function that consumes a channel of
//   integers and sums them ready for printing.
//
// A worked example is available in the main branch of the repo:
//     github.com/tealeg/mapreduce
//
// ─────────────────────────────────────────────────────────
// Questions to answer after the exercise (together with your mentor)
// ─────────────────────────────────────────────────────────
// - Where did you get stuck or hit gotchas?
// - Why do we need multiple instances of the mapper and reducer?
// - What affect does changing the number of instances have?
// - Is this approach simpler than the starting point?
// - Does this code run faster than the starting point?
// - Based on the answers to the above, why would we consider an
//   approach like this?
//
// ─────────────────
// Further exploration
// ─────────────────
//
// If you'd like to explore further you may attempt the following
// variations on your own, or arrange further mentoring sessions to
// explore them:
//
// - try using buffered channels.
//   - does this impact the programs performance?
// - create a version of the program that counts occurrences of each word.
//   - what new problems does this introduce?

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

// countWords iterates over the content of io.Reader breaking it down
// into lines and summing the lengths of those lines as its output.
// This is not necessarily the idiomatic way to achieve this goal in
// Go, but rather it is intended to give you the building blocks
// needed to break this imperative procedure into a map-reduce and
// aggregation pattern.
func countWords(input io.Reader) int {
	var count int

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		count += len(strings.Fields(scanner.Text()))
	}
	return count
}

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Please supply the path to exactly 1 plain text file.")
	}

	path := os.Args[1]

	input, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := input.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	count := countWords(input)
	fmt.Printf("\t%d %s\n", count, filepath.Base(path))
}

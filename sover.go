package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"

	"github.com/coreos/go-semver/semver"
	"github.com/mattn/go-isatty"
)

var stdin io.Reader
var stdout io.Writer

var reverse bool

func init() {
	if isatty.IsTerminal(os.Stdin.Fd()) || isatty.IsCygwinTerminal(os.Stdin.Fd()) {
		stdin = bytes.NewBufferString("")
	} else {
		stdin = os.Stdin
	}
	stdout = os.Stdout

	flag.BoolVar(&reverse, "r", false, "sort in reverse order")
}

func main() {
	flag.Parse()
	args := strings.Join(flag.Args(), "\n")

	versions := make([]*version, 0)
	sc := bufio.NewScanner(io.MultiReader(
		stdin,
		bytes.NewBufferString("\n"),
		bytes.NewBufferString(args),
	))

	for sc.Scan() {
		t := sc.Text()
		if len(t) == 0 {
			continue
		}
		semv := semver.New(strings.TrimPrefix(t, "v"))
		versions = append(versions, &version{semv: *semv, raw: t})
	}

	sort.Slice(versions, func(i, j int) bool {
		return versions[i].semv.Compare(versions[j].semv) < 0
	})

	for i := range versions {
		if reverse {
			i = len(versions) - i - 1
		}
		fmt.Fprintln(stdout, versions[i].raw)
	}
}

type version struct {
	semv semver.Version
	raw  string
}

package path_matcher

import (
	"github.com/docker/docker/pkg/fileutils"
	"path/filepath"

	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func newPatternMatcher(patterns []string) *fileutils.PatternMatcher {
	m, err := fileutils.NewPatternMatcher(patterns)
	if err != nil {
		panic(err)
	}

	return m
}

type dockerfileIgnorePathMatchEntry struct {
	baseBase        string
	patternMatcher  *fileutils.PatternMatcher
	matchedPaths    []string
	notMatchedPaths []string
}

var _ = DescribeTable("DockerfileIgnore_MatchPath", func(e dockerfileIgnorePathMatchEntry) {
	pathMatcher := NewDockerfileIgnorePathMatcher(e.baseBase, e.patternMatcher)

	for _, matchedPath := range e.matchedPaths {
		Ω(pathMatcher.MatchPath(matchedPath)).Should(BeTrue())
	}

	for _, notMatchedPath := range e.notMatchedPaths {
		Ω(pathMatcher.MatchPath(notMatchedPath)).Should(BeFalse())
	}
},
	Entry("basePath is equal to the path (base)", dockerfileIgnorePathMatchEntry{
		baseBase:     filepath.Join("a", "b", "c"),
		matchedPaths: []string{filepath.Join("a", "b", "c")},
	}),

	Entry("path is relative to the basePath (base)", dockerfileIgnorePathMatchEntry{
		baseBase:     filepath.Join("a", "b", "c"),
		matchedPaths: []string{filepath.Join("a", "b", "c", "d")},
	}),
	Entry("path is relative to the basePath (exclude record)", dockerfileIgnorePathMatchEntry{
		baseBase:        filepath.Join("a", "b", "c"),
		patternMatcher:  newPatternMatcher([]string{"d"}),
		matchedPaths:    []string{filepath.Join("a", "b", "c", "de")},
		notMatchedPaths: []string{filepath.Join("a", "b", "c", "d"), filepath.Join("a", "b", "c", "d", "e")},
	}),
	Entry("path is relative to the basePath (exclude with exclusion)", dockerfileIgnorePathMatchEntry{
		baseBase:        filepath.Join("a", "b", "c"),
		patternMatcher:  newPatternMatcher([]string{"d", "!d/e"}),
		matchedPaths:    []string{filepath.Join("a", "b", "c", "d", "e")},
		notMatchedPaths: []string{filepath.Join("a", "b", "c", "d")},
	}),

	Entry("path is not relative to the basePath (base)", dockerfileIgnorePathMatchEntry{
		baseBase:        filepath.Join("a", "b", "c"),
		notMatchedPaths: []string{filepath.Join("a", "b", "d"), "b"},
	}),
	Entry("path is not relative to the basePath(exclude record)", dockerfileIgnorePathMatchEntry{
		baseBase:        filepath.Join("a", "b", "c"),
		patternMatcher:  newPatternMatcher([]string{"d"}),
		notMatchedPaths: []string{filepath.Join("a", "b", "d"), "b"},
	}),
	Entry("path is not relative to the basePath (exclude record with exclusion)", dockerfileIgnorePathMatchEntry{
		baseBase:        filepath.Join("a", "b", "c"),
		patternMatcher:  newPatternMatcher([]string{"d", "!d/e"}),
		notMatchedPaths: []string{filepath.Join("a", "b", "d"), "b"},
	}),

	//Entry("basePath is relative to the path (base)", dockerfileIgnorePathMatchEntry{
	//	baseBase:        filepath.Join("a", "b", "c"),
	//	notMatchedPaths: []string{filepath.Join("a")},
	//}),
	//Entry("basePath is relative to the path (includePaths)", dockerfileIgnorePathMatchEntry{
	//	baseBase:        filepath.Join("a", "b", "c"),
	//	includePaths:    []string{"d"},
	//	notMatchedPaths: []string{filepath.Join("a")},
	//}),
	//Entry("basePath is relative to the path (excludePaths)", dockerfileIgnorePathMatchEntry{
	//	baseBase:        filepath.Join("a", "b", "c"),
	//	excludePaths:    []string{"d"},
	//	notMatchedPaths: []string{filepath.Join("a")},
	//}),
	//Entry("basePath is relative to the path (includePaths and excludePaths)", dockerfileIgnorePathMatchEntry{
	//	baseBase:        filepath.Join("a", "b", "c"),
	//	includePaths:    []string{"d"},
	//	excludePaths:    []string{"e"},
	//	notMatchedPaths: []string{filepath.Join("a")},
	//}),
	//
	//Entry("glob completion by default (includePaths)", dockerfileIgnorePathMatchEntry{
	//	includePaths: []string{
	//		"a",
	//		filepath.Join("b", "*"),
	//		filepath.Join("c", "**"),
	//		filepath.Join("d", "**", "*"),
	//	},
	//	matchedPaths: []string{
	//		filepath.Join("a", "b", "c", "d"),
	//		filepath.Join("b", "b", "c", "d"),
	//		filepath.Join("c", "b", "c", "d"),
	//		filepath.Join("d", "b", "c", "d"),
	//	},
	//}),
	//Entry("glob completion by default (excludePaths)", dockerfileIgnorePathMatchEntry{
	//	excludePaths: []string{
	//		"a",
	//		filepath.Join("b", "*"),
	//		filepath.Join("c", "**"),
	//		filepath.Join("d", "**", "*"),
	//	},
	//	notMatchedPaths: []string{
	//		filepath.Join("a", "b", "c", "d"),
	//		filepath.Join("b", "b", "c", "d"),
	//		filepath.Join("c", "b", "c", "d"),
	//		filepath.Join("d", "b", "c", "d"),
	//	},
	//}),
	//Entry("glob completion by default (includePaths and excludePaths)", dockerfileIgnorePathMatchEntry{
	//	includePaths: []string{
	//		"a",
	//		filepath.Join("b", "*"),
	//		filepath.Join("c", "**"),
	//		filepath.Join("d", "**", "*"),
	//	},
	//	excludePaths: []string{
	//		"a",
	//		filepath.Join("b", "*"),
	//		filepath.Join("c", "**"),
	//		filepath.Join("d", "**", "*"),
	//	},
	//	notMatchedPaths: []string{
	//		filepath.Join("a", "b", "c", "d"),
	//		filepath.Join("b", "b", "c", "d"),
	//		filepath.Join("c", "b", "c", "d"),
	//		filepath.Join("d", "b", "c", "d"),
	//	},
	//}),
)

type dockerfileIgnoreProcessDirOrSubmodulePath struct {
	baseBase               string
	patternMatcher         *fileutils.PatternMatcher
	matchedPaths           []string
	shouldWalkThroughPaths []string
	notMatchedPaths        []string
}

var _ = DescribeTable("DockerfileIgnore_ProcessDirOrSubmodulePath", func(e dockerfileIgnoreProcessDirOrSubmodulePath) {
	pathMatcher := NewDockerfileIgnorePathMatcher(e.baseBase, e.patternMatcher)

	for _, matchedPath := range e.matchedPaths {
		isMatched, shouldWalkThrough := pathMatcher.ProcessDirOrSubmodulePath(matchedPath)
		Ω(isMatched).Should(BeTrue())
		Ω(shouldWalkThrough).Should(BeFalse())
	}

	for _, shouldWalkThroughPath := range e.shouldWalkThroughPaths {
		isMatched, shouldWalkThrough := pathMatcher.ProcessDirOrSubmodulePath(shouldWalkThroughPath)
		Ω(isMatched).Should(BeFalse())
		Ω(shouldWalkThrough).Should(BeTrue())
	}

	for _, notMatchedPath := range e.notMatchedPaths {
		isMatched, shouldWalkThrough := pathMatcher.ProcessDirOrSubmodulePath(notMatchedPath)
		Ω(isMatched).Should(BeFalse())
		Ω(shouldWalkThrough).Should(BeFalse())
	}
},
	Entry("basePath is equal to the path (base)", dockerfileIgnoreProcessDirOrSubmodulePath{
		baseBase:     filepath.Join("a", "b", "c"),
		matchedPaths: []string{filepath.Join("a", "b", "c")},
	}),

	Entry("path is relative to the basePath (base)", dockerfileIgnoreProcessDirOrSubmodulePath{
		baseBase:     filepath.Join("a", "b", "c"),
		matchedPaths: []string{filepath.Join("a", "b", "c", "d")},
	}),
	Entry("path is relative to the basePath (exclude record)", dockerfileIgnoreProcessDirOrSubmodulePath{
		baseBase:        filepath.Join("a", "b", "c"),
		patternMatcher:  newPatternMatcher([]string{"d"}),
		matchedPaths:    []string{filepath.Join("a", "b", "c", "de")},
		notMatchedPaths: []string{filepath.Join("a", "b", "c", "d"), filepath.Join("a", "b", "c", "d", "e")},
	}),
	Entry("path is relative to the basePath (exclude with exclusion)", dockerfileIgnoreProcessDirOrSubmodulePath{
		baseBase:               filepath.Join("a", "b", "c"),
		patternMatcher:         newPatternMatcher([]string{"d", "!d/e"}),
		matchedPaths:           []string{filepath.Join("a", "b", "c", "d", "e")},
		shouldWalkThroughPaths: []string{filepath.Join("a", "b", "c", "d")},
	}),

//	Entry("basePath is equal to the path (includePaths)", dockerfileIgnoreProcessDirOrSubmodulePath{
//		baseBase:               filepath.Join("a", "b", "c"),
//		includePaths:           []string{"d"},
//		shouldWalkThroughPaths: []string{filepath.Join("a", "b", "c")},
//	}),
//	Entry("basePath is equal to the path (excludePaths)", dockerfileIgnoreProcessDirOrSubmodulePath{
//		baseBase:               filepath.Join("a", "b", "c"),
//		excludePaths:           []string{"d"},
//		shouldWalkThroughPaths: []string{filepath.Join("a", "b", "c")},
//	}),
//	Entry("basePath is equal to the path (includePaths and excludePaths)", dockerfileIgnoreProcessDirOrSubmodulePath{
//		baseBase:               filepath.Join("a", "b", "c"),
//		includePaths:           []string{"d"},
//		excludePaths:           []string{"e"},
//		shouldWalkThroughPaths: []string{filepath.Join("a", "b", "c")},
//	}),
//	Entry("basePath is equal to the path, includePath ''", dockerfileIgnoreProcessDirOrSubmodulePath{
//		baseBase:     filepath.Join("a", "b", "c"),
//		includePaths: []string{""},
//		matchedPaths: []string{filepath.Join("a", "b", "c")},
//	}),
//	Entry("basePath is equal to the path, excludePath ''", dockerfileIgnoreProcessDirOrSubmodulePath{
//		baseBase:        filepath.Join("a", "b", "c"),
//		excludePaths:    []string{""},
//		notMatchedPaths: []string{filepath.Join("a", "b", "c")},
//	}),
//	Entry("basePath is equal to the path, includePath '', excludePath ''", dockerfileIgnoreProcessDirOrSubmodulePath{
//		baseBase:        filepath.Join("a", "b", "c"),
//		includePaths:    []string{""},
//		excludePaths:    []string{""},
//		notMatchedPaths: []string{filepath.Join("a", "b", "c")},
//	}),
//
//	Entry("path is relative to the basePath (base)", dockerfileIgnoreProcessDirOrSubmodulePath{
//		baseBase:     filepath.Join("a", "b", "c"),
//		matchedPaths: []string{filepath.Join("a", "b", "c", "d")},
//	}),
//	Entry("path is relative to the basePath (includePaths)", dockerfileIgnoreProcessDirOrSubmodulePath{
//		baseBase:        filepath.Join("a", "b", "c"),
//		includePaths:    []string{"d"},
//		matchedPaths:    []string{filepath.Join("a", "b", "c", "d")},
//		notMatchedPaths: []string{filepath.Join("a", "b", "c", "e"), filepath.Join("a", "b", "c", "de")},
//	}),
//	Entry("path is relative to the basePath (excludePaths)", dockerfileIgnoreProcessDirOrSubmodulePath{
//		baseBase:        filepath.Join("a", "b", "c"),
//		excludePaths:    []string{"d"},
//		matchedPaths:    []string{filepath.Join("a", "b", "c", "e"), filepath.Join("a", "b", "c", "de")},
//		notMatchedPaths: []string{filepath.Join("a", "b", "c", "d")},
//	}),
//	Entry("path is relative to the basePath (includePaths and excludePaths)", dockerfileIgnoreProcessDirOrSubmodulePath{
//		baseBase:        filepath.Join("a", "b", "c"),
//		includePaths:    []string{"d"},
//		excludePaths:    []string{"e"},
//		matchedPaths:    []string{filepath.Join("a", "b", "c", "d")},
//		notMatchedPaths: []string{filepath.Join("a", "b", "c", "e")},
//	}),
//
//	Entry("path is not relative to the basePath (base)", dockerfileIgnoreProcessDirOrSubmodulePath{
//		baseBase:        filepath.Join("a", "b", "c"),
//		notMatchedPaths: []string{filepath.Join("a", "b", "d"), "b"},
//	}),
//	Entry("path is not relative to the basePath(includePaths)", dockerfileIgnoreProcessDirOrSubmodulePath{
//		baseBase:        filepath.Join("a", "b", "c"),
//		includePaths:    []string{"d"},
//		notMatchedPaths: []string{filepath.Join("a", "b", "d"), "b"},
//	}),
//	Entry("path is not relative to the basePath (excludePaths)", dockerfileIgnoreProcessDirOrSubmodulePath{
//		baseBase:        filepath.Join("a", "b", "c"),
//		excludePaths:    []string{"d"},
//		notMatchedPaths: []string{filepath.Join("a", "b", "d"), "b"},
//	}),
//	Entry("path is not relative to the basePath (includePaths and excludePaths)", dockerfileIgnoreProcessDirOrSubmodulePath{
//		baseBase:        filepath.Join("a", "b", "c"),
//		includePaths:    []string{"d"},
//		excludePaths:    []string{"e"},
//		notMatchedPaths: []string{filepath.Join("a", "b", "d"), "b"},
//	}),
//
//	Entry("basePath is relative to the path (base)", dockerfileIgnoreProcessDirOrSubmodulePath{
//		baseBase:               filepath.Join("a", "b", "c"),
//		shouldWalkThroughPaths: []string{filepath.Join("a")},
//	}),
//	Entry("basePath is relative to the path (includePaths)", dockerfileIgnoreProcessDirOrSubmodulePath{
//		baseBase:               filepath.Join("a", "b", "c"),
//		includePaths:           []string{"d"},
//		shouldWalkThroughPaths: []string{filepath.Join("a")},
//	}),
//	Entry("basePath is relative to the path (excludePaths)", dockerfileIgnoreProcessDirOrSubmodulePath{
//		baseBase:               filepath.Join("a", "b", "c"),
//		excludePaths:           []string{"d"},
//		shouldWalkThroughPaths: []string{filepath.Join("a")},
//	}),
//	Entry("basePath is relative to the path (includePaths and excludePaths)", dockerfileIgnoreProcessDirOrSubmodulePath{
//		baseBase:               filepath.Join("a", "b", "c"),
//		includePaths:           []string{"d"},
//		excludePaths:           []string{"e"},
//		shouldWalkThroughPaths: []string{filepath.Join("a")},
//	}),
//
//	Entry("glob completion by default (includePaths)", dockerfileIgnoreProcessDirOrSubmodulePath{
//		includePaths: []string{
//			"a",
//			filepath.Join("b", "*"),
//			filepath.Join("c", "**"),
//			filepath.Join("d", "**", "*"),
//		},
//		matchedPaths: []string{
//			filepath.Join("a", "b", "c", "d"),
//			filepath.Join("b", "b", "c", "d"),
//			filepath.Join("c", "b", "c", "d"),
//			filepath.Join("d", "b", "c", "d"),
//		},
//	}),
//	Entry("glob completion by default (excludePaths)", dockerfileIgnoreProcessDirOrSubmodulePath{
//		excludePaths: []string{
//			"a",
//			filepath.Join("b", "*"),
//			filepath.Join("c", "**"),
//			filepath.Join("d", "**", "*"),
//		},
//		notMatchedPaths: []string{
//			filepath.Join("a", "b", "c", "d"),
//			filepath.Join("b", "b", "c", "d"),
//			filepath.Join("c", "b", "c", "d"),
//			filepath.Join("d", "b", "c", "d"),
//		},
//	}),
//	Entry("glob completion by default (includePaths and excludePaths)", dockerfileIgnoreProcessDirOrSubmodulePath{
//		includePaths: []string{
//			"a",
//			filepath.Join("b", "*"),
//			filepath.Join("c", "**"),
//			filepath.Join("d", "**", "*"),
//		},
//		excludePaths: []string{
//			"a",
//			filepath.Join("b", "*"),
//			filepath.Join("c", "**"),
//			filepath.Join("d", "**", "*"),
//		},
//		notMatchedPaths: []string{
//			filepath.Join("a", "b", "c", "d"),
//			filepath.Join("b", "b", "c", "d"),
//			filepath.Join("c", "b", "c", "d"),
//			filepath.Join("d", "b", "c", "d"),
//		},
//	}),
)

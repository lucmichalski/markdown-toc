package toc

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuild(t *testing.T) {
	testCases := map[string]struct {
		data        []byte
		header      string
		addHeader   bool
		expectedToC []string
	}{
		"success - empty markdown": {
			data:      []byte(``),
			header:    "# Table of Contents",
			addHeader: true,
			expectedToC: []string{
				"<!-- ToC start -->",
				"# Table of Contents\n",
				"<!-- ToC end -->\n",
			},
		},
		"success - full example": {
			data: []byte(`
Header 1
========

Some content

Header 2
--------

Some content

# Header 3

Some content

## Header 4

Some content

### Header 5

Some content

#### Header 6

Some content
`),
			header:    "# Table of Contents",
			addHeader: true,
			expectedToC: []string{
				"<!-- ToC start -->",
				"# Table of Contents\n",
				"- [Header 1](#header-1)",
				"  - [Header 2](#header-2)",
				"- [Header 3](#header-3)",
				"  - [Header 4](#header-4)",
				"    - [Header 5](#header-5)",
				"      - [Header 6](#header-6)",
				"<!-- ToC end -->\n",
			},
		},
	}

	for name, testCase := range testCases {
		t.Run(name, func(t *testing.T) {
			toc, err := Build(testCase.data, testCase.header, testCase.addHeader)
			assert.Nil(t, err)
			assert.Equal(t, testCase.expectedToC, toc)
		})
	}
}
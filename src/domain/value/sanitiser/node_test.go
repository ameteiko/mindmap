package sanitiser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitiseSlug(t *testing.T) {
	tcs := []struct {
		desc          string
		slug          string
		sanitisedSlug string
	}{
		{
			desc:          "ForACleanSlug",
			slug:          "shell",
			sanitisedSlug: "shell",
		},
		{
			desc:          "ForAnEmptySlug",
			slug:          "",
			sanitisedSlug: "",
		},
		{
			desc:          "ForAnEmptyWhitespacesSlug",
			slug:          "  ",
			sanitisedSlug: "",
		},
		{
			desc:          "ForASlugWrappedWithWhitespaces",
			slug:          " shell  ",
			sanitisedSlug: "shell",
		},
		{
			desc:          "ForASlugWithCapitalLetters",
			slug:          "SheLL",
			sanitisedSlug: "shell",
		},

		{
			desc:          "ForASlugWithSeveralWords",
			slug:          " SheLL  ScripTing env    variables   ",
			sanitisedSlug: "shell-scripting-env-variables",
		},
	}

	for i := range tcs {
		tc := tcs[i]
		t.Run(tc.desc, func(t *testing.T) {
			nodeSanitiser := Node{}

			sanitisedSlug := nodeSanitiser.sanitiseSlug(tc.slug)

			assert.Equal(t, tc.sanitisedSlug, sanitisedSlug, "Sanitised slug doesn't match.")
		})
	}
}
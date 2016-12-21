package filter

import "strings"

// Dev is the official politically correct filter
func Dev(in string) string {

	r := strings.NewReplacer(
		"it's very verbose", "it's Java",
		"make it more performant", "use Go",
		"refactor the code", "add 10.000 sprints",
		"your code is fine", "your code is shit",
		"this code is shit", "this code was written by another dev and I don't understand it",
		"dan", "TROLOLLOLLLOLLLOOOLOLO",
	)

	// Replace all pairs.
	return r.Replace(in)
}

// Division translator is the official D&T filter: useful for starters.
func Division(in string) string {
	r := strings.NewReplacer(
		"infosec", "the Dark Side of the Force",
		"don't worry", "be worried",
	)

	// Replace all pairs.
	return r.Replace(in)

}

// UK is...
func UK(in string) string {
	r := strings.NewReplacer(
		"i hear what you say", "i disagree and do not want to discuss it further",
		"with the greatest respect", "you are an idiot",
		"that's not bad", "that's good",
		"that is a very brave proposal", "you are insane",
		"quite good", "a bit disappointing",
		"i would suggest", "do it or prepare to justify yourself",
		"oh, incidentally", "the primary purpose of our discussion is",
		"oh, by the way", "the primary purpose of our discussion is",
		"i was a bit disappointed that", "i was annoyed that",
		"very interesting", "this is clearly nonsense",
		"i'll bear that in mind", "i've forgotten it already",
		"i'm sure it's my fault", "it's your fault",
		"you must come for dinner", "it's not an invitation, i'm just being polite",
		"i almost agree", "i don't agree at all",
		"i only have a few minor comments", "please rewrite it completely",
		"could we consider some other option", "i don't like your idea",
	)

	// Replace all pairs.
	return r.Replace(in)

}

// Apply applies all the filters
func Apply(s *string, filters ...func(in string) string) {
	newS := strings.ToLower(*s)
	for _, f := range filters {
		newS = f(newS)
	}
	*s = newS
}

package hello

// ShareWith return a greeting string.
func ShareWith(who string) string {
	if len(who) == 0 {
		who = "you"
	}
	return "One for " + who + ", one for me."
}

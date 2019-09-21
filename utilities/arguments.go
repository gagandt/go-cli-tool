package utilities

func FindIndexOfCommandInOsArgs(args []string, argToFind string) int {
	for index, arg := range(args) {
		if arg == argToFind {
			return index
		}
	}
	return -1
}

func OsArgsHaveSufficientNumberOfArgs(lengthOfOsArgs, currentIndex, NumberOfElementsNeeded int) bool {
	if currentIndex + NumberOfElementsNeeded < lengthOfOsArgs {
		return true
	}
	return false
}
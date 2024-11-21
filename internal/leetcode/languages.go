package leetcode

var (
	AvailableLanguage = map[string]string{
		"cpp":  "C++",
		"java": "Java",
		"py":   "Python",
		"ts":   "TypeScript",
		"sql":  "MySQL",
		"sh":   "Shell",
	}
)

func ValidLanguage(language string) bool {
	_, exists := AvailableLanguage[language]
	return exists
}

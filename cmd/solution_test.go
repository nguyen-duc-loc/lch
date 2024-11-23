package cmd

import (
	"bytes"
	"math"
	"testing"

	"github.com/nguyen-duc-loc/lch/internal/leetcode"
	"github.com/stretchr/testify/require"
)

func TestSolutionActions(t *testing.T) {
	testCases := []struct {
		name      string
		problemID int32
		lang      string
		expectErr error
		expectOut string
	}{
		{
			name:      "ProblemNotFound",
			problemID: math.MaxInt32,
			lang:      "cpp",
			expectErr: leetcode.ErrProblemNotFound,
		},
		{
			name:      "UnsupportedLanguage",
			problemID: 1,
			lang:      "unsupportedLanguage",
			expectErr: leetcode.ErrLanguageNotSupported,
		},
		{
			name:      "SolutionNotWrittenInCpp",
			problemID: 175,
			lang:      "cpp",
			expectErr: leetcode.ErrUnableToGetSolution(175, "cpp"),
		},
		{
			name:      "SolutionNotWrittenInJava",
			problemID: 175,
			lang:      "java",
			expectErr: leetcode.ErrUnableToGetSolution(175, "java"),
		},
		{
			name:      "SolutionNotWrittenInPython",
			problemID: 175,
			lang:      "py",
			expectErr: leetcode.ErrUnableToGetSolution(175, "py"),
		},
		{
			name:      "SolutionNotWrittenInTypescript",
			problemID: 175,
			lang:      "ts",
			expectErr: leetcode.ErrUnableToGetSolution(175, "ts"),
		},
		{
			name:      "SolutionNotWrittenInMySQL",
			problemID: 1,
			lang:      "sql",
			expectErr: leetcode.ErrUnableToGetSolution(1, "sql"),
		},
		{
			name:      "SolutionNotWrittenInShell",
			problemID: 1,
			lang:      "sh",
			expectErr: leetcode.ErrUnableToGetSolution(1, "sh"),
		},
		{
			name:      "CppSolutionOK",
			problemID: 1,
			lang:      "cpp",
			expectOut: `class Solution {
 public:
  vector<int> twoSum(vector<int>& nums, int target) {
    unordered_map<int, int> numToIndex;

    for (int i = 0; i < nums.size(); ++i) {
      if (const auto it = numToIndex.find(target - nums[i]);
          it != numToIndex.cend())
        return {it->second, i};
      numToIndex[nums[i]] = i;
    }

    throw;
  }
};
`,
		},
		{
			name:      "JavaSolutionOK",
			problemID: 1,
			lang:      "java",
			expectOut: `class Solution {
  public int[] twoSum(int[] nums, int target) {
    Map<Integer, Integer> numToIndex = new HashMap<>();

    for (int i = 0; i < nums.length; ++i) {
      if (numToIndex.containsKey(target - nums[i]))
        return new int[] {numToIndex.get(target - nums[i]), i};
      numToIndex.put(nums[i], i);
    }

    throw new IllegalArgumentException();
  }
}
`,
		},
		{
			name:      "PythonSolutionOK",
			problemID: 1,
			lang:      "py",
			expectOut: `class Solution:
  def twoSum(self, nums: list[int], target: int) -> list[int]:
    numToIndex = {}

    for i, num in enumerate(nums):
      if target - num in numToIndex:
        return numToIndex[target - num], i
      numToIndex[num] = i
`,
		},
		{
			name:      "TSSolutionOK",
			problemID: 2618,
			lang:      "ts",
			expectOut: `function checkIfInstanceOf(obj: any, classFunction: any): boolean {
  while (obj != null) {
    if (obj.constructor === classFunction) {
      return true;
    }
    obj = Object.getPrototypeOf(obj);
  }
  return false;
}
`,
		},
		{
			name:      "MySQLSolutionOK",
			problemID: 175,
			lang:      "sql",
			expectOut: `SELECT
  Person.firstName,
  Person.lastName,
  Address.city,
  Address.state
FROM Person
LEFT JOIN Address
  USING (personId);
`,
		},
		{
			name:      "ShellSolutionOK",
			problemID: 192,
			lang:      "sh",
			expectOut: `cat words.txt | tr -s ' ' '\n' | sort | uniq -c | sort -r | awk '{ print $2, $1 }'
`,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var out bytes.Buffer

			if err := solutionActions(&out, tc.problemID, tc.lang); err != nil {
				require.Error(t, tc.expectErr)
				require.Empty(t, tc.expectOut)
				require.EqualError(t, err, tc.expectErr.Error())
			} else {
				require.Equal(t, out.String(), tc.expectOut)
			}
		})
	}
}

// Code generated by "stringer -type SandboxTestResult .\sandbox\sandbox.go"; DO NOT EDIT.

package sandbox

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NoTest-0]
	_ = x[TestFailed-1]
	_ = x[TestPassed-2]
}

const _SandboxTestResult_name = "NoTestTestFailedTestPassed"

var _SandboxTestResult_index = [...]uint8{0, 6, 16, 26}

func (i SandboxTestResult) String() string {
	if i < 0 || i >= SandboxTestResult(len(_SandboxTestResult_index)-1) {
		return "SandboxTestResult(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _SandboxTestResult_name[_SandboxTestResult_index[i]:_SandboxTestResult_index[i+1]]
}

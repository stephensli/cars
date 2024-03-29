// Code generated by "stringer -type=ContainerStatus"; DO NOT EDIT.

package sandbox

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NotRan-0]
	_ = x[Created-1]
	_ = x[Running-2]
	_ = x[Killing-3]
	_ = x[Killed-4]
	_ = x[Finished-5]
	_ = x[MemoryConstraintExceeded-6]
	_ = x[TimeLimitExceeded-7]
	_ = x[ProvidedTestFailed-8]
	_ = x[CompilationFailed-9]
	_ = x[RunTimeError-10]
	_ = x[NonDeterministicError-11]
}

const _ContainerStatus_name = "NotRanCreatedRunningKillingKilledFinishedMemoryConstraintExceededTimeLimitExceededProvidedTestFailedCompilationFailedRunTimeErrorNonDeterministicError"

var _ContainerStatus_index = [...]uint8{0, 6, 13, 20, 27, 33, 41, 65, 82, 100, 117, 129, 150}

func (i ContainerStatus) String() string {
	if i < 0 || i >= ContainerStatus(len(_ContainerStatus_index)-1) {
		return "ContainerStatus(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ContainerStatus_name[_ContainerStatus_index[i]:_ContainerStatus_index[i+1]]
}

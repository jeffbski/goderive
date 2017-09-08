// Code generated by goderive DO NOT EDIT.

package sort

import (
	"sort"
	"strings"
)

// deriveCompare returns:
//   * 0 if this and that are equal,
//   * -1 is this is smaller and
//   * +1 is this is bigger.
func deriveCompare(this, that []*MyStruct) int {
	if this == nil {
		if that == nil {
			return 0
		}
		return -1
	}
	if that == nil {
		return 1
	}
	if len(this) != len(that) {
		if len(this) < len(that) {
			return -1
		}
		return 1
	}
	for i := 0; i < len(this); i++ {
		if c := deriveCompare_(this[i], that[i]); c != 0 {
			return c
		}
	}
	return 0
}

// deriveSort sorts the slice inplace and also returns it.
func deriveSort(list []*MyStruct) []*MyStruct {
	sort.Slice(list, func(i, j int) bool { return deriveCompare_(list[i], list[j]) < 0 })
	return list
}

// deriveCompare_ returns:
//   * 0 if this and that are equal,
//   * -1 is this is smaller and
//   * +1 is this is bigger.
func deriveCompare_(this, that *MyStruct) int {
	if this == nil {
		if that == nil {
			return 0
		}
		return -1
	}
	if that == nil {
		return 1
	}
	if c := deriveCompare_i(this.Int64, that.Int64); c != 0 {
		return c
	}
	if c := deriveCompare_1(this.StringPtr, that.StringPtr); c != 0 {
		return c
	}
	return 0
}

// deriveCompare_i returns:
//   * 0 if this and that are equal,
//   * -1 is this is smaller and
//   * +1 is this is bigger.
func deriveCompare_i(this, that int64) int {
	if this != that {
		if this < that {
			return -1
		} else {
			return 1
		}
	}
	return 0
}

// deriveCompare_1 returns:
//   * 0 if this and that are equal,
//   * -1 is this is smaller and
//   * +1 is this is bigger.
func deriveCompare_1(this, that *string) int {
	if this == nil {
		if that == nil {
			return 0
		}
		return -1
	}
	if that == nil {
		return 1
	}
	return deriveCompare_s(*this, *that)
}

// deriveCompare_s returns:
//   * 0 if this and that are equal,
//   * -1 is this is smaller and
//   * +1 is this is bigger.
func deriveCompare_s(this, that string) int {
	return strings.Compare(this, that)
}

package kit

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type sliceInt64Value []int64

func newSliceInt64Value(val []int64, p *[]int64) *sliceInt64Value {
	*p = val
	return (*sliceInt64Value)(p)
}

func (s *sliceInt64Value) Set(val string) error {
	split := strings.Split(val, ",")
	res := make([]int64, len(split))
	for i := range split {
		parse, err := strconv.ParseInt(split[i], 10, 64)
		if err != nil {
			return err
		}
		res[i] = parse
	}
	*s = res
	return nil
}

func (s *sliceInt64Value) Get() any { return []int64(*s) }

func (s *sliceInt64Value) String() string { return fmt.Sprintf("%v", *s) }

func SliceInt64Var(p *[]int64, name string, value []int64, usage string) {
	flag.CommandLine.Var(newSliceInt64Value(value, p), name, usage)
}

type sliceStrValue []string

func newSliceStrValue(val []string, p *[]string) *sliceStrValue {
	*p = val
	return (*sliceStrValue)(p)
}

func (s *sliceStrValue) Set(val string) error {
	split := strings.Split(val, ",")
	res := make([]string, len(split))
	for i := range split {
		res[i] = split[i]
	}
	*s = res
	return nil
}

func (s *sliceStrValue) Get() any { return []string(*s) }

func (s *sliceStrValue) String() string { return fmt.Sprintf("%v", *s) }

func SliceStrVar(p *[]string, name string, value []string, usage string) {
	flag.CommandLine.Var(newSliceStrValue(value, p), name, usage)
}

type iTime int64

func (s *iTime) Get() any { return int64(*s) }

func (s *iTime) String() string { return fmt.Sprintf("%v", *s) }

func (s *iTime) Set(val string) (err error) {
	sec, err := parseTime(val)
	if err != nil {
		return err
	}
	*s = iTime(sec)
	return nil
}

func newTime(val int64, p *int64) *iTime {
	*p = val
	return (*iTime)(p)
}

func Time(p *int64, name string, value int64, usage string) {
	flag.CommandLine.Var(newTime(value, p), name, usage)
}

type second int64

func (c *second) Get() any { return int64(*c) }

func (c *second) String() string { return fmt.Sprintf("%v", *c) }

func (c *second) Set(s string) (err error) {
	val, err := parseDuration(s)
	if err != nil {
		return err
	}
	*c = second(val.Seconds())
	return nil
}

func newSecond(val int64, p *int64) *second {
	*p = val
	return (*second)(p)
}

func Second(p *int64, name string, value int64, usage string) {
	flag.CommandLine.Var(newSecond(value, p), name, usage)
}

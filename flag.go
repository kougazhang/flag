package kit

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
	"time"
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

func SliceStr(p *[]string, name string, value []string, usage string) {
	flag.CommandLine.Var(newSliceStrValue(value, p), name, usage)
}

type iTime time.Time

func (s *iTime) Get() any { return time.Time(*s) }

func (s *iTime) String() string { return fmt.Sprintf("%v", *s) }

func (s *iTime) Set(val string) (err error) {
	tm, err := time.ParseInLocation("2006-01-02_15:04:05", val, time.Local)
	if err != nil {
		return err
	}
	*s = iTime(tm)
	return nil
}

func newTime(val time.Time, p *time.Time) *iTime {
	*p = val
	return (*iTime)(p)
}

func Time(p *time.Time, name string, value time.Time, usage string) {
	flag.CommandLine.Var(newTime(value, p), name, usage)
}

type duration time.Duration

func (s *duration) Get() any { return time.Duration(*s) }

func (s *duration) String() string { return fmt.Sprintf("%v", *s) }

func (s *duration) Set(val string) (err error) {
	tm, err := parseDuration(val)
	if err != nil {
		return err
	}
	*s = duration(tm)
	return nil
}

func newDuration(val time.Duration, p *time.Duration) *duration {
	*p = val
	return (*duration)(p)
}

func Duration(p *time.Duration, name string, value time.Duration, usage string) {
	flag.CommandLine.Var(newDuration(value, p), name, usage)
}

type date time.Time

func (s *date) Get() any { return time.Time(*s) }

func (s *date) String() string { return fmt.Sprintf("%v", *s) }

func (s *date) Set(val string) (err error) {
	tm, err := time.ParseInLocation("2006-01-02", val, time.Local)
	if err != nil {
		return err
	}
	*s = date(tm)
	return nil
}

func newDate(val time.Time, p *time.Time) *date {
	*p = val
	return (*date)(p)
}

func Date(p *time.Time, name string, value time.Time, usage string) {
	flag.CommandLine.Var(newDate(value, p), name, usage)
}

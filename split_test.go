package mafan

import (
	. "launchpad.net/gocheck"
	"strings"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type SplitTestSuite struct{}

var _ = Suite(&SplitTestSuite{})

func (s *SplitTestSuite) TestSplit(c *C) {
	c.Check(strings.Join(Split("上海十大接吻聖地"), " "), Equals, "上海 十大 接吻 聖地")
}

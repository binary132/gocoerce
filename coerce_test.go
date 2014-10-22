package coerce_test

import gc "gopkg.in/check.v1"

type coerceSuite struct{}

var _ = gc.Suite(certSuite{})

func (s *coerceSuite) TestInit(c *gc.C) {
	c.Assert(true, gc.Equals, true)
}

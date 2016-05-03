// Copyright 2015 ThoughtWorks, Inc.

// This file is part of Gauge.

// Gauge is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// Gauge is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with Gauge.  If not, see <http://www.gnu.org/licenses/>.

package event

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MySuite struct{}

var _ = Suite(&MySuite{})

func (s *MySuite) TestRegisterForOneTopic(c *C) {
	ch := make(chan ExecutionEvent)
	eventRegistry = nil

	Register(ch, StepEnd)

	c.Assert(len(eventRegistry), Equals, 1)
	c.Assert(eventRegistry[StepEnd][0], Equals, ch)
}

func (s *MySuite) TestRegisterForMultipleTopics(c *C) {
	ch := make(chan ExecutionEvent)
	eventRegistry = nil

	Register(ch, StepEnd, StepStart, SpecEnd, SpecStart)

	c.Assert(len(eventRegistry), Equals, 4)
	c.Assert(eventRegistry[StepEnd][0], Equals, ch)
	c.Assert(eventRegistry[StepStart][0], Equals, ch)
	c.Assert(eventRegistry[SpecEnd][0], Equals, ch)
	c.Assert(eventRegistry[SpecEnd][0], Equals, ch)
}

func (s *MySuite) TestMultipleSubscribersRegisteringForMultipleEvent(c *C) {
	eventRegistry = nil

	ch1 := make(chan ExecutionEvent)
	Register(ch1, StepStart, StepEnd)

	ch2 := make(chan ExecutionEvent)
	Register(ch2, StepStart, StepEnd)

	ch3 := make(chan ExecutionEvent)
	Register(ch3, SpecStart, SpecEnd, StepStart, StepEnd)

	c.Assert(len(eventRegistry[SpecStart]), Equals, 1)
	c.Assert(contains(eventRegistry[SpecStart], ch3), Equals, true)

	c.Assert(len(eventRegistry[SpecEnd]), Equals, 1)
	c.Assert(contains(eventRegistry[SpecEnd], ch3), Equals, true)

	c.Assert(len(eventRegistry[StepStart]), Equals, 3)
	c.Assert(contains(eventRegistry[StepStart], ch1), Equals, true)
	c.Assert(contains(eventRegistry[StepStart], ch2), Equals, true)
	c.Assert(contains(eventRegistry[StepStart], ch3), Equals, true)

	c.Assert(len(eventRegistry[StepEnd]), Equals, 3)
	c.Assert(contains(eventRegistry[StepEnd], ch1), Equals, true)
	c.Assert(contains(eventRegistry[StepEnd], ch2), Equals, true)
	c.Assert(contains(eventRegistry[StepEnd], ch3), Equals, true)
}

func contains(arr []chan ExecutionEvent, key chan ExecutionEvent) bool {
	for _, k := range arr {
		if k == key {
			return true
		}
	}
	return false
}

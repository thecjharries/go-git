package plumbing

import . "gopkg.in/check.v1"

type ReferenceSuite struct{}

var _ = Suite(&ReferenceSuite{})

var (
	ExampleReferenceName = NewReferenceName("refs/heads/v4")
)

func (s *ReferenceSuite) TestReferenceTypeString(c *C) {
	c.Assert(SymbolicReference.String(), Equals, "symbolic-reference")
}

func (s *ReferenceSuite) TestReferenceNameShort(c *C) {
	c.Assert(ExampleReferenceName.Short(), Equals, "v4")
}

func (s *ReferenceSuite) TestReferenceNameWithSlash(c *C) {
	r := NewReferenceName("refs/remotes/origin/feature/AllowSlashes")
	c.Assert(r.Short(), Equals, "origin/feature/AllowSlashes")
}

func (s *ReferenceSuite) TestReferenceNameNote(c *C) {
	r := NewReferenceName("refs/notes/foo")
	c.Assert(r.Short(), Equals, "notes/foo")
}

func (s *ReferenceSuite) TestNewReferenceFromStrings(c *C) {
	r := NewReferenceFromStrings("refs/heads/v4", "6ecf0ef2c2dffb796033e5a02219af86ec6584e5")
	c.Assert(r.Type(), Equals, HashReference)
	c.Assert(r.Name(), Equals, ExampleReferenceName)
	c.Assert(r.Hash(), Equals, NewHash("6ecf0ef2c2dffb796033e5a02219af86ec6584e5"))

	r = NewReferenceFromStrings("HEAD", "ref: refs/heads/v4")
	c.Assert(r.Type(), Equals, SymbolicReference)
	c.Assert(r.Name(), Equals, HEAD)
	c.Assert(r.Target(), Equals, ExampleReferenceName)
}

func (s *ReferenceSuite) TestNewSymbolicReference(c *C) {
	r := NewSymbolicReference(HEAD, ExampleReferenceName)
	c.Assert(r.Type(), Equals, SymbolicReference)
	c.Assert(r.Name(), Equals, HEAD)
	c.Assert(r.Target(), Equals, ExampleReferenceName)
}

func (s *ReferenceSuite) TestNewHashReference(c *C) {
	r := NewHashReference(ExampleReferenceName, NewHash("6ecf0ef2c2dffb796033e5a02219af86ec6584e5"))
	c.Assert(r.Type(), Equals, HashReference)
	c.Assert(r.Name(), Equals, ExampleReferenceName)
	c.Assert(r.Hash(), Equals, NewHash("6ecf0ef2c2dffb796033e5a02219af86ec6584e5"))
}

func (s *ReferenceSuite) TestNewBranchReferenceName(c *C) {
	r := NewBranchReferenceName("foo")
	c.Assert(r.String(), Equals, "refs/heads/foo")
}

func (s *ReferenceSuite) TestNewNoteReferenceName(c *C) {
	r := NewNoteReferenceName("foo")
	c.Assert(r.String(), Equals, "refs/notes/foo")
}

func (s *ReferenceSuite) TestNewRemoteReferenceName(c *C) {
	r := NewRemoteReferenceName("bar", "foo")
	c.Assert(r.String(), Equals, "refs/remotes/bar/foo")
}

func (s *ReferenceSuite) TestNewRemoteHEADReferenceName(c *C) {
	r := NewRemoteHEADReferenceName("foo")
	c.Assert(r.String(), Equals, "refs/remotes/foo/HEAD")
}

func (s *ReferenceSuite) TestNewTagReferenceName(c *C) {
	r := NewTagReferenceName("foo")
	c.Assert(r.String(), Equals, "refs/tags/foo")
}

func (s *ReferenceSuite) TestIsBranch(c *C) {
	r := ExampleReferenceName
	c.Assert(r.IsBranch(), Equals, true)
}

func (s *ReferenceSuite) TestIsNote(c *C) {
	r := NewReferenceName("refs/notes/foo")
	c.Assert(r.IsNote(), Equals, true)
}

func (s *ReferenceSuite) TestIsRemote(c *C) {
	r := NewReferenceName("refs/remotes/origin/master")
	c.Assert(r.IsRemote(), Equals, true)
}

func (s *ReferenceSuite) TestIsTag(c *C) {
	r := NewReferenceName("refs/tags/v3.1.")
	c.Assert(r.IsTag(), Equals, true)
}

var (
	RefNames = [][]string{
		[]string{
			"origin/master",
			"origin",
			"master",
		},
		[]string{
			"remotes/origin/master",
			"remotes",
			"origin",
			"master",
		},
		[]string{
			"refs/remotes/origin/master",
			"refs",
			"remotes",
			"origin",
			"master",
		},
	}
)

func (s *ReferenceSuite) TestReferenceNameFormat(c *C) {
	for _, ref_name_slice := range RefNames {
		new_name_method := NewReferenceName(ref_name_slice[1:]...)
		old_name_method := NewReferenceName(ref_name_slice[0])

		c.Assert(new_name_method.Name, Equals, old_name_method.Name)
		c.Assert(new_name_method.String(), Equals, old_name_method.String())
	}

}

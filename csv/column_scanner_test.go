package csv

import (
	"errors"
	"testing"

	"github.com/smarty/assertions/should"
	"github.com/smarty/gunit"
)

func TestColumnScannerFixture(t *testing.T) {
	gunit.Run(new(ColumnScannerFixture), t)
}

type ColumnScannerFixture struct {
	*gunit.Fixture

	scanner *ColumnScanner
	err     error
}

func (this *ColumnScannerFixture) Setup() {
	this.scanner, this.err = NewColumnScanner(NewScanner(reader(csvCanon)))
	this.So(this.err, should.BeNil)
	this.So(this.scanner.Header(), should.Resemble, []string{"first_name", "last_name", "username"})
}

func ScanAllUsers(scanner *ColumnScanner) []User {
	users := []User{}
	for scanner.Scan() {
		users = append(users, User{
			FirstName: scanner.Column(scanner.Header()[0]),
			LastName:  scanner.Column(scanner.Header()[1]),
			Username:  scanner.Column(scanner.Header()[2]),
		})
	}
	return users
}

func (this *ColumnScannerFixture) TestReadColumns() {
	users := ScanAllUsers(this.scanner)

	this.So(this.scanner.Error(), should.BeNil)
	this.So(users, should.Resemble, []User{
		{FirstName: "Rob", LastName: "Pike", Username: "rob"},
		{FirstName: "Ken", LastName: "Thompson", Username: "ken"},
		{FirstName: "Robert", LastName: "Griesemer", Username: "gri"},
	})
}

func (this *ColumnScannerFixture) TestCannotReadHeader() {
	scanner, err := NewColumnScanner(NewScanner(new(ErrorReader)))
	this.So(scanner, should.BeNil)
	this.So(err, should.NotBeNil)
}

func (this *ColumnScannerFixture) TestColumnNotFound_Error() {
	this.scanner.Scan()
	value, err := this.scanner.ColumnErr("nope")
	this.So(value, should.BeBlank)
	this.So(err, should.NotBeNil)
}

func (this *ColumnScannerFixture) TestColumnNotFound_Panic() {
	this.scanner.Scan()
	this.So(func() { this.scanner.Column("nope") }, should.Panic)
}

// TestDuplicateColumnNames confirms that duplicated/repeated
// column names results in the last repeated column being
// added to the map and used to retrieve values for that name.
func (this *ColumnScannerFixture) TestDuplicateColumnNames() {
	scanner, err := NewColumnScanner(NewScanner(reader([]string{
		"Col1,Col2,Col2",
		"foo,bar,baz",
	})))
	this.So(err, should.BeNil)
	this.So(scanner.Header(), should.Resemble, []string{"Col1", "Col2", "Col2"})
	scanner.Scan()
	this.So(scanner.Column("Col2"), should.Equal, "baz")
}

func (this *ColumnScannerFixture) TestColumnOpt_ToUpperHeader() {
	data := append([]string{"first_name,LAST_NAME,uSeRNaMe"}, csvCanon[1:]...)
	scanner, err := NewColumnScanner(
		NewScanner(reader(data)),
		ColumnOpts.ToUpperHeader())
	this.So(err, should.BeNil)

	users := ScanAllUsers(scanner)
	this.So(this.scanner.Error(), should.BeNil)
	this.So(users, should.Resemble, []User{
		{FirstName: "Rob", LastName: "Pike", Username: "rob"},
		{FirstName: "Ken", LastName: "Thompson", Username: "ken"},
		{FirstName: "Robert", LastName: "Griesemer", Username: "gri"},
	})
}

type User struct {
	FirstName string
	LastName  string
	Username  string
}

type ErrorReader struct{}

func (this *ErrorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("ERROR")
}

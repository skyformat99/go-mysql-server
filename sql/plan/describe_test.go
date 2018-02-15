package plan

import (
	"io"
	"testing"

	"github.com/src-d/go-mysql-server/mem"
	"github.com/src-d/go-mysql-server/sql"
	"github.com/stretchr/testify/assert"
)

func TestDescribe(t *testing.T) {
	assert := assert.New(t)

	table := mem.NewTable("test", sql.Schema{
		{Name: "c1", Type: sql.Text},
		{Name: "c2", Type: sql.Int32},
	})

	d := NewDescribe(table)
	iter, err := d.RowIter()
	assert.Nil(err)
	assert.NotNil(iter)

	n, err := iter.Next()
	assert.Nil(err)
	assert.Equal(sql.NewRow("c1", "TEXT"), n)

	n, err = iter.Next()
	assert.Nil(err)
	assert.Equal(sql.NewRow("c2", "INT32"), n)

	n, err = iter.Next()
	assert.Equal(io.EOF, err)
	assert.Nil(n)
}

func TestDescribe_Empty(t *testing.T) {
	assert := assert.New(t)

	d := NewDescribe(NewUnresolvedTable("test_table"))

	iter, err := d.RowIter()
	assert.Nil(err)
	assert.NotNil(iter)

	n, err := iter.Next()
	assert.Equal(io.EOF, err)
	assert.Nil(n)
}

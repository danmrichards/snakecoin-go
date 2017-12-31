package models

import (
	"testing"
	"time"
)

func TestBlock_String(t *testing.T) {
	tt := []struct {
		name           string
		block          *Block
		expectedString string
	}{
		{
			"test block 1",
			NewBlock(
				1,
				time.Date(2017, 12, 31, 15, 15, 24, 0, time.UTC),
				[]byte("test block 1"),
				[]byte("0"),
			),
			"12017-12-31T15:15:24Ztest block 10",
		},
		{
			"test block 2",
			NewBlock(
				2,
				time.Date(2017, 12, 31, 15, 15, 24, 0, time.UTC),
				[]byte("test block 2"),
				[]byte("1"),
			),
			"22017-12-31T15:15:24Ztest block 21",
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			if tc.block.String() != tc.expectedString {
				t.Errorf("TestBlock_String: %s, expected %s, got %s", tc.name, tc.expectedString, tc.block.String())
			}
		})
	}
}

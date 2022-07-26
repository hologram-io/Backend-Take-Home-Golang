package usage_parser

import (
	"net"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestUsageParser(t *testing.T) {
	t.Run("Parse", func(t *testing.T) {
		t.Run("Given a single string", func(t *testing.T) {
			t.Run("Given an ID that does not end with 4 or 6", func(t *testing.T) {
				var input []string = []string{
					"7291,293451",
				}
				var expected []*UsageResult = []*UsageResult{
					NewBasicUsage(7291, 293451),
				}

				t.Run("Then it will return basic string data", func(t *testing.T) {
					result := NewParser().Parse(input)

					if diff := cmp.Diff(expected, result); diff != "" {
						t.Errorf("-expected +received:\n%s", diff)
					}
				})
			})
			t.Run("Given an ID that ends with 4", func(t *testing.T) {
				var input []string = []string{
					"7194,b33,394,495593,192",
				}
				var expected []*UsageResult = []*UsageResult{
					NewExtendedUsage(7194, 495593, 192, "b33", 394),
				}

				t.Run("Then it will return extended string data", func(t *testing.T) {
					result := NewParser().Parse(input)

					if diff := cmp.Diff(expected, result); diff != "" {
						t.Errorf("-expected +received:\n%s", diff)
					}
				})
			})
			t.Run("Given an ID that ends with 6", func(t *testing.T) {
				var input []string = []string{
					"316,0e893279227712cac0014aff",
				}
				var expected []*UsageResult = []*UsageResult{
					NewHexUsage(316, 12921, 578228938, net.ParseIP("192.1.74.255"), 3721),
				}

				t.Run("Then it will return hex string data", func(t *testing.T) {
					result := NewParser().Parse(input)

					if diff := cmp.Diff(expected, result); diff != "" {
						t.Errorf("-expected +received:\n%s", diff)
					}
				})
			})
		})

		t.Run("Given an array of strings", func(t *testing.T) {
			var input []string = []string{
				"4,0d39f,0,495594,214",
				"16,be833279000000c063e5e63d",
				"9991,2935",
			}
			var expected []*UsageResult = []*UsageResult{
				NewExtendedUsage(4, 495594, 214, "0d39f", 0),
				NewHexUsage(16, 12921, 192, net.ParseIP("99.229.230.61"), 48771),
				NewBasicUsage(9991, 2935),
			}

			t.Run("Then it will parse each string according to its ID and return an array of data", func(t *testing.T) {
				result := NewParser().Parse(input)

				if diff := cmp.Diff(expected, result); diff != "" {
					t.Errorf("-expected +received:\n%s", diff)
				}
			})
		})
	})
}

package crypto

import (
	"reflect"
	"testing"

	"github.com/edznux/wonderxss/storage/models"
)

func TestGenerateSRIHashes(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want models.SRIHashes
	}{
		{
			name: "Test empty string",
			args: args{data: ""},
			want: models.SRIHashes{
				SHA256: "sha256-47DEQpj8HBSa+/TImW+5JCeuQeRkm5NMpJWZG3hSuFU=",
				SHA384: "sha384-OLBgp1GsljhM2TJ+sbHjaiH9txEUvgdDTAzHv2P24donTt6/529l+9Ua0vFImLlb",
				SHA512: "sha512-z4PhNX7vuL3xVChQ1m2AB9Yg5AULVxXcg/SpIdNs6c5H0NE8XYXysP+DGNKHfuwvY7kxvUdBeoGlODJ6+SfaPg==",
			},
		},
		{
			name: "Test number only string",
			args: args{data: "1234567890"},
			want: models.SRIHashes{
				SHA256: "sha256-x3Xnt1ft5jDNCqERO9ECZhqziCnKUqZCKreChi8mhkY=",
				SHA384: "sha384-7YRfi08qbV2oajvskDUtkW1qZuNCDXIOFkOa3yOPEpGCyMZPxOyMHmUGvCtIiLr5",
				SHA512: "sha512-ErAyJqbYvpxujNXlXcbHkgyqo53xSquS1ePqk0DRyKTT0LjkMU8fbvExukvxzrkYarh8gBrw1clbG++4ztriuQ==",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GenerateSRIHashes(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateSRIHashes() = %v, want %v", got, tt.want)
			}
		})
	}
}

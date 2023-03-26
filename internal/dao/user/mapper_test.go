package user

import "testing"

func Test_parseStructToSlices(t *testing.T) {
	tests := []struct {
		name  string
		obj   any
		want  string
		want1 string
	}{
		{
			name: "Happy Path",
			obj: struct {
				Name string `json:"name,omitempty" db:"name"`
				Job  string `json:"job,omitempty" db:"job"`
			}{Name: "NAME", Job: "JOB"},
			want:  "name, job",
			want1: "'NAME', 'JOB'",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseStructToSlices(tt.obj)
			if got != tt.want {
				t.Errorf("parseStructToSlices() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseStructToSlices() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

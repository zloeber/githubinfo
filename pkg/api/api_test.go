package api

import (
	"reflect"
	"testing"
)

func TestIsValidProject(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "success",
			args: args{
				str: "vendor/repo",
			},
			want: true,
		},
		{
			name: "failure",
			args: args{
				str: "repo",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidProject(tt.args.str); got != tt.want {
				t.Errorf("IsValidProject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestProjectJSON(t *testing.T) {
	type args struct {
		project string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				project: "zloeber/githubinfo",
			},
			want: "",
		},
		{
			name: "failure",
			args: args{
				project: "repo",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ProjectJSON(tt.args.project); got != tt.want {
				t.Errorf("ProjectJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReleasesJSON(t *testing.T) {
	type args struct {
		project string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReleasesJSON(tt.args.project); got != tt.want {
				t.Errorf("ReleasesJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDescription(t *testing.T) {
	type args struct {
		payload string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Description(tt.args.payload); got != tt.want {
				t.Errorf("Description() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLicense(t *testing.T) {
	type args struct {
		payload string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := License(tt.args.payload); got != tt.want {
				t.Errorf("License() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestReleaseURLs(t *testing.T) {
	type args struct {
		payload string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReleaseURLs(tt.args.payload); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReleaseURLs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsJSON(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsJSON(tt.args.str); got != tt.want {
				t.Errorf("IsJSON() = %v, want %v", got, tt.want)
			}
		})
	}
}

package main

import (
	"testing"
)

func TestFileExist(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "file exists",
			args: args{
				name: "README.md",
			},
			want: true,
		},
		{
			name: "file does not exist",
			args: args{
				name: "test.txt",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FileExist(tt.args.name); got != tt.want {
				t.Errorf("FileExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDirExist(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "dir does not exist",
			args: args{
				name: "testdir",
			},
			want: false,
		},
		{
			name: "dir does exist",
			args: args{
				name: "../touchx",
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DirExist(tt.args.name); got != tt.want {
				t.Errorf("DirExist() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsDir(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is dir",
			args: args{
				name: "../touchx",
			},
			want: true,
		},
		{
			name: "is not a dir",
			args: args{
				name: "README.md",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsDir(tt.args.name); got != tt.want {
				t.Errorf("IsDir() = %v, want %v", got, tt.want)
			}
		})
	}
}

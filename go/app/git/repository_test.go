package git

import (
	"os"
	"strings"
	"testing"
)

func TestOpen(t *testing.T) {
	type args struct {
		repositoryPath string
	}
	tests := []struct {
		name     string
		args     args
		want     string
		wantErr  bool
		errorMsg string
	}{
		{
			name: "success to open no remote repo",
			args: args{repositoryPath: "testdata/noremoterepo"},
			want: getWd() + "/testdata/noremoterepo",
		}, {
			name: "success to open one remote repo",
			args: args{repositoryPath: "testdata/oneremoterepo"},
			want: getWd() + "/testdata/oneremoterepo",
		}, {
			name: "success to open two remotes repo",
			args: args{repositoryPath: "testdata/tworemotesrepo"},
			want: getWd() + "/testdata/tworemotesrepo",
		}, {
			name: "failed to open denied config repo",
			args: args{repositoryPath: "testdata/deniedconfigrepo"},
			want: getWd() + "/testdata/deniedconfigrepo",
		}, {
			name:     "fail to open no repo directory",
			args:     args{repositoryPath: "testdata/norepo"},
			wantErr:  true,
			errorMsg: "repository does not exist",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Open(tt.args.repositoryPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Open() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err != nil && !strings.Contains(err.Error(), tt.errorMsg) {
				t.Errorf("Open() error = %v, must contain %v", err.Error(), tt.errorMsg)
				return
			}

			gotString := ""
			if got != nil {
				gotString = got.rootPath
			}

			if gotString != tt.want {
				t.Errorf("Open() = %v, want %v", got.rootPath, tt.want)
			}
		})
	}
}

func getWd() string {
	wd, _ := os.Getwd()
	return wd
}

func getNoRemoteRepoOpened() *Repository {
	repo, _ := Open("testdata/noremoterepo")
	return repo
}

func getOneRemoteRepoOpened() *Repository {
	repo, _ := Open("testdata/oneremoterepo")
	return repo
}

func getTwoRemotesRepoOpened() *Repository {
	repo, _ := Open("testdata/tworemotesrepo")
	return repo
}

func getDeniedConfigRepoOpened() *Repository {
	repo, _ := Open("testdata/deniedconfigrepo")
	return repo
}

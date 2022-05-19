package git

import (
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
		want     *Repository
		wantErr  bool
		errorMsg string
	}{
		{
			name: "success to open no remote repo",
			args: args{repositoryPath: "testdata/noremoterepo"},
			want: getNoRemoteRepoOpened(),
		}, {
			name: "success to open one remote repo",
			args: args{repositoryPath: "testdata/oneremoterepo"},
			want: getOneRemoteRepoOpened(),
		}, {
			name: "success to open two remotes repo",
			args: args{repositoryPath: "testdata/tworemotesrepo"},
			want: getTwoRemotesRepoOpened(),
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
			if !got.equals(tt.want) {
				t.Errorf("Open() got = %v, want %v", got, tt.want)
			}
		})
	}
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

package git

import (
	"github.com/go-git/go-git/v5/plumbing"
	"reflect"
	"sort"
	"testing"
)

func TestRemote_GetRefPrefix(t *testing.T) {
	tests := []struct {
		name   string
		remote func() *Remote
		want   string
	}{
		{name: "success to get first remote prefix", remote: getBackupRemote, want: "backup/"},
		{name: "success to get second remote prefix", remote: getSecondRemote, want: "origin/"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.remote()
			if got := r.GetRefPrefix(); got != tt.want {
				t.Errorf("GetRefPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemote_IsParentOfBranch(t *testing.T) {
	type args struct {
		ref *plumbing.Reference
	}
	tests := []struct {
		name   string
		remote func() *Remote
		args   args
		want   bool
	}{
		{
			name:   "success to check if remote is parent of branch",
			remote: getBackupRemote,
			args:   args{ref: plumbing.NewReferenceFromStrings("refs/remotes/backup/branch", "branch")},
			want:   true,
		}, {
			name:   "success to check if remote is not parent of branch",
			remote: getBackupRemote,
			args:   args{ref: plumbing.NewReferenceFromStrings("refs/remotes/origin/branch", "branch")},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.remote()
			if got := r.IsParentOfBranch(tt.args.ref); got != tt.want {
				t.Errorf("IsParentOfBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemote_Name(t *testing.T) {
	tests := []struct {
		name   string
		remote func() *Remote
		want   string
	}{
		{name: "success to get first remote name", remote: getBackupRemote, want: "backup"},
		{name: "success to get second remote name", remote: getSecondRemote, want: "origin"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.remote()
			if got := r.Name(); got != tt.want {
				t.Errorf("Name() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetRemotes(t *testing.T) {
	tests := []struct {
		name       string
		repository func() *Repository
		want       []string
		wantErr    bool
	}{
		{name: "success to get no remote", repository: getNoRemoteRepoOpened, want: []string{}},
		{name: "success to get one remote", repository: getOneRemoteRepoOpened, want: []string{"origin"}},
		{name: "success to get two remotes", repository: getTwoRemotesRepoOpened, want: []string{"backup", "origin"}},
		{name: "success to get one remote", repository: getDeniedConfigRepoOpened, want: nil, wantErr: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := tt.repository()
			got, err := r.GetRemotes()
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRemotes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got == nil {
				if tt.want != nil {
					t.Errorf("GetRemotes() = nil, want %v", tt.want)
				}
			} else {
				gotAsString := make([]string, len(got))
				for i, remote := range got {
					gotAsString[i] = remote.Name()
				}
				sort.Strings(gotAsString)
				if !reflect.DeepEqual(gotAsString, tt.want) {
					t.Errorf("GetRemotes() = %v, want %v", gotAsString, tt.want)
				}
			}
		})
	}
}

func getRemotes() []*Remote {
	repo, _ := Open("testdata/tworemotesrepo")

	remotes, _ := repo.GetRemotes()
	sort.Slice(remotes, func(i, j int) bool {
		return remotes[i].Name() < remotes[j].Name()
	})

	return remotes
}

func getBackupRemote() *Remote {

	return getRemotes()[0]
}

func getSecondRemote() *Remote {

	return getRemotes()[1]
}

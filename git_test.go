package main

import "testing"

func Test_GenerateRepoURL(t *testing.T) {
	tests := []struct {
		name        string
		args        Args
		org, repo   string
		expectedURL string
	}{
		{
			name: "GitHub URL",
			args: Args{
				Enterprise: false,
			},
			org:         "exampleOrg",
			repo:        "exampleRepo",
			expectedURL: "https://github.com/exampleOrg/exampleRepo.git",
		},
		{
			name: "Enterprise URL",
			args: Args{
				Enterprise:    true,
				EnterpriseUrl: "https://enterprise.example.com",
			},
			org:         "exampleOrg",
			repo:        "exampleRepo",
			expectedURL: "https://enterprise.example.com/exampleOrg/exampleRepo.git",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args = tt.args
			got := GenerateRepoURL(tt.org, tt.repo)
			if got != tt.expectedURL {
				t.Errorf("generateRepoURL() = %v, want %v", got, tt.expectedURL)
			}
		})
	}
}

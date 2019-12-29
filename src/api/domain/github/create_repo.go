package github

/*
{
  "name": "golang-rest-api",
  "description": "golang someting",
  "homepage": "https://github.com",
  "private": true,
  "has_issues": true,
  "has_projects": true,
  "has_wiki": true
}
*/

type CreateRepoRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Homepage    string `json:"homepage"`
	Private     bool   `json:"private"`
	HasIssues   bool   `json:"has_issues"`
	HasProjects bool   `json:"has_projects"`
	HasWiki     bool   `json:"has_wiki"`
}

type CreateRepoResponse struct {
	Id          int         `json:"id"`
	Name        string      `json:"name"`
	FullName    string      `json:"full_name"`
	Owner       RepoOwner   `json:"owner"`
	Permissions Permissions `json:"permissions"`
}

type RepoOwner struct {
	Login   string `json:"login"`
	Id      int    `json:"id"`
	Url     string `json:"url"`
	HtmlUrl string `json:"html_url"`
}

type Permissions struct {
	Admin bool `json:"admin"`
	Push  bool `json:"push"`
	Pull  bool `json:"pull"`
}

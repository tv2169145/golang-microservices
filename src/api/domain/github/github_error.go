package github

/*
{
    "message": "Repository creation failed.",
    "errors": [
        {
            "resource": "Repository",
            "code": "custom",
            "field": "name",
            "message": "name already exists on this account"
        }
    ],
    "documentation_url": "https://developer.github.com/v3/repos/#create"
}
*/

type GithubErrorResponse struct {
	StatusCode       int           `json:"status_code"`
	Message          string        `json:"message"`
	Errors           []GithubError `json:"errors"`
	DocumentationUrl string        `json:"documentation_url"`
}

type GithubError struct {
	Resource string `json:"resource"`
	Code     string `json:"code"`
	Field    string `json:"field"`
	Message  string `json:"message"`
}

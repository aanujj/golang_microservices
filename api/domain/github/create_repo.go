package github

/*func CreateRepo(){
	request := map[string]interface{}{
		"name":"Hello Github From Gin",
		"private":false,
	}
}*/

type CreateRepoRequest struct {
	Name string
	Description string
	Homepage string
	Private bool
	HasIssues bool
	HasProjects bool
	HasWiki bool
}
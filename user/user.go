/*
GITHUB API VERSION 3

http://developer.github.com/v3/users/

TODO: Add save users
TODO: Add Last userid seen as optional paramter to GetAll
TODO: Use link headers to implement paging.
*/

package user

import (
	"gohub/client"
)

type UserController struct {
	Client *client.GithubClient
}

type User struct {
	Login        string
	Id           int
	Avatar_url   string
	Gravatar_id  string
	Url          string
	Name         string
	Company      string
	Blog         string
	Location     string
	Email        string
	Hireable     bool
	Bio          string
	Public_repos int
	Public_gists int
	Followers    int
	Following    int
	Html_url     string
	Created_at   string
	Type         string
}

type CurrentUser struct {
	Login               string
	Id                  int
	Avatar_url          string
	Gravatar_id         string
	Url                 string
	Html_url            string
	Gists_url           string
	Starred_url         string
	Subscriptions_url   string
	Organizations_url   string
	Repos_url           string
	Events_url          string
	Received_events_url string
	Type                string
	Name                string
	Company             string
	Blog                string
	Location            string
	Email               string
	Hireable            bool
	Bio                 string
	Public_repos        int
	Followers           int
	Following           int
	Created_at          string
	Updated_at          string
	Public_gists        int
	Total_private_repos int
	Owned_private_repos int
	Disk_usage          int
	Collaborators       int
	Plan                Plan
	Private_gists       int
	Site_admin          bool
}

type Plan struct {
	Name          string
	Space         int
	Collaborators int
	Private_repos int
}

const gh_user_url = "user"
const gh_users_url = "users"

func (gc *UserController) GetCurrent() (*CurrentUser, error) {
	user, err := gc.Client.MakeRequest(gh_user_url, new(CurrentUser))

	return user.(*CurrentUser), err
}

func (gc *UserController) Get(username string) (*User, error) {
	user, err := gc.Client.MakeRequest(gh_users_url+"/"+username, new(User))

	return user.(*User), err
}

func (gc *UserController) GetAll() (*[]User, error) {
	users, err := gc.Client.MakeRequest(gh_users_url, new([]User))

	return users.(*[]User), err
}

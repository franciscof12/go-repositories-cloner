package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type repo struct {
	Name string `json:"name"`
	Url  string `json:"clone_url"`
}

func (r repo) Title() string       { return r.Name }
func (r repo) Description() string { return r.Url }
func (r repo) FilterValue() string { return r.Name }

type model struct {
	list list.Model
}

func (m model) Init() tea.Cmd {
	return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			selectedItem := m.list.SelectedItem()
			if r, ok := selectedItem.(repo); ok {
				repoUrl := r.Description()
				err := cloneRepo(repoUrl)
				if err != nil {
					fmt.Println("Error cloning repository:", err)
					return m, tea.Quit
				}
				fmt.Printf("Successfully cloned %s\n", repoUrl)
				return m, tea.Quit
			}
		}
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	return docStyle.Render(m.list.View())
}

func fetchUserRepos(username string) ([]repo, error) {
	resp, err := http.Get(fmt.Sprintf("https://api.github.com/users/%s/repos", username))

	body, err := ioutil.ReadAll(resp.Body)

	var repos []repo
	if err = json.Unmarshal(body, &repos); err != nil {
		return nil, err
	}
	return repos, nil
}

func cloneRepo(url string) error {
	cmd := exec.Command("git", "clone", url)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	fmt.Println("Enter a GitHub username and press Enter to get their repositories:")
	var username string
	fmt.Scanln(&username)

	repos, err := fetchUserRepos(username)
	if err != nil {
		fmt.Println("Error fetching repositories:", err)
		os.Exit(1)
	}

	items := make([]list.Item, len(repos))
	for i, repo := range repos {
		items[i] = repo
	}

	m := model{list: list.New(items, list.NewDefaultDelegate(), 0, 0)}
	m.list.Title = fmt.Sprintf("Repos of %s", username)

	p := tea.NewProgram(m, tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}

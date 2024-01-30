# go-repositories-cloner 

A simple CLI tool built in Go that fetches the list of repositories for a given GitHub username. The user is presented with a TUI (Text-based User Interface) from which they can navigate and select a repository to clone it to the local machine.
<img width="958" alt="Captura de pantalla 2023-10-19 a las 23 23 19" src="https://github.com/franciscof12/go-repositories-cloner/assets/123760628/7aa63b3a-3190-49ec-81c2-c24de3bcc1de">
<img width="748" alt="Captura de pantalla 2023-10-19 a las 23 23 38" src="https://github.com/franciscof12/go-repositories-cloner/assets/123760628/0c40b02a-5b3d-4f83-aff6-dcbe50e888ec">

# Features:
1. Fetches repositories of any GitHub user by their username.
2. Interactive TUI for displaying and navigating through the repositories using the bubbletea package.
3. Stylish rendering with the lipgloss package.
4. Option to clone the selected repository with a simple press of the enter key.
   
# How to Use:
1. Run the program.
2. Enter the desired GitHub username.
3. Navigate through the list of repositories using arrow keys.
4. Press enter to clone the highlighted repository.
5. Exit the program using ctrl+c.
   
# Dependencies:

1. bubbletea: For creating the TUI.
2. lipgloss: For stylish interface rendering.
3. The standard Go library.

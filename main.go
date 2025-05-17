package main

import "fmt"

type Team struct {
	Name   string
	Wins   int
	Losses int
	Points int
}

func addTeam(teams []Team, team Team) []Team {
	return append(teams, team)
}

func updateTeam(teams []Team, name string, newWins, newLosses, newPoints int) []Team {
	for i := range teams {
		if teams[i].Name == name {
			teams[i] = Team{name, newWins, newLosses, newPoints}
		}
	}
	return teams
}

func deleteTeam(teams []Team, name string) []Team {
	for i := range teams {
		if teams[i].Name == name {
			return append(teams[:i], teams[i+1:]...)
		}
	}
	return teams
}

func displayRanking(teams []Team) {
	selectionSort(teams)
	fmt.Println("\nKlasemen:")
	for i, team := range teams {
		fmt.Printf("%d. %s - Wins: %d, Losses: %d, Points: %d\n", i+1, team.Name, team.Wins, team.Losses, team.Points)
	}
}

func searchTeamSequential(teams []Team, name string) Team {
	for _, team := range teams {
		if team.Name == name {
			return team
		}
	}
	return Team{}
}

func selectionSort(teams []Team) {
	for i := range teams {
		maxIdx := i
		for j := i + 1; j < len(teams); j++ {
			if teams[j].Points > teams[maxIdx].Points {
				maxIdx = j
			}
		}
		teams[i], teams[maxIdx] = teams[maxIdx], teams[i]
	}
}

func showMenu() {
	fmt.Println("\nMenu:\n1. Klasemen\n2. Cari Tim\n3. Tambah Tim\n4. Update Tim\n5. Hapus Tim\n6. Keluar")
	fmt.Print("Pilih menu: ")
}

func main() {
	var teams []Team
	var choice int

	for {
		showMenu()
		fmt.Scan(&choice)

		switch choice {
		case 1:
			displayRanking(teams)
		case 2:
			var name string
			fmt.Print("Masukkan nama tim: ")
			fmt.Scan(&name)
			team := searchTeamSequential(teams, name)
			if (team != Team{}) {
				fmt.Printf("Tim ditemukan: %s - Wins: %d, Losses: %d, Points: %d\n", team.Name, team.Wins, team.Losses, team.Points)
			} else {
				fmt.Println("Tim tidak ditemukan.")
			}
		case 3:
			var name string
			var wins, losses, points int
			fmt.Print("Nama tim: ")
			fmt.Scan(&name)
			fmt.Print("Kemenangan, Kekalahan, Poin: ")
			fmt.Scan(&wins, &losses, &points)
			teams = addTeam(teams, Team{name, wins, losses, points})
		case 4:
			var name string
			var wins, losses, points int
			fmt.Print("Nama tim yang diupdate: ")
			fmt.Scan(&name)
			fmt.Print("Kemenangan, Kekalahan, Poin baru: ")
			fmt.Scan(&wins, &losses, &points)
			teams = updateTeam(teams, name, wins, losses, points)
		case 5:
			var name string
			fmt.Print("Nama tim yang dihapus: ")
			fmt.Scan(&name)
			teams = deleteTeam(teams, name)
		case 6:
			fmt.Println("Program selesai.")
			return
		default:
			fmt.Println("Pilihan tidak valid.")
		}
	}
}
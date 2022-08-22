package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl/ch04/github"
)

func main() {
	today := time.Now()
	queryPast1Month := fmt.Sprintf("created:>%s", today.AddDate(0, -1, 0).Format(time.RFC3339))
	queryPast1Year := fmt.Sprintf("created:>%s", today.AddDate(-1, 0, 0).Format(time.RFC3339))
	queryOver1Year := fmt.Sprintf("created:<%s", today.AddDate(-1, 0, 0).Format(time.RFC3339))

	result, err := github.SearchIssues(append(os.Args[1:], queryPast1Month))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("########")
	fmt.Printf("# %s\n", queryPast1Month)
	fmt.Println("########")
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	result, err = github.SearchIssues(append(os.Args[1:], queryPast1Year))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("########")
	fmt.Printf("# %s\n", queryPast1Year)
	fmt.Println("########")
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}

	result, err = github.SearchIssues(append(os.Args[1:], queryOver1Year))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("########")
	fmt.Printf("# %s\n", queryOver1Year)
	fmt.Println("########")
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}

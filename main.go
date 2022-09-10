package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"

	"github.com/samlhuillier/say/cmd"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "db",
	Long: "Root command",
}
var HelloCmd = &cobra.Command{
	Use:   "hello <name>",
	Short: "Get user data",
	Args:  cobra.ExactValidArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Hello %s!!!\n", args[0])
	},
	ValidArgsFunction: UserGet,
}

func getNames() []string {
	file, err := os.Open("common-names.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var names []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		names = append(names, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return names
}

func UserGet(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	rand.Seed(time.Now().UnixNano())
	return getNames(), cobra.ShellCompDirectiveNoFileComp //
}
func init() {
	RootCmd.AddCommand(HelloCmd, cmd.CompletionCmd)
}
func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

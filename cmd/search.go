package cmd

import (
	"context"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/google/generative-ai-go/genai"
	"github.com/spf13/cobra"
	"google.golang.org/api/option"
)

var numWords string = "150"

var searchCmd = &cobra.Command{
	Use:   "search [question]",
	Short: "Ask a question and get a response",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		res := getApiResponse(args)
		fmt.Println(res)
	},
}

func getApiResponse(args []string) string {
	userArgs := strings.Join(args[0:], " ")
	ctx := context.Background()
	client, err := genai.NewClient(ctx, option.WithAPIKey(os.Getenv("GEMINI_API_KEY")))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	_, err = strconv.Atoi(numWords)
	if err != nil {
		log.Fatalf("invalid number of words")
	}

	model := client.GenerativeModel("gemini-1.5-flash")
	resp, err := model.GenerateContent(ctx, genai.Text(userArgs+"in"+numWords+" words."))
	if err != nil {
		log.Fatalf("Failed to generate content: %v", err)
	}

	finalResponse := resp.Candidates[0].Content.Parts[0]
	return formatAsText(fmt.Sprint(finalResponse))
}

func formatAsText(input string) string {

	re := regexp.MustCompile(`^#{1,6}\s+(.*)`)
	input = re.ReplaceAllString(input, "$1")

	re = regexp.MustCompile(`\n---\n`)
	input = re.ReplaceAllString(input, "\n")

	re = regexp.MustCompile(`_([^_]+)_`)
	input = re.ReplaceAllString(input, "$1")

	re = regexp.MustCompile(`\*\*([^*]+)\*\*`)
	input = re.ReplaceAllString(input, "$1")

	re = regexp.MustCompile(`\n\* (.*)`)
	input = re.ReplaceAllString(input, "\n- $1")

	return input
}

func init() {
	searchCmd.Flags().StringVarP(&numWords, "words", "w", "150", "number of words")
}

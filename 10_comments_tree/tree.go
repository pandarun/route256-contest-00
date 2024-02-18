package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strconv"
)

var commentRegexp = `^(\d+)\s(-{0,1}\d+)\s(.+)$`

var compilesCommentRegex = regexp.MustCompile(commentRegexp)

func parseNumber9(item string) (int, error) {
	number, err := strconv.Atoi(item)
	if err != nil {
		fmt.Println("unable to parse int")
		return 0, fmt.Errorf("unable to parse int")
	}
	return number, nil
}

type Comment struct {
	id       int
	parentId int
	text     string
	replies  []*Comment
}

func tree(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)

	in.Scan()
	dataInputsAsStrings := in.Text()

	dataInputs, err := parseNumber9(dataInputsAsStrings)
	if err != nil {
		return fmt.Errorf("unable to parse data inputs length")
	}

	var dataInput = 0
	for dataInput < dataInputs {

		in.Scan()
		commentsLength := in.Text()

		n, err := parseNumber9(commentsLength)
		if err != nil {
			return fmt.Errorf("unable to parse n length")
		}

		comments := make([]*Comment, n)
		commentsById := make(map[int]*Comment)
		roots := []*Comment{}
		for i := 0; i < n; i++ {

			in.Scan()
			commentString := in.Text()

			matches := compilesCommentRegex.FindStringSubmatch(commentString)
			if len(matches) != 4 {
				return fmt.Errorf("unable to parse comment")
			}
			id, err := parseNumber9(matches[1])
			if err != nil {
				return fmt.Errorf("unable to parse id")
			}
			parenId, err := parseNumber9(matches[2])

			c := &Comment{id: id, parentId: parenId, text: matches[3]}
			comments[i] = c
			commentsById[id] = c
		}

		sort.Slice(comments, func(i, j int) bool {
			return comments[i].id < comments[j].id
		})

		for _, comment := range comments {
			if comment.parentId == -1 {
				roots = append(roots, comment)
			} else {
				parent := commentsById[comment.parentId]
				parent.replies = append(parent.replies, comment)
			}
		}

		var buffer bytes.Buffer
		for _, root := range roots {
			printComment(root, &buffer, 0, []bool{})
			buffer.WriteString("\n")
		}

		fmt.Fprint(output, buffer.String())

		dataInput++
	}

	return nil
}

func printComment(comment *Comment, buffer *bytes.Buffer, level int, printVertical []bool) {

	for i := 0; i < level-1; i++ {

		if level > 1 && printVertical[i] {
			buffer.WriteString("|")
		} else {
			buffer.WriteString(" ")
		}

		buffer.WriteString("  ")
	}

	if level > 0 {
		buffer.WriteString("|\n")
	}

	for i := 0; i < level-1; i++ {

		if level > 1 && printVertical[i] {
			buffer.WriteString("|")
		} else {
			buffer.WriteString(" ")
		}

		buffer.WriteString("  ")
	}

	if level > 0 {
		buffer.WriteString("|--")
	}

	buffer.WriteString(comment.text)
	buffer.WriteString("\n")
	for idx, reply := range comment.replies {

		printVertical = append(printVertical, len(comment.replies) > 1)

		printVertical[level] = idx < len(comment.replies)-1

		printComment(reply, buffer, level+1, printVertical)
	}
}

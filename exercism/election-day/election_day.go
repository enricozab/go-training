// Package electionday implements a new digital system to count the votes
package electionday

import "fmt"

// NewVoteCounter returns a new vote counter with
// a given number of inital votes.
func NewVoteCounter(initialVotes int) *int {
	counter := &initialVotes

	return counter
}

// VoteCount extracts the number of votes from a counter.
func VoteCount(counter *int) int {
	if counter == nil {
		return 0
	}

	votes := *counter

	return votes
}

// IncrementVoteCount increments the value in a vote counter
func IncrementVoteCount(counter *int, increment int) {
	votes := VoteCount(counter)
	*counter = votes + increment
}

// NewElectionResult creates a new election result

func NewElectionResult(candidateName string, votes int) *ElectionResult {
	return &ElectionResult{Name: candidateName, Votes: votes}
}

// DisplayResult creates a message with the result to be displayed
func DisplayResult(result *ElectionResult) string {
	temp := *result

	return fmt.Sprintf("%v (%v)", temp.Name, temp.Votes)
}

// DecrementVotesOfCandidate decrements by one the vote count of a candidate in a map
func DecrementVotesOfCandidate(results map[string]int, candidate string) {
	for key, val := range results {
		if key == candidate {
			results[key] = val - 1
		}
	}
}

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	problems := make([]string, 75)
	index := 0

	s := bufio.NewScanner(strings.NewReader(leetcodeList))
	for s.Scan() {
		problems[index] = s.Text()
		index++
	}

	// Fisher–Yates shuffle
	for i := len(problems) - 1; i >= 1; i-- {
		j := rand.Intn(i)
		temp := problems[j]
		problems[j] = problems[i]
		problems[i] = temp
	}

	for i, problem := range problems {
		fmt.Printf("%d. %s\n", i+1, problem)
	}
}

const leetcodeList = `Merge Strings Alternately
Greatest Common Divisor of Strings
Kids With the Greatest Number of Candies
Can Place Flowers
Reverse Vowels of a String
Reverse Words in a String
Product of Array Except Self
Increasing Triplet Subsequence
String Compression
Move Zeroes
Is Subsequence
Container With Most Water
Max Number of K-Sum Pairs
Maximum Average Subarray I
Maximum Number of Vowels in a Substring of Given Length
Max Consecutive Ones III
Longest Subarray of 1's After Deleting One Element
Find the Highest Altitude
Find Pivot Index
Find the Difference of Two Arrays
Unique Number of Occurrences
Determine if Two Strings Are Close
Equal Row and Column Pairs
Removing Stars From a String
Asteroid Collision
Decode String
Number of Recent Calls
Dota2 Senate
Delete the Middle Node of a Linked List
Odd Even Linked List
Reverse Linked List
Maximum Twin Sum of a Linked List
Maximum Depth of Binary Tree
Leaf-Similar Trees
Count Good Nodes in Binary Tree
Path Sum III
Longest ZigZag Path in a Binary Tree
Lowest Common Ancestor of a Binary Tree
Binary Tree Right Side View
Maximum Level Sum of a Binary Tree
Search in a Binary Search Tree
Delete Node in a BST
Keys and Rooms
Number of Provinces
Reorder Routes to Make All Paths Lead to the City Zero
Evaluate Division
Nearest Exit from Entrance in Maze
Rotting Oranges
Kth Largest Element in an Array
Smallest Number in Infinite Set
Maximum Subsequence Score
Total Cost to Hire K Workers
Guess Number Higher or Lower
Successful Pairs of Spells and Potions
Find Peak Element
Koko Eating Bananas
Letter Combinations of a Phone Number
Combination Sum III
N-th Tribonacci Number
Min Cost Climbing Stairs
House Robber
Domino and Tromino Tiling
Unique Paths
Longest Common Subsequence
Best Time to Buy and Sell Stock with Transaction Fee
Edit Distance
Counting Bits
Single Number
Minimum Flips to Make a OR b Equal to c
Implement Trie (Prefix Tree)
Search Suggestions System
Non-overlapping Intervals
Minimum Number of Arrows to Burst Balloons
Daily Temperatures
Online Stock Span`

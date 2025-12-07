package main

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type RunRequest struct {
	Code string `json:"code"`
}

type RunResponse struct {
	Status string `json:"status"`
	Output string `json:"output"`
}

type FixRequest struct {
	Code string `json:"code"`
}

type FixResponse struct {
	FixedCode string `json:"fixed_code"`
}

type HelpRequest struct {
	Query string `json:"query"`
}

type HelpResponse struct {
	Help string `json:"help"`
}

func main() {
	app := fiber.New(fiber.Config{
		AppName: "Code Assist API",
	})

	// Serve static files
	app.Static("/", "./static")

	// API Routes
	app.Post("/run", handleRun)
	app.Post("/autofix", handleAutoFix)
	app.Post("/help", handleHelp)

	// Serve index.html for the root path
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("./static/index.html")
	})

	// Start server
	address := ":3001"
	fmt.Printf("Server running on http://localhost%s\n", address)
	if err := app.Listen(address); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

func handleRun(c *fiber.Ctx) error {
	var req RunRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(RunResponse{
			Status: "error",
			Output: fmt.Sprintf("Invalid request format: %v", err),
		})
	}

	// Simulate code execution (basic syntax checking)
	if strings.TrimSpace(req.Code) == "" {
		return c.JSON(RunResponse{
			Status: "error",
			Output: "Empty code provided",
		})
	}

	// Check for common syntax errors
	if strings.Contains(req.Code, "func main()") && !strings.Contains(req.Code, "package main") {
		return c.JSON(RunResponse{
			Status: "error",
			Output: "Missing 'package main' declaration",
		})
	}

	// If no errors found
	return c.JSON(RunResponse{
		Status: "success",
		Output: "Code executed successfully (simulated)",
	})
}

func handleAutoFix(c *fiber.Ctx) error {
	var req FixRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(FixResponse{
			FixedCode: fmt.Sprintf("Invalid request format: %v", err),
		})
	}

	if req.Code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(FixResponse{
			FixedCode: "No code provided to fix",
		})
	}

	fixed := autoFixCode(req.Code)
	return c.JSON(FixResponse{
		FixedCode: fixed,
	})
}

func handleHelp(c *fiber.Ctx) error {
	req := new(HelpRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(HelpResponse{
			Help: fmt.Sprintf("Invalid request format: %v", err),
		})
	}

	if req.Query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(HelpResponse{
			Help: "No query provided. Please provide a topic to get help with.",
		})
	}

	helpText := getHelpText(req.Query)
	return c.JSON(HelpResponse{
		Help: helpText,
	})
}

func autoFixCode(code string) string {
	// Remove trailing whitespace from each line
	lines := strings.Split(code, "\n")
	for i, line := range lines {
		lines[i] = strings.TrimRight(line, " \t")
	}

	// Join lines back together
	fixed := strings.Join(lines, "\n")

	// Fix common bracket mismatches (simple version)
	openBraces := strings.Count(fixed, "{")
	closeBraces := strings.Count(fixed, "}")
	if openBraces > closeBraces {
		fixed += strings.Repeat("\n}", openBraces-closeBraces)
	}

	// Normalize line endings
	fixed = strings.ReplaceAll(fixed, "\r\n", "\n")
	fixed = strings.ReplaceAll(fixed, "\r", "\n")

	return fixed
}

func getHelpText(query string) string {
	switch strings.ToLower(query) {
	case "error":
		return "Common errors include missing semicolons, undefined variables, or incorrect package imports."
	case "loop":
		return "In Go, use 'for' for loops. Example: for i := 0; i < 5; i++ { ... }"
	case "function":
		return "Define functions with 'func' keyword. Example: func add(a, b int) int { return a + b }"
	case "syntax":
		return "Go syntax requires proper formatting. Use 'go fmt' to format your code."
	default:
		return "I can help with: error, loop, function, syntax. Ask about any of these topics."
	}
}

// Copyright 2025 Erst Users
// SPDX-License-Identifier: Apache-2.0

package visualizer

import (
	"os"
	"strings"

	"github.com/dotandev/hintents/internal/terminal"
)

var defaultRenderer terminal.Renderer = terminal.NewANSIRenderer()

// ColorEnabled reports whether ANSI color output should be used.
// Checks NO_COLOR and TERM=dumb environment variables on every call
// so that tests can control color via env vars dynamically.
func ColorEnabled() bool {
	// NO_COLOR takes precedence over everything
	if _, ok := os.LookupEnv("NO_COLOR"); ok {
		return false
	}
	// FORCE_COLOR enables colors unconditionally
	if os.Getenv("FORCE_COLOR") != "" {
		return true
	}
	// dumb terminal
	if os.Getenv("TERM") == "dumb" {
		return false
	}
	return defaultRenderer.IsTTY()
}

// colorMap maps color names to ANSI SGR codes.
var colorMap = map[string]string{
	"red":     sgrRed,
	"green":   sgrGreen,
	"yellow":  sgrYellow,
	"blue":    sgrBlue,
	"magenta": sgrMagenta,
	"cyan":    sgrCyan,
	"bold":    sgrBold,
	"dim":     sgrDim,
}

// Colorize returns text with ANSI color if enabled, otherwise plain text.
func Colorize(text string, color string) string {
	if !ColorEnabled() {
		return text
	}
	code, ok := colorMap[strings.ToLower(color)]
	if !ok {
		return text
	}
	return code + text + sgrReset
}

// ContractBoundary returns a visual separator for cross-contract call transitions.
func ContractBoundary(fromContract, toContract string) string {
	if ColorEnabled() {
		return sgrMagenta + sgrBold + "--- contract boundary: " + fromContract + " -> " + toContract + " ---" + sgrReset
	}
	return "--- contract boundary: " + fromContract + " -> " + toContract + " ---"
}

// Success returns a success indicator: colored checkmark if enabled, "[OK]" otherwise.
func Success() string {
	if ColorEnabled() {
		return themeColors("success") + "[OK]" + sgrReset
	}
	return "[OK]"
}

// Warning returns a warning indicator.
func Warning() string {
	if ColorEnabled() {
		return themeColors("warning") + "[!]" + sgrReset
	}
	return "[!]"
}

// Error returns an error indicator.
func Error() string {
	if ColorEnabled() {
		return themeColors("error") + "[X]" + sgrReset
	}
	return "[X]"
}

// Info returns an info indicator with theme-aware coloring.
func Info() string {
	if ColorEnabled() {
		return themeColors("info") + "[i]" + sgrReset
	}
	return "[i]"
}

// Symbol returns a symbol that may be styled; when colors disabled, returns plain ASCII equivalent.
//
//nolint:gocyclo
func Symbol(name string) string {
	if ColorEnabled() {
		return defaultRenderer.Symbol(name)
	}
	// Return plain ASCII equivalents (no ANSI, no Unicode symbols)
	switch name {
	case "check":
		return "[OK]"
	case "cross":
		return "[X]"
	case "warn":
		return "[!]"
	case "arrow_r":
		return "->"
	case "arrow_l":
		return "<-"
	case "target":
		return ">>"
	case "pin":
		return "*"
	case "wrench":
		return "[*]"
	case "chart":
		return "[#]"
	case "list":
		return "[.]"
	case "play":
		return ">"
	case "book":
		return "[?]"
	case "wave":
		return "~"
	case "magnify":
		return "[?]"
	case "logs":
		return "[Logs]"
	case "events":
		return "[Events]"
	default:
		return name
	}
}

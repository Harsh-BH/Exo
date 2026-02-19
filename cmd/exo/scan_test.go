package exo

import (
	"os"
	"path/filepath"
	"testing"
)

// ── isBinaryFile ─────────────────────────────────────────────────────────────

func TestIsBinaryFile_TextFile(t *testing.T) {
	f := tempFile(t, "hello, world\nno null bytes here\n")
	if isBinaryFile(f) {
		t.Error("expected text file to NOT be detected as binary")
	}
}

func TestIsBinaryFile_BinaryFile(t *testing.T) {
	// Write a file that contains a null byte — classic binary signature
	tmp := filepath.Join(t.TempDir(), "binary.bin")
	content := []byte{'H', 'e', 'l', 'l', 'o', 0x00, 'W', 'o', 'r', 'l', 'd'}
	if err := os.WriteFile(tmp, content, 0600); err != nil {
		t.Fatal(err)
	}
	if !isBinaryFile(tmp) {
		t.Error("expected file with null byte to be detected as binary")
	}
}

func TestIsBinaryFile_Missing(t *testing.T) {
	// A missing file should return false (not binary)
	if isBinaryFile("/nonexistent/path/file.bin") {
		t.Error("expected missing file to return false")
	}
}

func TestIsBinaryFile_EmptyFile(t *testing.T) {
	f := tempFile(t, "")
	// Empty files have no null bytes — not binary
	if isBinaryFile(f) {
		t.Error("expected empty file to NOT be detected as binary")
	}
}

// ── secretPatterns ────────────────────────────────────────────────────────────

func TestSecretPatterns_AWSAccessKey(t *testing.T) {
	line := `AWS_ACCESS_KEY_ID=AKIAIOSFODNN7EXAMPLE`
	if !matchesAnyPattern(line) {
		t.Error("expected AWS access key pattern to match")
	}
}

func TestSecretPatterns_PrivateKey(t *testing.T) {
	line := `-----BEGIN RSA PRIVATE KEY-----`
	if !matchesAnyPattern(line) {
		t.Error("expected private key pattern to match")
	}
}

func TestSecretPatterns_GitHubToken(t *testing.T) {
	// gh[pousr]_ followed by 36+ [A-Za-z0-9_] chars
	line := `token=ghp_ABCDEFGHIJKLMNOPQRSTUVWXYZabcdef0123`
	if !matchesAnyPattern(line) {
		t.Error("expected GitHub token pattern to match")
	}
}

func TestSecretPatterns_GoogleAPIKey(t *testing.T) {
	// AIza followed by exactly 35 [A-Za-z0-9\-_] chars (total 39 chars)
	line := `api_key = AIzaSyDJR3AnotARealKeyButLooksLike1XXXX`
	if !matchesAnyPattern(line) {
		t.Error("expected Google API key pattern to match")
	}
}

func TestSecretPatterns_NoMatch(t *testing.T) {
	line := `PORT=8080`
	if matchesAnyPattern(line) {
		t.Errorf("expected plain PORT assignment NOT to match any secret pattern")
	}
}

// ── safeValues tightening ─────────────────────────────────────────────────────

func TestSafeValues_DoesNotContainSecret(t *testing.T) {
	for _, sv := range safeValues {
		if sv == "secret" {
			t.Error("'secret' should NOT be in safeValues — it would suppress real secret detections")
		}
		if sv == "password" {
			t.Error("'password' should NOT be in safeValues — it would suppress real password detections")
		}
	}
}

func TestSafeValues_ContainsExpectedPlaceholders(t *testing.T) {
	expected := []string{"change-me", "placeholder", "example", "${"}
	svSet := make(map[string]bool)
	for _, sv := range safeValues {
		svSet[sv] = true
	}
	for _, e := range expected {
		if !svSet[e] {
			t.Errorf("expected safeValues to contain %q", e)
		}
	}
}

// ── helpers ───────────────────────────────────────────────────────────────────

// matchesAnyPattern returns true if line matches at least one secretPattern.
func matchesAnyPattern(line string) bool {
	for _, p := range secretPatterns {
		if p.pattern.MatchString(line) {
			return true
		}
	}
	return false
}

// tempFile writes content to a temp file and returns its path.
func tempFile(t *testing.T, content string) string {
	t.Helper()
	tmp := filepath.Join(t.TempDir(), "testfile.txt")
	if err := os.WriteFile(tmp, []byte(content), 0600); err != nil {
		t.Fatal(err)
	}
	return tmp
}

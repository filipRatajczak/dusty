package main

// CleanablePaths defines safe directories that can be cleaned without harming user data.
//
// The following directories are generally safe to clean:
//
//   - ~/Library/Caches/                   → Application cache (user-level), always safe to delete.
//   - ~/Library/Logs/                     → Log files from apps, often large and unimportant.
//   - ~/.Trash/                           → User’s trash folder; files pending deletion.
//   - ~/Downloads/                        → Downloaded files; consider deleting only files older than X days.
//   - ~/Library/Containers/*/Data/Library/Caches → App sandbox cache (used by apps from App Store).
//   - ~/Library/Safari/Favicon Cache/     → Safari favicons, harmless to delete.
//   - /private/tmp/ and /tmp/             → Temporary files; clean only files owned by the current user and older than 24–48h.
//   - /private/var/folders/.../T/         → Per-user temp files; similar to /tmp/.
//
// ⚠️ DO NOT clean the following unless the user explicitly confirms:
//
//   - ~/Library/Preferences/              → App preferences/config files.
//   - ~/Library/Application Support/      → App data, may contain plugins, sessions, saved work.
//   - ~/Library/Keychains/                → Passwords and security credentials.
//   - ~/Library/Mobile Documents/         → iCloud files.
//   - /System, /Library, /Applications    → Core system directories.
//   - Any files modified in the last 24h  → May be actively used by running apps.
//
// Recommendation:
//   - Always support a "dry-run" or "report only" mode.
//   - Log deleted paths and file sizes for rollback and transparency.
//   - Prefer interactive confirmation for aggressive or ambiguous targets.
//
// Rule of thumb: if you're unsure — **don't delete automatically**.

func main() {
	err := Execute()
	if err != nil {
		return
	}
}

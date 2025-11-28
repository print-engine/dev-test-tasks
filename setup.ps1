# PowerShell Setup Script for Windows

Clear-Host

Write-Host "================================================================" -ForegroundColor Blue
Write-Host "" -ForegroundColor Blue
Write-Host "     Senior Backend Developer Assessment Setup" -ForegroundColor Blue
Write-Host "     WebSocket Notifications Service in Go" -ForegroundColor Blue
Write-Host "" -ForegroundColor Blue
Write-Host "================================================================" -ForegroundColor Blue
Write-Host ""
Write-Host "Welcome to the technical assessment!" -ForegroundColor Yellow
Write-Host ""
Write-Host "This setup script will configure your environment with the"
Write-Host "Anthropic API key needed for Claude Code assistance."
Write-Host ""
Write-Host "Please paste the API key provided by the interviewer:" -ForegroundColor Yellow

$PlainKey = Read-Host

if ([string]::IsNullOrWhiteSpace($PlainKey)) {
    Write-Host ""
    Write-Host "WARNING: No API key provided. Please run the setup again." -ForegroundColor Yellow
    Read-Host "Press Enter to exit"
    exit 1
}

# Set for current session
$env:ANTHROPIC_API_KEY = $PlainKey

# Set for future sessions (user level)
[System.Environment]::SetEnvironmentVariable("ANTHROPIC_API_KEY", $PlainKey, [System.EnvironmentVariableTarget]::User)

Write-Host ""
Write-Host "✓ API key configured successfully!" -ForegroundColor Green
Write-Host ""
Write-Host "Verifying... API key is set in your environment." -ForegroundColor Green
Write-Host ""
Write-Host "================================================================" -ForegroundColor Blue
Write-Host "                    You're All Set! 🚀" -ForegroundColor Blue
Write-Host "================================================================" -ForegroundColor Blue
Write-Host ""
Write-Host "Next steps:" -ForegroundColor Green
Write-Host ""
Write-Host "  1. Type: " -NoNewline
Write-Host "claude" -ForegroundColor Yellow
Write-Host "  2. Start vibing! ✨"
Write-Host ""
Write-Host "Tips:" -ForegroundColor Blue
Write-Host "  • Read the " -NoNewline
Write-Host "TASK.md" -ForegroundColor Yellow -NoNewline
Write-Host " file for requirements"
Write-Host "  • The " -NoNewline
Write-Host "CLAUDE.md" -ForegroundColor Yellow -NoNewline
Write-Host " file has helpful context"
Write-Host "  • Ask Claude questions - this is about AI collaboration!"
Write-Host "  • Focus on core features first, then stretch goals"
Write-Host ""
Write-Host "Good luck! 🎯" -ForegroundColor Green
Write-Host ""
Read-Host "Press Enter to continue"
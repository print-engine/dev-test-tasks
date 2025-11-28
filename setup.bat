@echo off

cls

echo ================================================================
echo.
echo      Senior Backend Developer Assessment Setup
echo      WebSocket Notifications Service in Go
echo.
echo ================================================================
echo.
echo Welcome to the technical assessment!
echo.
echo This setup script will configure your environment with the
echo Anthropic API key needed for Claude Code assistance.
echo.
echo Please paste the API key provided by the interviewer:
set /p ANTHROPIC_API_KEY=

if "%ANTHROPIC_API_KEY%"=="" (
    echo.
    echo WARNING: No API key provided. Please run the setup again.
    pause
    exit /b 1
)

:: Export for future sessions
setx ANTHROPIC_API_KEY "%ANTHROPIC_API_KEY%" >nul 2>&1

echo.
echo OK API key configured successfully!
echo.
echo IMPORTANT: The API key is now set for future terminal sessions.
echo For it to work in THIS terminal, please run this command:
echo.
echo   set ANTHROPIC_API_KEY=%ANTHROPIC_API_KEY%
echo.
echo Or simply open a NEW Command Prompt window.
echo.
echo ================================================================
echo                     You're All Set!
echo ================================================================
echo.
echo Next steps:
echo.
echo   1. Open a NEW Command Prompt (or run the set command above)
echo   2. Type: claude
echo   3. Start vibing!
echo.
echo Tips:
echo   * Read the TASK.md file for requirements
echo   * The CLAUDE.md file has helpful context
echo   * Ask Claude questions - this is about AI collaboration!
echo   * Focus on core features first, then stretch goals
echo.
echo Good luck!
echo.
pause
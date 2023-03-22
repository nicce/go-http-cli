#!/bin/bash

# Get the latest tag
latest_tag=$(git describe --tags --always --abbrev=0)

# Get the previous tag
previous_tag=$(git describe --tags --always --abbrev=0 "$latest_tag^")

# Get the commit logs since the last tag
logs=$(git log "$previous_tag".."$latest_tag" --no-merges --all --pretty=format:'%s (%h)')

# Initialize the changelog
changelog="## Changelog\n\n"

# Parse the commit logs and group them by commit type
while read -r log; do
  commit_type=$(echo "$log" | cut -d '(' -f1)
  case $commit_type in
    feat*)
      feat_logs+="* ${log}\n"
      ;;
    !feat*)
      feat_logs+="* ${log} 💥breaking change\n"
      ;;
    fix*)
      fix_logs+="* ${log}\n"
      ;;
    docs*)
      docs_logs+="* ${log}\n"
      ;;
    refactor*)
      refactor_logs+="* ${log}\n"
      ;;
    perf*)
      perf_logs+="* ${log}\n"
      ;;
    test*)
      test_logs+="* ${log}\n"
      ;;
    chore*)
      chore_logs+="* ${log}\n"
      ;;
    build*)
      build_logs+="* ${log}\n"
      ;;
    *)
      unknown_logs+="* $log\n"
      ;;
  esac
done <<< "$logs"

# Build the changelog by combining the commit type sections
if [ -n "$feat_logs" ]; then
  changelog+="### 🚀 Features\n\n$feat_logs\n"
fi

if [ -n "$fix_logs" ]; then
  changelog+="### 🐛 Bug Fixes\n\n$fix_logs\n"
fi

if [ -n "$docs_logs" ]; then
  changelog+="### 📝 Documentation\n\n$docs_logs\n"
fi

if [ -n "$refactor_logs" ]; then
  changelog+="### 🔧 Refactors\n\n$refactor_logs\n"
fi

if [ -n "$perf_logs" ]; then
  changelog+="### ⚡ Performance Improvements\n\n$perf_logs\n"
fi

if [ -n "$test_logs" ]; then
  changelog+="### ✅ Tests\n\n$test_logs\n"
fi

if [ -n "$chore_logs" ]; then
  changelog+="### 🧹 Chores\n\n$chore_logs\n"
fi

if [ -n "$build_logs" ]; then
  changelog+="### 🤖 Build\n\n$build_logs\n"
fi

if [ -n "$unknown_logs" ]; then
  changelog+="### ❓ Other Changes\n\n$unknown_logs\n"
fi

# Output the changelog
echo -e "$changelog"


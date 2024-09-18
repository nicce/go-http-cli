#!/bin/bash
# This script calculate the next version number based on conventional commits.
# At the moment only, major, feat and fix gets a new version number.

# Set the starting version number
# Get the latest tag
latest_tag=$(git describe --tags --always --abbrev=0)

should_bump_major=false
should_bump_minor=false
should_bump_revision=false

# initialize version
version=""

# Get the commit messages between the two tags
commits=$(git log --pretty=format:%s --always "$latest_tag"..HEAD)

# Loop through the commit messages and flag what versions should be bumped.
for commit in $commits
do
  case $commit in
    !feat*)
      should_bump_major=true
      break;
      ;;
    feat*) # Increment the minor version number for a new feature
      should_bump_minor=true
      ;;
    fix*|build\(deps\)*) # Increment the patch version number for a bug fix
      should_bump_revision=true
      ;;
    *) # Ignore other commit types
      ;;
  esac
done

if $should_bump_major ; then version=v$(echo "$latest_tag" | awk -F"." '{print v$1+1 ".0" ".0"}')
elif $should_bump_minor ; then version=$(echo "$latest_tag" | awk -F"." '{print $1 "." $2+1 ".0"}')
elif $should_bump_revision ; then version=$(echo "$latest_tag" | awk -F"." '{print $1 "." $2 "." $3+1}') ; fi


# Output the final version number
echo "$version"

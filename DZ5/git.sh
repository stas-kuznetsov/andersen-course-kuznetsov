#!/bin/bash

if [ "$#" -eq "0" ]
  then
    echo 'you need to use this script with repository address in format:'
    echo 'git.sh https://github.com/$user/$repo'
    echo 'Also you can to enter repository address here: '
    read repo
  else
    repo=$1
fi

repo_user_and_name=${repo#*github.com/}
repo_name=${repo_user_and_name#*/}

number_requests=$(curl https://api.github.com/repos/$repo_user_and_name/pulls | jq .[].title | wc -l)


echo "The number of pull requests in repo $repo is $number_requests"

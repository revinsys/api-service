#!/bin/bash
argsCount=$#
echo $argsCount
if [[ "$argsCount" == "1" ]]
then
        echo "Run replace"
        gitPath=$(echo $1 | sed 's/\//\\\//g')
        findFiles=$(find ./ -name "*.go")

        for file in $findFiles
        do
                filePath=$(echo `pwd`$file | sed 's/\.\/\//\//g')
                sed -i '' "s/{GIT_PATH}/$gitPath/" $filePath
        done

        echo "Remove template git and init new git project"
        rm -rf ./.git
        git init > /dev/null
        git remote add origin https://$1 > /dev/null

        echo "Initialize go project"
        go mod init $1 > /dev/null
        go get github.com/gorilla/mux
        go get github.com/sirupsen/logrus
        go get github.com/jmoiron/sqlx
        go get github.com/lib/pq
else
        echo "Run command on format: ./init.sh <git_path>"
fi

#!/bin/bash
#bash command to import csv data to database
#pay attention to db name and its directory
#.mode csv

srcTable=$1
dstTable=$2
srcCSVDirectory=$3
dstCSVDirectory=$4

sqlite3 ipmap.db "create table '"$srcTable"'(ipmap TEXT);"
sqlite3 ipmap.db "create table '"$dstTable"'(ipmap TEXT);"

for f in $srcCSVDirectory
do
    sqlite3 ipmap.db ".import '"$f"' $srcTable"
done

for f in $dstCSVDirectory
do
    sqlite3 ipmap.db ".import '"$f"' $dstTable"
done
# sqlite3 ipmap.db ".import /tmp/city.csv dest"

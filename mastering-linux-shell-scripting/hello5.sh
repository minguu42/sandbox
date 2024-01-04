#!/bin/bash
if [ "$#" -lt 1 ] ; then
	read -p "Enter a name: "
	name=$REPLY
else
	name=$1
fi
echo "Hello $name"
if [ "mokhtar" \> "Mokhtar" ]; then
	echo "文字列1は文字列2より大きい"
else
	echo "文字列1は文字列2より小さい"
fi
exit 0

#!/bin/bash
shopt -s nocasematch # 大文字と小文字の区別をオフにする

read -p "Type color or mono for script output: "
if [[ $REPLY =~ colou?r ]] ; then
    source $HOME/snippets/color
fi

echo -e "${GREEN}This is $0 $RESET"

shopt -u nocasematch # 大文字と小文字の区別をリセットする
exit 0

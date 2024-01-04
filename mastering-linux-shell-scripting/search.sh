#!/bin/bash
usage="使用法: search.sh <ファイル> <検索文字列> <操作>"
if [ ! "$#" -eq 3 ] ; then
  echo "$usage"
  exit 2
fi

[ ! -f "$1" ] && exit 3

case "$3" in
  [cC])
    msg="$1の中で$2にマッチする行数を数えます"
    opt="-c"
    ;;
  [pP])
    msg="$1の中で$2にマッチする行を表示します"
    opt=""
    ;;
  [dD])
    msg="$1から$2にマッチする行を除いて全て表示します"
    opt="-v"
    ;;
  *)
    echo "$1 $2 $3を評価できません"
    exit 1
    ;;
esac

echo $msg
grep $opt $2 $1
exit 0

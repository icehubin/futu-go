#!/bin/sh

#直接引用的 https://github.com/FutunnOpen/py-futu-api/tree/master/futu/common/pb

sed -i "" "s:github.com/futuopen/ftapi4go/::g" `ls *.proto`
ls *.proto |awk '{print "protoc -I=./ --go_out=../ ./"$1}'|sh

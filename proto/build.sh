#!/bin/sh

#直接引用的 https://github.com/FutunnOpen/py-futu-api/tree/master/futu/common/pb

sed -i "" "s:github.com/futuopen/ftapi4go/:github.com/icehubin/futu-go/:g" `ls *.proto`
ls *.proto |awk '{print "protoc -I=./ --go_out=../ ./"$1}'|sh
sed -i "" "s:github.com/icehubin/futu-go/:github.com/futuopen/ftapi4go/:g" `ls *.proto`
cp -rf ../github.com/icehubin/futu-go/pb/* ../pb
rm -rf ../github.com

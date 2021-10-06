#!/bin/sh

#使用模板快速生成一个adapt，然后修改代码

PB=$1
CLASS=${PB/_/}
INPORT=`echo ${CLASS} | tr 'A-Z' 'a-z'`

if [ ! -f "../futuproto/${PB}.proto" ]; then
    echo "file not exist"
    exit 0
fi
cp ExampleAdapt.go ${CLASS}.go
sed -i "" "s:ExampleAdapt:${CLASS}:g" ${CLASS}.go
sed -i "" "s:Example_Adapt:${PB}:g" ${CLASS}.go
sed -i "" "s:keepalive:${INPORT}:g" ${CLASS}.go
cp ExampleAdapt_test.go ${CLASS}_test.go
sed -i "" "s:ExampleAdapt:${CLASS}:g" ${CLASS}_test.go
sed -i "" "s:Example_Adapt:${PB}:g" ${CLASS}_test.go

echo "gen ${PB} success"
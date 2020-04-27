#!/bin/bash

d="" # delimiter
q="'"

ack "console\\.(log|error)" src | while read line
do
	arr=(${line//:/ })
	filename=${arr[0]}
	linenum=${arr[1]}

	a="$filename $linenum"

	m='console.([a-z]+)\(('$q'\[.*\]\\n'$q', )?(.*)\)'
	r='console.\1('$q'\['$a'\]\\n'$q', \3)'
	str="${linenum}s$d${m}$d${r}$d"

	cp $filename ${filename}_
	sed -r "$str" ${filename}_ > $filename
	rm ${filename}_
done

#!/bin/bash
g++ /shared/code.cpp -o /a.out 2>/shared/erroresCompilacion.txt
if [ $? = 0 ];then
	rm /shared/erroresCompilacion.txt
	for entry in "/shared/inputs"/*
	do
		inputID=$(echo $entry | sed 's/^\/shared\/inputs\///g')
		timeout 60s /a.out <"$entry"  >/shared/outputs/$inputID
		if [ $? = 124 ];then
			echo "tiempoExcedido" >/shared/errors/$inputID
		elif [ $? != 0 ];then
			echo "errorRuntime" >/shared/errors/$inputID
		else
			echo "correcto" >/shared/errors/$inputID
		fi
	done
fi
#!/bin/bash
timeout 60s g++ /shared/code.cpp -o /a.out 2>/shared/erroresCompilacion.txt
resp_comp=$?
if [ $resp_comp = 0 ];then
	rm /shared/erroresCompilacion.txt
	for entry in "/shared/inputs"/*
	do
		inputID=$(echo $entry | sed 's/^\/shared\/inputs\///g')
		timeout 60s /a.out <"$entry"  >/shared/outputs/$inputID
		resp_ejec=$?
		if [ $resp_ejec = 124 ];then
			echo "tiempoExcedido" >/shared/errors/$inputID
		elif [ $resp_ejec != 0 ];then
			echo "errorRuntime" >/shared/errors/$inputID
		else
			echo "correcto" >/shared/errors/$inputID
		fi
	done
elif [ $resp_comp = 124 ];then
	echo "Tiempo de compilacion excesivamente largo" >/shared/erroresCompilacion.txt
fi
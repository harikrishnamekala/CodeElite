gcc -lm /vol/main.c -o /vol/main 2> /vol/errors.txt
timeout 5 ./vol/main < /vol/input.txt  > /vol/data.txt

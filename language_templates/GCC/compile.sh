gcc -lm /home/main.c -o /home/main 2> /home/errors.txt
timeout 5 ./home/main < /home/input.txt  > /home/output.txt

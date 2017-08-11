gmcs /home/Solution.cs -out:/home/Solution.exe &> /home/errors.txt
timeout 5 mono /home/Solution.exe < /home/input.txt > /home/output.txt

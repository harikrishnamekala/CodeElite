timeout 5 vbnc -nologo -quiet /home/file.vb -out:/home/file.exe &> /home/errors.txt
mono /home/file.exe < /home/input.txt > /home/output.txt

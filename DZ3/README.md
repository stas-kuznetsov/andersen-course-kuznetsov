Actions of my script:

1. Script checks: how many arguments have entered by user

2. If user have not enterd any arguments: 
2.1. Outputs message at console "Enter process name or PID"
2.2. Reads value, that enter user (process name or PID) and saves it to variable 'ent'
2.3. Outputs another message at console "How maximum lines do you want to output?"
2.4. Reads value, that enter user (number of lines) and saves it to variable 'lines'

3. If user have entered only one argunent (name or PID process):
3.1. Reads value, that have entered by user in command line (process name or PID) and save it to variable 'ent'
3.2. Setups variable 'lines' to 5

4. If user have entered two (or more) arguments:
4.1 Reads first argument, that have entered by user in command line (process name or PID) and saves it to variable 'ent'
4.2 Reads second aegument, that have entered by user in command line (number of lines) and saves it to variable 'lines'

5. Run one-line command from Homework 3, but with my modifications: value of variable 'ent' (process name or PID) and value of variable 'lines' (number of lines) are enter into this one-line command

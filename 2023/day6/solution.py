#!usr/bin/env python3

FILE = open('problem.txt', 'r').read()

def sln_one():
    lines = FILE.split('\n')
    sum = 1

    time = [int(i) for i in lines[0].split(':')[1].split()]
    distance = [int(i) for i in lines[1].split(':')[1].split()]
    
    # enumerate through the races
    for r in range(len(time)):
        ways = 0
        
        for j in range(time[r]+1):
            if(j * (time[r] - j) > distance[r]):
                ways += 1
        
        sum *= ways 

    print(f'[+] part-1 solution: {sum}')

def sln_two():
    lines = FILE.split('\n')
    sum = 0

    time = int(''.join(lines[0].split(':')[1].split()))
    distance = int(''.join(lines[1].split(':')[1].split()))

    for t in range(time+1):
        if(t * (time - t) > distance):
            sum += 1

    print(f'[+] part-2 solution: {sum}')

if __name__ == '__main__':
    sln_one()
    sln_two()
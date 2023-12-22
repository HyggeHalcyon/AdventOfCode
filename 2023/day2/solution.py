#!usr/bin/env python3

FILE = open('problem.txt', 'r').read()
MAXRED = 12
MAXGREEN = 13
MAXBLUE = 14

def sln_one():
    lines = FILE.split('\n')
    sum = 0

    # enumerate each line
    for i, line in enumerate(lines):
        cubes = {
            'red': 0,
            'green': 0,
            'blue': 0
        }

        game, set = line.split(":")
        set = [s.strip() for s in set.replace(";", ",").split(",")] # convert sets of game into list of cubes
        
        # enumerate each cubes
        for s in set:
            num, cube = s.split(" ")
            if(int(num) > cubes[cube]):
                cubes[cube] = int(num)

        if(cubes['red'] <= MAXRED and cubes['green'] <= MAXGREEN and cubes['blue'] <= MAXBLUE):
            sum += i+1

    print(f'[+] part-1 solution: {sum}')

def sln_two():
    lines = FILE.split('\n')
    sum = 0

    # enumerate each line
    for i, line in enumerate(lines):
        cubes = {
            'red': 0,
            'green': 0,
            'blue': 0
        }

        game, set = line.split(":")
        set = [s.strip() for s in set.replace(";", ",").split(",")] # convert sets of game into list of cubes
        
        # enumerate each cubes
        for s in set:
            num, cube = s.split(" ")
            if(int(num) > cubes[cube]):
                cubes[cube] = int(num)

        sum += cubes['red'] * cubes['green'] * cubes['blue']

    print(f'[+] part-2 solution: {sum}')

if __name__ == '__main__':
    sln_one()
    sln_two()
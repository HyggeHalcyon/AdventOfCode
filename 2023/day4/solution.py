#!usr/bin/env python3

FILE = open('problem.txt', 'r').read()

def sln_one():
    lines = FILE.split('\n')
    sum = 0

    for line in lines:
        _, scratchcard = line.split(':')
        left, right = scratchcard.split('|')
        wcard = [int(c) for c in left.split()]
        pcard = [int(c) for c in right.split()]

        wins = len(set(wcard) & set(pcard))
        if wins > 0:
            sum += 2**(wins - 1)    
    
    print(f'[+] part-1 solution: {sum}')

def sln_two():
    lines = FILE.split('\n')
    scratchcards = [0] * len(lines)

    for i, line in enumerate(lines):
        scratchcards[i] += 1

        _, scratchcard = line.split(':')
        left, right = scratchcard.split('|')
        wcard = [int(c) for c in left.split()]
        pcard = [int(c) for c in right.split()]

        wins = len(set(wcard) & set(pcard))
        for w in range(wins):
            scratchcards[i+1+w] += scratchcards[i]

    print(f'[+] part-2 solution: {sum(scratchcards)}')

if __name__ == '__main__':
    sln_one()
    sln_two()
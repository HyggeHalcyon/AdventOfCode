#!/usr/bin/env python3

FILE = open('problem.txt', 'r').read()
WORD_TO_DIGITS = {
    "one": 1,
    "two": 2,
    "three": 3,
    "four": 4,
    "five": 5,
    "six": 6,
    "seven": 7,
    "eight": 8,
    "nine": 9
}

def isdigit(char: str) -> bool:
    char = ord(char)
    if(char >= 0x30 and char <= 0x39):
        return True
    return False

def sln_one():
    lines = FILE.split('\n')
    sum = 0

    # enumerate each line
    for line in lines:
        firstDigit, lastDigit = '', ''

        # enumerate each character
        for char in line:
           if isdigit(char):
                if(firstDigit == ''):
                    firstDigit = char
                lastDigit = char
        
        sum += int(firstDigit + lastDigit)

    print(f'[+] part-1 solution: {sum}')

def sln_two():
    lines = FILE.split('\n')
    sum = 0

    # enumerate each line
    for line in lines:
        firstDigit, lastDigit = '', ''

        # enumerate each character
        for idx, char in enumerate(line):
            num = 0

            if isdigit(char):
                num = char
            else:
                for key in WORD_TO_DIGITS.keys():
                    if(line[idx:].startswith(key)):
                        num = str(WORD_TO_DIGITS[key])
                        break
                else:
                    continue

            if(firstDigit == ''):
                firstDigit = num
            lastDigit = num
        
        sum += int(firstDigit + lastDigit)

    print(f'[+] part-2 solution: {sum}')

if __name__ == '__main__':
    sln_one()
    sln_two()
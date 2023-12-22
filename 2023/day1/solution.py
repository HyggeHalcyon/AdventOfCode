#!/usr/bin/env python3
# --- Day 1: Trebuchet?! ---
import sys

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

def part_one_sln(filename: str) -> int:
    # --- Part One ---
    sln = 0

    with open(filename) as file:
        for line in file:
            firstDigit = -1
            lastDigit = -1

            for char in line:
                if(isdigit(char)):
                    if(firstDigit == -1):
                        firstDigit = int(char)
                    lastDigit = int(char)

            # print(f'[!] {line.strip()} -> {firstDigit}:{lastDigit}')
            sln += int(str(firstDigit) + str(lastDigit))
    
    return sln

def part_two_sln(filename: str) -> int:
    # --- Part Two ---
    sln = 0

    with open(filename) as file:
        for line in file:
            firstDigit = -1
            lastDigit = -1
            buff = ''

            for char in line:
                num = 0

                if(isdigit(char)):
                    num = int(char)
                    buff = ''
                else:
                    buff += char
                    for key in WORD_TO_DIGITS:
                        if key in buff:
                            num = WORD_TO_DIGITS[key]
                            buff = ''
                            break
                    else:
                        continue

                if(firstDigit == -1):
                    firstDigit = num
                lastDigit = num

            # print(f'[!] {line.strip()} -> {firstDigit}:{lastDigit} || {buff.strip()}')
            sln += int(str(firstDigit) + str(lastDigit))
    
    return sln

if __name__ == '__main__':
    if(len(sys.argv) < 2):
        print(f"[*] Usage: {sys.argv[0]} <file>")
        sys.exit(1)

    sln = part_one_sln(sys.argv[1]) # 55816
    print(f'[+] Solution Part One: {sln}')

    sln = part_two_sln(sys.argv[1]) # 54970 ?? why am i missing a 10?
    print(f'[+] Solution Part Two: {sln}')
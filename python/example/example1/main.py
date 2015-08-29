import random

'''
    猜数字游戏
'''
number = random.randint(0,100)
print("hello, Number guessing Game: between 0 and 100 inclusive.")

guessString = input("guess a number: ")
guess = int(guessString)

while 0 <= guess <=100:
    if guess > number:
        print("Guessed Too High, Please guess again!")
    elif guess < number:
        print("Guess Too Low, Please guess again!")
    else:
        print("You guessed it. The number was:", number)
        break
    guessString = input("Guess a number: ")
    guess = int(guessString)
else:
    print("Your guess was not between 0 and 100, the again!")

# Number guessing game cli!

A simple number guessing game where the computer generates a random number and the user has to guess it. The user has a limited number of tries to get the correct number. If the user guesses correctly, they win and the game ends. Otherwise, the game will continue until the user runs out of tries.

## Sample Game Play

```txt
Welcome to the Number Guessing Game!
I'm thinking of a number between 1 and 100.
You have 5 chances to guess the correct number.

Please select the difficulty level:
1. Easy (10 chances)
2. Medium (5 chances)
3. Hard (3 chances)

Enter your choice: 2

Great! You have selected the Medium difficulty level.
Let's start the game!

Enter your guess: 50
Incorrect! The number is less than 50.

Enter your guess: 25
Incorrect! The number is greater than 25.

Enter your guess: 35
Incorrect! The number is less than 35.

Enter your guess: 30
Congratulations! You guessed the correct number in 4 attempts.
```

### Additional Features Added.

1. Allow the user to play multiple rounds of the game (i.e., keep playing until the user decides to quit). You can do this by asking the user if they want to play again after each round.
2. Add a timer to see how long it takes the user to guess the number.
3. Implement a hint system that provides clues to the user if they are stuck.
4. Keep track of the user’s high score (i.e., the fewest number of attempts it took to guess the number under a specific difficulty level).

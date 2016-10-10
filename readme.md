# Rank
Rank is a simple **Ran**dom **K**ey generator.

usage examples:

```
# Default will return a 32 character hexadecimal key
rank
> 3072d4f55328059b285366252840c802

# generate a 50 character hexadecimal random key
rank -l 50
> 3ad3d0df8c7727c9671427d8a7b6824e2a1ff9d79492ea39a8

# generate multiple keys
rank -n 5
> db87d6d447c85e8e4940e08b4ece58c6
> 3523a1bdc5c64c2c589866876e612a7a
> 503a3c5ceb1453b98ffb7cc788a7a518
> 28d8af2f4ddfaeec410ff3171d82ec94
> 0283007be6c2f78c449001c337bde31c

# base64
rank -m base64
> z6QUAeuIvP0EZg-7NYpfwU1O4hMSZ4b6

# all together
rank -l 50 -n 2 -m base64
> guSOMF0y0hpD0eg4edkY9CjTEmI5yI9rEWyzGoOvQ6sO9bmEqE
> 7HVxLCyKUKKT7H50zGAyqzxqSVB-VZPGYbrmnOpn6UYVfTs0Jh
```

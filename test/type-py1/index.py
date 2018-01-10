import sys

who = sys.argv[1] if len(sys.argv) > 1 and sys.argv[1] else 'Python'

print('Hi ' + who + "!")

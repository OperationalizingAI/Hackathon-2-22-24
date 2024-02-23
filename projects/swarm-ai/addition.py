def add(x, y):
    if not (isinstance(x, (int, float)) and isinstance(y, (int, float))):
        raise TypeError('Inputs must be numbers')
    return x + y
import unittest
from fibonacci import fibonacci

class TestFibonacci(unittest.TestCase):
    def test_first_number(self):
        self.assertEqual(fibonacci(1), 0)

    def test_second_number(self):
        self.assertEqual(fibonacci(2), 1)
    
    def test_third_number(self):
        self.assertEqual(fibonacci(3), 1)
    
    def test_sixth_number(self):
        self.assertEqual(fibonacci(6), 5)

    def test_thirteenth_number(self):
        self.assertEqual(fibonacci(13), 144)

if __name__ == '__main__':
    unittest.main()
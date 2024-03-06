import unittest
from addition import add

class TestAddFunction(unittest.TestCase):
    def test_add_positive_integers(self):
        self.assertEqual(add(1, 2), 3)

    def test_add_negative_integers(self):
        self.assertEqual(add(-1, -2), -3)

    def test_add_positive_and_negative_integer(self):
        self.assertEqual(add(-1, 2), 1)

    def test_add_zero(self):
        self.assertEqual(add(0, 0), 0)

    # If you want your function to handle floats as well
    def test_add_floats(self):
        self.assertAlmostEqual(add(1.1, 2.2), 3.3, places=2)

    # If you want to ensure type checking (e.g., both arguments must be int or float)
    def test_add_non_numbers(self):
        with self.assertRaises(TypeError):
            add('a', 'b')

if __name__ == '__main__':
    unittest.main()

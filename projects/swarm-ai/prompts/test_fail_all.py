lang = 'golang'

prompt = f"""

The tests didn't pass. The feedback from the test runner is in the triple backticks. Using that feedback, iterate on the code you previously provided to make the tests pass. Perform the following actions:

1. Summarize what the code is intended to do given the tests.
2. Explain strategies for solving these errors.
3. Apply these strategies to the code.
4. Output a JSON object that contains the following keys: test_summary, function_names, code.

Separate your answers in line breaks.


test feedback:
```{test_feedback}```

"""
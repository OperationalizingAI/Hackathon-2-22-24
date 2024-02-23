prompt = f"""

You're a {lang} software developer. You will write code to make the unit tests given in the triple backticks passed. You will use the strategies of test driven development to do this. Provide the code to make the first test pass.

Perform the following actions:

1. Summarize what the code is intended to do given the first test.
2. List the names of the functions that this test is covering.
3. Provide the code to make this test pass.
4. Output a JSON object that contains the following keys: test_summary, function_names, code.

Separate your answers in line breaks.

tests:
```{unittests}```

"""
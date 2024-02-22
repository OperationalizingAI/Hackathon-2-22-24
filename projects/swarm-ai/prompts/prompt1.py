lang = 'golang'

prompt = f"""

You're a {lang} sofware developer. You will be given {error_log}. Inspect the errors, revise the written code to ensure the unit tests given in the triple backticks passed. 

Perform the following actions:

1. Summarize what the code is intended to do given the tests.
2. List the names of the functions that these tests are covering.
3. Provide the code.
4. Output a JSON object that containts the following keys: test_summary, function_names, code.

Separate your answers in line breaks.


tests:
```{unittests}```

"""
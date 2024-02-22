
prompt = f"""

The code didn't compile. You will be given the error log. Inspect the errors, revise the written code to ensure the unit tests given in the triple backticks passed. 

Perform the following actions:

1. Summarize what the code is intended to do given the tests.
2. List the names of the functions that these tests are covering.
3. Provide the code.
4. Highlight diffs.
5. Output a JSON object that containts the following keys: test_summary, function_names, code, diffs.


Separate your answers in line breaks.


error log:
```{error_log}```

"""
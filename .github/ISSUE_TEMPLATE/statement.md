---
name: Statement
about: Problem statement
title: Technical interview exercise for [NAME]
labels: ''
assignees: miguelff

---

:wave:, 

### Build a simple data pipeline

You will need [docker](https://docs.docker.com/get-docker/) installed. The following [command](https://github.com/hasura/hiring_log_gen#usage)  generates log data:

`docker run hasura/hiring_log_gen`

A log line has the shape below. Units for sizes are expressed in bytes, and times in seconds.

```
2022/04/05 18:38:34 {"project_id":"9a6e673e-c404-44be-9ec7-f608cdad895a","timestamp":"2022-04-05T18:38:34.10973506+02:00","operation":{"name":"insert_magazine","runtime":0.49770060098541913,"request_id":"298f63b9-f171-40ee-9730-6c608da3035c","response_size":25,"request_size":41,"http_status":200}}
```

The exercise consists of building a pipeline that would let us understand:

* The amount of data transferred by any project in a specific time window, so we can calculate a project's bill.

* What are, per project, the most time-consuming queries.

* For the sake of the exercise, don't worry about the volume of data yet, assume we will receive at most 1K requests per minute. You can tweak how fast the log generator will generate traces, by using the `-projects`, `-rate`, and `-speed` [parameters](https://github.com/hasura/hiring_log_gen#usage).

### Other requirements:

* Provide us with a way to execute the solution, this can be a GitHub repository containing a docker-compose file and instructions to run the pipeline.
* Provide a description of the solution, its trade-offs, and the compromises you made, in the form of a README.md, file in the repository, or 


Bonus points:

* Simplicity
* Presence of tests
* If you identify other potentially relevant information for the business or the system's operators, that can be extracted from the log.
* If you provide a good UX/DX in your solution.
* If you Identify the scalability challenges that the pipeline would face if we received 100x the volume of log data and some potential solutions.

You have as much time as you want to solve the exercise, but don't aim at building the perfect solution, just one that's simple and good enough for the requirements described.

Thank you for your time, looking forward to hearing back from you!

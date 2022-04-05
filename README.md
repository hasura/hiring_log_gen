# hiring_log_gen

Log generator for technical challenges that in Hasura's hiring processes.

## Usage

```
% docker run hasura/hiring_log_gen -h
  -projects uint
        number of projects. Valid between 1 and 1000. (default 25)
  -rate uint
        traces to emit per second, accross all projects. Valid between 1 and 10000 (default 1000)
  -speed uint
        relative speed of time. Valid between 1 -realtime- and 86400 -each second counts for a day-. (default 360)

```

```
% docker run hasura/hiring_log_gen -projects 2 -rate 10 -speed 180
```

Generates something like:

```
2022/04/05 18:38:34 {"project_id":"e615f4da-d790-499f-a67a-07f288c9b3ee","timestamp":"2022-04-05T18:38:34.1049305+02:00","operation":{"name":"update_greece","runtime":0.34890843682963724,"request_id":"8d5f3a74-73e6-4ac2-8b0e-71279b8580be","response_size":70,"request_size":391,"http_status":200}}
2022/04/05 18:38:34 {"project_id":"9a6e673e-c404-44be-9ec7-f608cdad895a","timestamp":"2022-04-05T18:38:34.10973506+02:00","operation":{"name":"insert_magazine","runtime":0.49770060098541913,"request_id":"298f63b9-f171-40ee-9730-6c608da3035c","response_size":25,"request_size":41,"http_status":200}}
2022/04/05 18:38:34 {"project_id":"e615f4da-d790-499f-a67a-07f288c9b3ee","timestamp":"2022-04-05T18:38:36.128352+02:00","operation":{"name":"select_planet","runtime":0.3610665872273873,"request_id":"92f236db-0a20-4610-affe-a088f2fe395e","response_size":82,"request_size":163,"http_status":200}}
2022/04/05 18:38:34 {"project_id":"9a6e673e-c404-44be-9ec7-f608cdad895a","timestamp":"2022-04-05T18:38:37.935592+02:00","operation":{"name":"update_piano","runtime":0.7038179451558593,"request_id":"44493055-0b6b-4c22-827b-22678420225c","response_size":97,"request_size":217,"http_status":200}}
2022/04/05 18:38:34 {"project_id":"e615f4da-d790-499f-a67a-07f288c9b3ee","timestamp":"2022-04-05T18:38:39.73786156+02:00","operation":{"name":"select_helicopter","runtime":0.9756046885083646,"request_id":"12d65025-f2e6-4e15-b5c4-bc119bf0ce45","response_size":87,"request_size":808,"http_status":200}}
...
```

## Questions?

Ask <a href="mailto:miguel@hasura.io">the maintainer</a>

Licensed under the MIT [license](./LICENSE)



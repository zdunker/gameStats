# game statistics

## playaround project with some fun game stats, doing some analysis, trends, etc.

### 0.0.1 keep design as simple as possible

#### configs, infrastructure(db, cache, file server ...), should be initialized and save in the main process memory, which will not change until a full server redeployment.
#### once server is up, requests coming in should first visit our own data storage, use caution to make OpenDota API calls unless we don't have the requested data in our storage.
#### any requested data we get back from OpenDota API should be saved in our storage.
#### data saved in our storage could serve user's requests, and also serve our own background data analysis tasks.


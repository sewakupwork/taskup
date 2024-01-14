
# Taskup

Create Api's for manage Task CRUD operations to send reminders for task to specific user.


## Deployment

 To run this project

```bash
  docker-compose up --build -d
```

wait for some time, build will take some time.

meanwhile

import collection to postman. file is already in github repo.

taskup-api.postman_collection.json

perform testing

# To stop the project gracefully.
```
docker-compose down
```

# Note

make sure port 8080 is available in your system. you need to tune the port in docker compose file.
